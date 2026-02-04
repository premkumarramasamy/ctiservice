package client

import (
	"context"
	"ctiservice/internal/config"
	"ctiservice/internal/messages"
	"ctiservice/internal/protocol"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"
)

// EventHandler is called when an event message is received.
type EventHandler func(msg protocol.Message)

// Client manages the connection to a CTI server.
type Client struct {
	cfg       *config.Config
	logger    *slog.Logger
	handler   EventHandler
	session   *Session
	heartbeat *Heartbeat

	mu        sync.Mutex
	conn      net.Conn
	reader    *Reader
	closeChan chan struct{}
}

// New creates a new CTI client.
func New(cfg *config.Config, logger *slog.Logger, handler EventHandler) *Client {
	c := &Client{
		cfg:       cfg,
		logger:    logger,
		handler:   handler,
		session:   NewSession(),
		closeChan: make(chan struct{}),
	}

	// Create heartbeat manager (will be started after session opens)
	c.heartbeat = NewHeartbeat(
		cfg.HeartbeatInterval,
		c.sendHeartbeat,
		c.onHeartbeatFailure,
		logger.With("component", "heartbeat"),
	)

	return c
}

// Run connects to the CTI server and processes messages until the context is canceled.
func (c *Client) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			c.logger.Info("context canceled, shutting down")
			c.Close()
			return ctx.Err()
		default:
		}

		if err := c.connect(ctx); err != nil {
			c.logger.Error("connection failed", "error", err)
			c.waitForRetry(ctx)
			continue
		}

		if err := c.open(ctx); err != nil {
			c.logger.Error("failed to open session", "error", err)
			c.disconnect()
			c.waitForRetry(ctx)
			continue
		}

		// Start heartbeat
		heartbeatCtx, cancelHeartbeat := context.WithCancel(ctx)
		heartbeatDone := make(chan struct{})
		go func() {
			defer close(heartbeatDone)
			c.heartbeat.Run(heartbeatCtx)
		}()

		// Process messages until error or context canceled
		err := c.processMessages(ctx)
		cancelHeartbeat()
		<-heartbeatDone // wait for the goroutine to finish before reconnecting

		if err != nil {
			c.logger.Error("message processing error", "error", err)
		}

		c.disconnect()

		// Check if we should retry
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			c.waitForRetry(ctx)
		}
	}
}

// connect establishes a TCP connection to the CTI server.
func (c *Client) connect(ctx context.Context) error {
	c.session.SetState(StateConnecting)

	addr := fmt.Sprintf("%s:%d", c.cfg.ServerHost, c.cfg.ServerPort)
	c.logger.Info("connecting to CTI server", "address", addr)

	dialer := net.Dialer{
		Timeout: 30 * time.Second,
	}

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		c.session.SetState(StateDisconnected)
		return fmt.Errorf("failed to connect: %w", err)
	}

	c.mu.Lock()
	c.conn = conn
	c.reader = NewReader(conn)
	c.mu.Unlock()

	c.session.SetState(StateConnected)
	c.logger.Info("connected to CTI server")

	return nil
}

// open sends OPEN_REQ and waits for OPEN_CONF.
func (c *Client) open(ctx context.Context) error {
	c.session.SetState(StateOpening)

	openReq := &messages.OpenReq{
		InvokeID:          c.session.NextInvokeID(),
		VersionNumber:     24, // Protocol version
		IdleTimeout:       uint32(c.cfg.IdleTimeout.Seconds()),
		PeripheralID:      c.cfg.PeripheralID,
		ServicesRequested: c.cfg.ServicesRequested,
		CallMsgMask:       0xFFFFFFFF, // All call events
		AgentStateMask:    0xFFFFFFFF, // All agent state events
		ConfigMsgMask:     0xFFFFFFFF, // All config events
		ClientID:          c.cfg.ClientID,
	}

	if err := c.sendMessage(openReq); err != nil {
		return fmt.Errorf("failed to send OPEN_REQ: %w", err)
	}

	c.logger.Info("sent OPEN_REQ", "invokeID", openReq.InvokeID)

	// Wait for OPEN_CONF with timeout
	deadline := time.Now().Add(30 * time.Second)
	c.conn.SetReadDeadline(deadline)
	defer c.conn.SetReadDeadline(time.Time{})

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		msg, err := c.reader.ReadMessage()
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}

		switch m := msg.(type) {
		case *messages.OpenConf:
			c.session.SetMonitorID(m.MonitorID)
			c.session.SetOpenConfDetails(m.ServicesGranted, m.FltPeripheralID, m.AgentState)
			c.session.SetState(StateOpen)
			c.logger.Info("session opened",
				"monitorID", m.MonitorID,
				"servicesGranted", m.ServicesGranted,
				"peripheralID", m.FltPeripheralID,
				"agentState", protocol.AgentStateName(m.AgentState))
			return nil

		case *messages.FailureConf:
			return fmt.Errorf("OPEN_REQ rejected with status %d", m.Status)

		case *messages.FailureEvent:
			return fmt.Errorf("failure event received with status %d", m.Status)

		default:
			c.logger.Warn("unexpected message during open",
				"type", protocol.MessageTypeName(msg.Type()))
		}
	}
}

