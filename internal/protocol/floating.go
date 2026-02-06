package protocol

import (
	"encoding/binary"
	"fmt"
)

// FloatingFieldParser parses floating fields from message data.
type FloatingFieldParser struct {
	data   []byte
	offset int
}

// NewFloatingFieldParser creates a parser for the floating part of a message.
func NewFloatingFieldParser(data []byte) *FloatingFieldParser {
	return &FloatingFieldParser{
		data:   data,
		offset: 0,
	}
}

// HasMore returns true if there are more floating fields to parse.
func (p *FloatingFieldParser) HasMore() bool {
	// Need at least 4 bytes for tag (2) + length (2)
	return p.offset+4 <= len(p.data)
}

// Next reads the next floating field.
// Returns tag, data, and any error.
// Protocol Version 24: Tag is USHORT (2 bytes), Length is USHORT (2 bytes)
func (p *FloatingFieldParser) Next() (uint16, []byte, error) {
	if !p.HasMore() {
		return 0, nil, fmt.Errorf("no more floating fields")
	}

	// Read 2-byte tag (USHORT)
	tag := binary.BigEndian.Uint16(p.data[p.offset : p.offset+2])
	p.offset += 2

	// Read 2-byte length (USHORT)
	length := int(binary.BigEndian.Uint16(p.data[p.offset : p.offset+2]))
	p.offset += 2

	// Validate we have enough data
	if p.offset+length > len(p.data) {
		return 0, nil, fmt.Errorf("floating field length %d exceeds remaining data %d",
			length, len(p.data)-p.offset)
	}

	// Read data
	data := p.data[p.offset : p.offset+length]
	p.offset += length

	return tag, data, nil
}

// ParseAll parses all floating fields into a map.
func (p *FloatingFieldParser) ParseAll() (map[uint16][]byte, error) {
	fields := make(map[uint16][]byte)

	for p.HasMore() {
		tag, data, err := p.Next()
		if err != nil {
			return fields, err
		}
		fields[tag] = data
	}

	return fields, nil
}

// floatingFieldEntry stores a tag and its data for internal use.
type floatingFieldEntry struct {
	tag  uint16
	data []byte
}

// FloatingFields holds parsed floating fields from a message.
// Supports repeated fields with the same tag.
type FloatingFields struct {
	fields    map[uint16][]byte      // First occurrence of each tag
	allFields []floatingFieldEntry   // All fields in order (for repeated tags)
}

// ParseFloatingFields parses the floating part of a message.
func ParseFloatingFields(data []byte) (*FloatingFields, error) {
	parser := NewFloatingFieldParser(data)

	ff := &FloatingFields{
		fields:    make(map[uint16][]byte),
		allFields: make([]floatingFieldEntry, 0),
	}

	for parser.HasMore() {
		tag, fieldData, err := parser.Next()
		if err != nil {
			return nil, err
		}

		// Store in allFields for repeated tag support
		ff.allFields = append(ff.allFields, floatingFieldEntry{tag: tag, data: fieldData})

		// Store first occurrence in map for quick lookup
		if _, exists := ff.fields[tag]; !exists {
			ff.fields[tag] = fieldData
		}
	}

	return ff, nil
}

// Has returns true if the field with the given tag exists.
func (f *FloatingFields) Has(tag uint16) bool {
	_, ok := f.fields[tag]
	return ok
}

// GetBytes returns the raw bytes for a field, or nil if not found.
func (f *FloatingFields) GetBytes(tag uint16) []byte {
	return f.fields[tag]
}

// GetString returns a null-terminated string field.
func (f *FloatingFields) GetString(tag uint16) string {
	data := f.fields[tag]
	if data == nil {
		return ""
	}
	// Find null terminator
	for i, b := range data {
		if b == 0 {
			return string(data[:i])
		}
	}
	return string(data)
}

// GetUint16 returns a uint16 field.
func (f *FloatingFields) GetUint16(tag uint16) uint16 {
	data := f.fields[tag]
	if len(data) < 2 {
		return 0
	}
	return binary.BigEndian.Uint16(data)
}

