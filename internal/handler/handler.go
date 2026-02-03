// Package handler defines event handlers for CTI messages.
package handler

import (
	"ctiservice/internal/protocol"
)

// EventHandler is the interface for handling CTI events.
type EventHandler interface {
	// Handle processes a received CTI message.
	Handle(msg protocol.Message)
}

// EventHandlerFunc is a function type that implements EventHandler.
type EventHandlerFunc func(msg protocol.Message)

// Handle implements EventHandler.
func (f EventHandlerFunc) Handle(msg protocol.Message) {
	f(msg)
}

// MultiHandler dispatches events to multiple handlers.
type MultiHandler struct {
	handlers []EventHandler
}

// NewMultiHandler creates a handler that dispatches to multiple handlers.
func NewMultiHandler(handlers ...EventHandler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

// Handle dispatches the message to all registered handlers.
func (m *MultiHandler) Handle(msg protocol.Message) {
	for _, h := range m.handlers {
		h.Handle(msg)
	}
}

// Add adds a handler to the multi-handler.
func (m *MultiHandler) Add(h EventHandler) {
	m.handlers = append(m.handlers, h)
}