// processMessages reads and handles messages from the server.
func (c *Client) processMessages(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.closeChan:
			return nil
		default:
		}

		// Set a read deadline to allow periodic context checks
		c.conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		msg, err := c.reader.ReadMessage()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue // Timeout is expected, check context and continue
			}
			return err
		}

		c.handleMessage(msg)
	}
}

// handleMessage processes a received message.
func (c *Client) handleMessage(msg protocol.Message) {
	msgType := msg.Type()
	c.logger.Debug("received message", "type", protocol.MessageTypeName(msgType))

	switch m := msg.(type) {
	case *messages.HeartbeatConf:
		c.heartbeat.Confirm(m.InvokeID)

	case *messages.CloseConf:
		c.logger.Info("received CLOSE_CONF")
		c.session.SetState(StateDisconnected)

	case *messages.FailureConf:
		c.logger.Error("received FAILURE_CONF", "status", m.Status, "invokeID", m.InvokeID)

	case *messages.FailureEvent:
		c.logger.Error("received FAILURE_EVENT", "status", m.Status)

	case *messages.SystemEvent:
		c.logger.Info("received SYSTEM_EVENT",
			"eventID", m.SystemEventID,
			"eventName", m.EventName())
		if c.handler != nil {
			c.handler(msg)
		}

	default:
		// Pass all other messages to the handler
		if c.handler != nil {
			c.handler(msg)
		}
	}
}

// sendMessage encodes and sends a message to the server.
func (c *Client) sendMessage(msg protocol.Message) error {
	data, err := protocol.EncodeMessage(msg)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return fmt.Errorf("not connected")
	}

	_, err = c.conn.Write(data)
	return err
}

// sendHeartbeat sends a heartbeat request using the shared session invokeID.
func (c *Client) sendHeartbeat() (uint32, error) {
	invokeID := c.session.NextInvokeID()
	return invokeID, c.sendMessage(&messages.HeartbeatReq{InvokeID: invokeID})
}

// onHeartbeatFailure is called when heartbeat fails.
func (c *Client) onHeartbeatFailure() {
	c.logger.Error("heartbeat failure, triggering reconnect")
	c.disconnect()
}

// disconnect closes the connection.
func (c *Client) disconnect() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	c.reader = nil
	c.session.Reset()
}

// Close gracefully closes the session.
func (c *Client) Close() error {
	if !c.session.IsOpen() {
		c.disconnect()
		return nil
	}

	c.session.SetState(StateClosing)
	c.logger.Info("closing session")

	closeReq := &messages.CloseReq{
		InvokeID: c.session.NextInvokeID(),
		Status:   0,
	}

	if err := c.sendMessage(closeReq); err != nil {
		c.logger.Warn("failed to send CLOSE_REQ", "error", err)
	}

	// Wait briefly for CLOSE_CONF
	time.Sleep(500 * time.Millisecond)

	close(c.closeChan)
	c.disconnect()
	return nil
}

// waitForRetry waits before retrying a connection.
func (c *Client) waitForRetry(ctx context.Context) {
	c.logger.Info("waiting before retry", "delay", c.cfg.ReconnectDelay)
	select {
	case <-ctx.Done():
	case <-time.After(c.cfg.ReconnectDelay):
	}
}

// State returns the current session state.
func (c *Client) State() SessionState {
	return c.session.State()
}
