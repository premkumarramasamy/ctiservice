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
	// Need at least 3 bytes for tag (2) + length (1)
	return p.offset+3 <= len(p.data)
}

// Next reads the next floating field.
// Returns tag, data, and any error.
func (p *FloatingFieldParser) Next() (uint16, []byte, error) {
	if !p.HasMore() {
		return 0, nil, fmt.Errorf("no more floating fields")
	}

	// Read 2-byte tag
	tag := binary.BigEndian.Uint16(p.data[p.offset : p.offset+2])
	p.offset += 2

	// Read 1-byte length
	length := int(p.data[p.offset])
	p.offset++

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

// FloatingFields holds parsed floating fields from a message.
type FloatingFields struct {
	fields map[uint16][]byte
}

// ParseFloatingFields parses the floating part of a message.
func ParseFloatingFields(data []byte) (*FloatingFields, error) {
	parser := NewFloatingFieldParser(data)
	fields, err := parser.ParseAll()
	if err != nil {
		return nil, err
	}
	return &FloatingFields{fields: fields}, nil
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
func (w *FloatingFieldWriter) WriteBytes(tag uint16, data []byte) {
	if len(data) > 255 {
		data = data[:255] // Truncate to max length
	}

	// Write tag (2 bytes)
	w.buf.WriteUint16(tag)
	// Write length (1 byte)
	w.buf.WriteUint8(uint8(len(data)))
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
