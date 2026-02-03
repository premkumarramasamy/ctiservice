// Package client implements the CTI client connection management.
package client

import (
	"ctiservice/internal/messages"
	"ctiservice/internal/protocol"
	"fmt"
	"io"
	"net"
)

// Reader reads and parses CTI messages from a TCP connection.
type Reader struct {
	conn     net.Conn
	registry *messages.Registry
}

// NewReader creates a new message reader.
func NewReader(conn net.Conn) *Reader {
	return &Reader{
		conn:     conn,
		registry: messages.NewRegistry(),
	}
}

// ReadMessage reads and parses a complete CTI message.
func (r *Reader) ReadMessage() (protocol.Message, error) {
	// Read the 8-byte header
	header, err := protocol.ReadHeader(r.conn)
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("connection closed: %w", err)
		}
		return nil, fmt.Errorf("failed to read message header: %w", err)
	}

	// Validate message length
	if header.MessageLength > protocol.MaxMessageSize {
		return nil, fmt.Errorf("message length %d exceeds maximum %d",
			header.MessageLength, protocol.MaxMessageSize)
	}

	// Read the message body
	body := make([]byte, header.MessageLength)
	if header.MessageLength > 0 {
		if _, err := io.ReadFull(r.conn, body); err != nil {
			return nil, fmt.Errorf("failed to read message body: %w", err)
		}
	}

	// Parse the message
	msg, err := r.registry.Parse(header.MessageType, body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse message: %w", err)
	}

	return msg, nil
}

// RawMessage contains the header and unparsed body of a message.
type RawMessage struct {
	Header *protocol.Header
	Body   []byte
}

// ReadRawMessage reads a message without parsing the body.
func (r *Reader) ReadRawMessage() (*RawMessage, error) {
	header, err := protocol.ReadHeader(r.conn)
	if err != nil {
		return nil, err
	}

	if header.MessageLength > protocol.MaxMessageSize {
		return nil, fmt.Errorf("message length %d exceeds maximum %d",
			header.MessageLength, protocol.MaxMessageSize)
	}

	body := make([]byte, header.MessageLength)
	if header.MessageLength > 0 {
		if _, err := io.ReadFull(r.conn, body); err != nil {
			return nil, err
		}
	}

	return &RawMessage{Header: header, Body: body}, nil
}
