package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Message is the interface that all CTI messages must implement.
type Message interface {
	// Type returns the message type ID.
	Type() uint32

	// Encode serializes the message body (without header) to bytes.
	Encode() ([]byte, error)

	// Decode deserializes the message body from bytes.
	Decode(data []byte) error
}

// Buffer provides helper methods for reading/writing message data.
type Buffer struct {
	*bytes.Buffer
}

// NewBuffer creates a new buffer for message encoding/decoding.
func NewBuffer(data []byte) *Buffer {
	return &Buffer{bytes.NewBuffer(data)}
}

// NewWriteBuffer creates a new empty buffer for writing.
func NewWriteBuffer() *Buffer {
	return &Buffer{new(bytes.Buffer)}
}

// ReadUint8 reads a single byte as uint8.
func (b *Buffer) ReadUint8() (uint8, error) {
	val, err := b.ReadByte()
	if err != nil {
		return 0, fmt.Errorf("failed to read uint8: %w", err)
	}
	return val, nil
}

// ReadInt8 reads a single byte as int8.
func (b *Buffer) ReadInt8() (int8, error) {
	val, err := b.ReadByte()
	if err != nil {
		return 0, fmt.Errorf("failed to read int8: %w", err)
	}
	return int8(val), nil
}

// ReadUint16 reads 2 bytes as big-endian uint16.
func (b *Buffer) ReadUint16() (uint16, error) {
	buf := make([]byte, 2)
	if _, err := io.ReadFull(b, buf); err != nil {
		return 0, fmt.Errorf("failed to read uint16: %w", err)
	}
	return binary.BigEndian.Uint16(buf), nil
}

// ReadInt16 reads 2 bytes as big-endian int16.
func (b *Buffer) ReadInt16() (int16, error) {
	val, err := b.ReadUint16()
	if err != nil {
		return 0, err
	}
	return int16(val), nil
}

// ReadUint32 reads 4 bytes as big-endian uint32.
func (b *Buffer) ReadUint32() (uint32, error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(b, buf); err != nil {
		return 0, fmt.Errorf("failed to read uint32: %w", err)
	}
	return binary.BigEndian.Uint32(buf), nil
}

// ReadInt32 reads 4 bytes as big-endian int32.
func (b *Buffer) ReadInt32() (int32, error) {
	val, err := b.ReadUint32()
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// ReadBytes reads n bytes from the buffer.
func (b *Buffer) ReadBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := io.ReadFull(b, buf); err != nil {
		return nil, fmt.Errorf("failed to read %d bytes: %w", n, err)
	}
	return buf, nil
}

// ReadFixedString reads a fixed-length null-padded string.
func (b *Buffer) ReadFixedString(n int) (string, error) {
	buf, err := b.ReadBytes(n)
	if err != nil {
		return "", err
	}
	// Find the null terminator
	for i, c := range buf {
		if c == 0 {
			return string(buf[:i]), nil
		}
	}
	return string(buf), nil
}

// WriteUint8 writes a uint8.
func (b *Buffer) WriteUint8(v uint8) error {
	return b.WriteByte(v)
}

// WriteInt8 writes an int8.
func (b *Buffer) WriteInt8(v int8) error {
	return b.WriteByte(byte(v))
}

// WriteUint16 writes a big-endian uint16.
func (b *Buffer) WriteUint16(v uint16) error {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, v)
	_, err := b.Write(buf)
	return err
}

// WriteInt16 writes a big-endian int16.
func (b *Buffer) WriteInt16(v int16) error {
	return b.WriteUint16(uint16(v))
}

// WriteUint32 writes a big-endian uint32.
func (b *Buffer) WriteUint32(v uint32) error {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, v)
	_, err := b.Write(buf)
	return err
}

// WriteInt32 writes a big-endian int32.
func (b *Buffer) WriteInt32(v int32) error {
	return b.WriteUint32(uint32(v))
}

// WriteFixedString writes a null-padded fixed-length string.
// If s is shorter than n, it's padded with nulls.
// If s is longer than n, it's truncated.
func (b *Buffer) WriteFixedString(s string, n int) error {
	buf := make([]byte, n)
	copy(buf, s)
	_, err := b.Write(buf)
	return err
}

// EncodeMessage creates a complete message with header and body.
func EncodeMessage(msg Message) ([]byte, error) {
	body, err := msg.Encode()
	if err != nil {
		return nil, fmt.Errorf("failed to encode message body: %w", err)
	}

	header := &Header{
		MessageLength: uint32(len(body)),
		MessageType:   msg.Type(),
	}

	result := make([]byte, HeaderSize+len(body))
	copy(result[0:HeaderSize], header.Bytes())
	copy(result[HeaderSize:], body)

	return result, nil
}
