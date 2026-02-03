package client

import (
	"sync"
)

// SessionState represents the current state of the CTI session.
type SessionState int

const (
	StateDisconnected SessionState = iota
	StateConnecting
	StateConnected
	StateOpening
	StateOpen
	StateClosing
)

// String returns the string representation of the session state.
func (s SessionState) String() string {
	switch s {
	case StateDisconnected:
		return "Disconnected"
	case StateConnecting:
		return "Connecting"
	case StateConnected:
		return "Connected"
	case StateOpening:
		return "Opening"
	case StateOpen:
		return "Open"
	case StateClosing:
		return "Closing"
	default:
		return "Unknown"
	}
}

// Session tracks the state of a CTI session.
type Session struct {
	mu        sync.RWMutex
	state     SessionState
	monitorID uint32
	invokeID  uint32

	// Session details from OPEN_CONF
	serviceGranted uint32
	peripheralID   uint32
	agentState     uint16
}

// NewSession creates a new session tracker.
func NewSession() *Session {
	return &Session{
		state: StateDisconnected,
	}
}

// State returns the current session state.
func (s *Session) State() SessionState {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.state
}

// SetState sets the session state.
func (s *Session) SetState(state SessionState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.state = state
}

// IsOpen returns true if the session is open and ready.
func (s *Session) IsOpen() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.state == StateOpen
}

// MonitorID returns the monitor ID assigned by the server.
func (s *Session) MonitorID() uint32 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.monitorID
}

// SetMonitorID sets the monitor ID.
func (s *Session) SetMonitorID(id uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.monitorID = id
}

// NextInvokeID returns the next invoke ID and increments the counter.
func (s *Session) NextInvokeID() uint32 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.invokeID++
	return s.invokeID
}

// ServiceGranted returns the services granted by the server.
func (s *Session) ServiceGranted() uint32 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.serviceGranted
}

// SetOpenConfDetails stores details from the OPEN_CONF response.
func (s *Session) SetOpenConfDetails(serviceGranted, peripheralID uint32, agentState uint16) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.serviceGranted = serviceGranted
	s.peripheralID = peripheralID
	s.agentState = agentState
}

// PeripheralID returns the peripheral ID.
func (s *Session) PeripheralID() uint32 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.peripheralID
}

// AgentState returns the agent state from OPEN_CONF.
func (s *Session) AgentState() uint16 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.agentState
}

// Reset resets the session to disconnected state.
func (s *Session) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.state = StateDisconnected
	s.monitorID = 0
	s.serviceGranted = 0
	s.peripheralID = 0
	s.agentState = 0
	// Don't reset invokeID - keep incrementing
}
