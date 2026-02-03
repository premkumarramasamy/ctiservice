package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Header represents the GED-188 message header.
// All messages start with this 8-byte header.
type Header struct {
	MessageLength uint32 // Length of the message body (excludes header)
	MessageType   uint32 // Message type identifier
}

// ReadHeader reads an 8-byte message header from the reader.
func ReadHeader(r io.Reader) (*Header, error) {
	buf := make([]byte, HeaderSize)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	return &Header{
		MessageLength: binary.BigEndian.Uint32(buf[0:4]),
		MessageType:   binary.BigEndian.Uint32(buf[4:8]),
	}, nil
}

// Write writes the header to the writer in big-endian format.
func (h *Header) Write(w io.Writer) error {
	buf := make([]byte, HeaderSize)
	binary.BigEndian.PutUint32(buf[0:4], h.MessageLength)
	binary.BigEndian.PutUint32(buf[4:8], h.MessageType)

	_, err := w.Write(buf)
	return err
}

// Bytes returns the header as a byte slice.
func (h *Header) Bytes() []byte {
	buf := make([]byte, HeaderSize)
	binary.BigEndian.PutUint32(buf[0:4], h.MessageLength)
	binary.BigEndian.PutUint32(buf[4:8], h.MessageType)
	return buf
}

// String returns a human-readable representation of the header.
func (h *Header) String() string {
	return fmt.Sprintf("Header{Type=%s(%d), Length=%d}",
		MessageTypeName(h.MessageType), h.MessageType, h.MessageLength)
}