// GetUint32 returns a uint32 field.
func (f *FloatingFields) GetUint32(tag uint16) uint32 {
	data := f.fields[tag]
	if len(data) < 4 {
		return 0
	}
	return binary.BigEndian.Uint32(data)
}

// GetInt32 returns an int32 field.
func (f *FloatingFields) GetInt32(tag uint16) int32 {
	return int32(f.GetUint32(tag))
}

// Tags returns all field tags present.
func (f *FloatingFields) Tags() []uint16 {
	tags := make([]uint16, 0, len(f.fields))
	for tag := range f.fields {
		tags = append(tags, tag)
	}
	return tags
}

// GetAllBytes returns all occurrences of a repeated field.
func (f *FloatingFields) GetAllBytes(tag uint16) [][]byte {
	var result [][]byte
	for _, field := range f.allFields {
		if field.tag == tag {
			result = append(result, field.data)
		}
	}
	return result
}

// GetAllStrings returns all string values for a repeated field.
func (f *FloatingFields) GetAllStrings(tag uint16) []string {
	allBytes := f.GetAllBytes(tag)
	result := make([]string, 0, len(allBytes))
	for _, data := range allBytes {
		// Find null terminator
		s := ""
		for i, b := range data {
			if b == 0 {
				s = string(data[:i])
				break
			}
		}
		if s == "" && len(data) > 0 {
			s = string(data)
		}
		result = append(result, s)
	}
	return result
}

// GetAllUint32 returns all uint32 values for a repeated field.
func (f *FloatingFields) GetAllUint32(tag uint16) []uint32 {
	allBytes := f.GetAllBytes(tag)
	result := make([]uint32, 0, len(allBytes))
	for _, data := range allBytes {
		if len(data) >= 4 {
			result = append(result, binary.BigEndian.Uint32(data))
		}
	}
	return result
}

// GetAllUint16 returns all uint16 values for a repeated field.
func (f *FloatingFields) GetAllUint16(tag uint16) []uint16 {
	allBytes := f.GetAllBytes(tag)
	result := make([]uint16, 0, len(allBytes))
	for _, data := range allBytes {
		if len(data) >= 2 {
			result = append(result, binary.BigEndian.Uint16(data))
		}
	}
	return result
}

// Count returns the number of occurrences of a tag.
func (f *FloatingFields) Count(tag uint16) int {
	count := 0
	for _, field := range f.allFields {
		if field.tag == tag {
			count++
		}
	}
	return count
}

// FloatingFieldWriter builds the floating part of a message.
type FloatingFieldWriter struct {
	buf *Buffer
}

// NewFloatingFieldWriter creates a writer for floating fields.
func NewFloatingFieldWriter() *FloatingFieldWriter {
	return &FloatingFieldWriter{
		buf: NewWriteBuffer(),
	}
}

// WriteString writes a string field.
func (w *FloatingFieldWriter) WriteString(tag uint16, s string) {
	data := append([]byte(s), 0) // Null-terminate
	w.WriteBytes(tag, data)
}

// WriteBytes writes a raw bytes field.
// Protocol Version 24: Tag is USHORT (2 bytes), Length is USHORT (2 bytes)
func (w *FloatingFieldWriter) WriteBytes(tag uint16, data []byte) {
	if len(data) > 65535 {
		data = data[:65535] // Truncate to max length (USHORT max)
	}

	// Write tag (2 bytes USHORT)
	w.buf.WriteUint16(tag)
	// Write length (2 bytes USHORT)
	w.buf.WriteUint16(uint16(len(data)))
	// Write data
	w.buf.Write(data)
}

// WriteUint16 writes a uint16 field.
func (w *FloatingFieldWriter) WriteUint16(tag uint16, v uint16) {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, v)
	w.WriteBytes(tag, data)
}

// WriteUint32 writes a uint32 field.
func (w *FloatingFieldWriter) WriteUint32(tag uint16, v uint32) {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, v)
	w.WriteBytes(tag, data)
}

// Bytes returns the encoded floating fields.
func (w *FloatingFieldWriter) Bytes() []byte {
	return w.buf.Bytes()
}
