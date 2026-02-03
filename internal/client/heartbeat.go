package client

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"
)

// Heartbeat manages the heartbeat mechanism for keeping the session alive.
type Heartbeat struct {
	interval    time.Duration
	sendFunc    func(invokeID uint32) error
	onFailure   func()
	logger      *slog.Logger

	mu           sync.Mutex
	unconfirmed  int
	lastInvokeID atomic.Uint32
	confirmed    chan uint32
	running      bool
}

// NewHeartbeat creates a new heartbeat manager.
func NewHeartbeat(interval time.Duration, sendFunc func(invokeID uint32) error, onFailure func(), logger *slog.Logger) *Heartbeat {
	return &Heartbeat{
		interval:  interval,
		sendFunc:  sendFunc,
		onFailure: onFailure,
		logger:    logger,
		confirmed: make(chan uint32, 10),
	}
}

// Run starts the heartbeat goroutine.
// It sends HEARTBEAT_REQ messages at the configured interval.
// If 3 heartbeats go unconfirmed, it calls the onFailure callback.
func (h *Heartbeat) Run(ctx context.Context) {
	h.mu.Lock()
	if h.running {
		h.mu.Unlock()
		return
	}
	h.running = true
	h.unconfirmed = 0
	h.mu.Unlock()

	ticker := time.NewTicker(h.interval)
	defer ticker.Stop()

	h.logger.Info("heartbeat started", "interval", h.interval)

	for {
		select {
		case <-ctx.Done():
			h.logger.Info("heartbeat stopped")
			h.mu.Lock()
			h.running = false
			h.mu.Unlock()
			return

		case <-ticker.C:
			h.mu.Lock()
			unconfirmed := h.unconfirmed
			h.mu.Unlock()

			if unconfirmed >= 3 {
				h.logger.Error("heartbeat failure: 3 heartbeats unconfirmed")
				h.mu.Lock()
				h.running = false
				h.mu.Unlock()
				h.onFailure()
				return
			}

			invokeID := h.lastInvokeID.Add(1)
			if err := h.sendFunc(invokeID); err != nil {
				h.logger.Error("failed to send heartbeat", "error", err)
				h.mu.Lock()
				h.unconfirmed++
				h.mu.Unlock()
			} else {
				h.logger.Debug("heartbeat sent", "invokeID", invokeID)
				h.mu.Lock()
				h.unconfirmed++
				h.mu.Unlock()
			}

		case invokeID := <-h.confirmed:
			h.logger.Debug("heartbeat confirmed", "invokeID", invokeID)
			h.mu.Lock()
			if h.unconfirmed > 0 {
				h.unconfirmed--
			}
			h.mu.Unlock()
		}
	}
}

// Confirm signals that a heartbeat confirmation was received.
func (h *Heartbeat) Confirm(invokeID uint32) {
	select {
	case h.confirmed <- invokeID:
	default:
		// Channel full, that's OK - we just need to clear the unconfirmed count
		h.mu.Lock()
		h.unconfirmed = 0
		h.mu.Unlock()
	}
}

// Stop stops the heartbeat.
func (h *Heartbeat) Stop() {
	h.mu.Lock()
	h.running = false
	h.mu.Unlock()
}

// IsRunning returns true if the heartbeat is active.
func (h *Heartbeat) IsRunning() bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.running
}
