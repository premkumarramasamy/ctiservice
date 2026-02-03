// Package messages defines all CTI protocol message types.
package messages

import (
	"ctiservice/internal/protocol"
)

// OpenReq is sent to initialize a session with the CTI server.
type OpenReq struct {
	InvokeID          uint32 // Client-assigned ID returned in response
	VersionNumber     uint32 // Protocol version number
	IdleTimeout       uint32 // Seconds of inactivity before server closes session
	PeripheralID      uint32 // Peripheral to connect to (0 for any)
	ServicesRequested uint32 // Bitmask of requested services
	CallMsgMask       uint32 // Call events to receive
	AgentStateMask    uint32 // Agent state events to receive
	ConfigMsgMask     uint32 // Config events to receive
	Reserved1         uint32 // Reserved
	Reserved2         uint32 // Reserved
	Reserved3         uint32 // Reserved

	// Floating fields
	ClientID       string // Client identifier
	ClientPassword string // Client password (if required)
}

func (m *OpenReq) Type() uint32 {
	return protocol.MsgTypeOpenReq
}

func (m *OpenReq) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()

	// Fixed part
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.VersionNumber)
	w.WriteUint32(m.IdleTimeout)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint32(m.ServicesRequested)
	w.WriteUint32(m.CallMsgMask)
	w.WriteUint32(m.AgentStateMask)
	w.WriteUint32(m.ConfigMsgMask)
	w.WriteUint32(m.Reserved1)
	w.WriteUint32(m.Reserved2)
	w.WriteUint32(m.Reserved3)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Floating part
	fw := protocol.NewFloatingFieldWriter()
	if m.ClientID != "" {
		fw.WriteString(protocol.TagClientID, m.ClientID)
	}
	if m.ClientPassword != "" {
		fw.WriteString(protocol.TagClientPassword, m.ClientPassword)
	}

	// Combine fixed and floating parts
	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *OpenReq) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)

	m.InvokeID = r.ReadUint32()
	m.VersionNumber = r.ReadUint32()
	m.IdleTimeout = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.ServicesRequested = r.ReadUint32()
	m.CallMsgMask = r.ReadUint32()
	m.AgentStateMask = r.ReadUint32()
	m.ConfigMsgMask = r.ReadUint32()
	m.Reserved1 = r.ReadUint32()
	m.Reserved2 = r.ReadUint32()
	m.Reserved3 = r.ReadUint32()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ClientID = ff.GetString(protocol.TagClientID)
		m.ClientPassword = ff.GetString(protocol.TagClientPassword)
	}

	return nil
}

// OpenConf is the server's response to OpenReq.
type OpenConf struct {
	InvokeID            uint32 // Matches InvokeID from OpenReq
	ServiceGranted      uint32 // Bitmask of granted services
	MonitorID           uint32 // Monitor ID assigned by server
	PGStatus            uint32 // Peripheral gateway status
	ICMCentralController uint32 // ICM central controller status
	AgentState          uint16 // Current agent state
	NumPeripherals      uint16 // Number of connected peripherals
	MultiLineAgentControl uint16 // Multi-line agent control flag
	Reserved1           uint16 // Reserved
	PeripheralID        uint32 // Peripheral ID
	AgentExtension      string // Agent's extension (floating)
	AgentID             string // Agent's ID (floating)
	AgentInstrument     string // Agent's instrument (floating)
}

func (m *OpenConf) Type() uint32 {
	return protocol.MsgTypeOpenConf
}

func (m *OpenConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()

	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.ServiceGranted)
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PGStatus)
	w.WriteUint32(m.ICMCentralController)
	w.WriteUint16(m.AgentState)
	w.WriteUint16(m.NumPeripherals)
	w.WriteUint16(m.MultiLineAgentControl)
	w.WriteUint16(m.Reserved1)
	w.WriteUint32(m.PeripheralID)

	return w.Bytes(), w.Error()
}

func (m *OpenConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)

	m.InvokeID = r.ReadUint32()
	m.ServiceGranted = r.ReadUint32()
	m.MonitorID = r.ReadUint32()
	m.PGStatus = r.ReadUint32()
	m.ICMCentralController = r.ReadUint32()
	m.AgentState = r.ReadUint16()
	m.NumPeripherals = r.ReadUint16()
	m.MultiLineAgentControl = r.ReadUint16()
	m.Reserved1 = r.ReadUint16()
	m.PeripheralID = r.ReadUint32()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.AgentExtension = ff.GetString(protocol.TagAgentExtension)
		m.AgentID = ff.GetString(protocol.TagAgentID)
		m.AgentInstrument = ff.GetString(protocol.TagAgentInstrument)
	}

	return nil
}

// HeartbeatReq is sent by the client to maintain the connection.
type HeartbeatReq struct {
	InvokeID uint32
}

func (m *HeartbeatReq) Type() uint32 {
	return protocol.MsgTypeHeartbeatReq
}

func (m *HeartbeatReq) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	return w.Bytes(), w.Error()
}

func (m *HeartbeatReq) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	return r.Error()
}

// HeartbeatConf is the server's response to HeartbeatReq.
type HeartbeatConf struct {
	InvokeID uint32
}

func (m *HeartbeatConf) Type() uint32 {
	return protocol.MsgTypeHeartbeatConf
}

func (m *HeartbeatConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	return w.Bytes(), w.Error()
}

func (m *HeartbeatConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	return r.Error()
}

// CloseReq is sent to gracefully close a session.
type CloseReq struct {
	InvokeID uint32
	Status   uint32
}

func (m *CloseReq) Type() uint32 {
	return protocol.MsgTypeCloseReq
}

func (m *CloseReq) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.Status)
	return w.Bytes(), w.Error()
}

func (m *CloseReq) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.Status = r.ReadUint32()
	return r.Error()
}

// CloseConf is the server's response to CloseReq.
type CloseConf struct {
	InvokeID uint32
}

func (m *CloseConf) Type() uint32 {
	return protocol.MsgTypeCloseConf
}

func (m *CloseConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	return w.Bytes(), w.Error()
}

func (m *CloseConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	return r.Error()
}
