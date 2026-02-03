// Package messages defines all CTI protocol message types.
package messages

import (
	"ctiservice/internal/protocol"
)

// OpenReq is sent to initialize a session with the CTI server.
// Protocol Version 24 - OPEN_REQ (MessageType = 3)
type OpenReq struct {
	// Fixed Part (44 bytes, excluding 8-byte header)
	InvokeID          uint32 // Client-assigned ID returned in response
	VersionNumber     uint32 // Protocol version number (24)
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
	ClientID          string // Tag 1 - Client identifier (max 64 bytes)
	ClientPassword    string // Tag 2 - Client password (max 64 bytes)
	ClientSignature   string // Tag 28 - Client signature (max 64 bytes)
	AgentExtension    string // Tag 3 - Agent's extension (max 16 bytes)
	AgentID           string // Tag 4 - Agent's ID (max 12 bytes)
	AgentInstrument   string // Tag 5 - Agent's instrument (max 64 bytes)
	ApplicationPathID int32  // Tag 90 - Application path ID
}

func (m *OpenReq) Type() uint32 {
	return protocol.MsgTypeOpenReq
}

func (m *OpenReq) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()

	// Fixed part (44 bytes)
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
	if m.ClientSignature != "" {
		fw.WriteString(protocol.TagCTIClientSignature, m.ClientSignature)
	}
	if m.AgentExtension != "" {
		fw.WriteString(protocol.TagAgentExtension, m.AgentExtension)
	}
	if m.AgentID != "" {
		fw.WriteString(protocol.TagAgentID, m.AgentID)
	}
	if m.AgentInstrument != "" {
		fw.WriteString(protocol.TagAgentInstrument, m.AgentInstrument)
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
		m.ClientSignature = ff.GetString(protocol.TagCTIClientSignature)
		m.AgentExtension = ff.GetString(protocol.TagAgentExtension)
		m.AgentID = ff.GetString(protocol.TagAgentID)
		m.AgentInstrument = ff.GetString(protocol.TagAgentInstrument)
	}

	return nil
}

// OpenConf is the server's response to OpenReq.
// Protocol Version 24 - OPEN_CONF (MessageType = 4)
type OpenConf struct {
	// Fixed Part (28 bytes, excluding 8-byte header)
	InvokeID                 uint32 // Matches InvokeID from OpenReq
	ServicesGranted          uint32 // Bitmask of granted services
	MonitorID                uint32 // Monitor ID assigned by server
	PGStatus                 uint32 // Peripheral gateway status
	ICMCentralControllerTime uint32 // ICM central controller time (TIME)
	PeripheralOnline         bool   // Peripheral online status (BOOL - 2 bytes)
	PeripheralType           uint16 // Type of peripheral (USHORT)
	AgentState               uint16 // Current agent state (USHORT)
	DepartmentID             int32  // Department ID (INT)
	SessionType              uint16 // Session type (USHORT)

	// Floating fields
	AgentExtension        string // Tag 3
	AgentID               string // Tag 4
	AgentInstrument       string // Tag 5
	NumPeripherals        uint16 // Tag 232
	FltPeripheralID       uint32 // Tag 6
	MultilineAgentControl uint16 // Tag 180
}

func (m *OpenConf) Type() uint32 {
	return protocol.MsgTypeOpenConf
}

func (m *OpenConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()

	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.ServicesGranted)
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PGStatus)
	w.WriteUint32(m.ICMCentralControllerTime)
	w.WriteBool(m.PeripheralOnline)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.AgentState)
	w.WriteInt32(m.DepartmentID)
	w.WriteUint16(m.SessionType)

	return w.Bytes(), w.Error()
}

func (m *OpenConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)

	m.InvokeID = r.ReadUint32()
	m.ServicesGranted = r.ReadUint32()
	m.MonitorID = r.ReadUint32()
	m.PGStatus = r.ReadUint32()
	m.ICMCentralControllerTime = r.ReadUint32()
	m.PeripheralOnline = r.ReadBool()
	m.PeripheralType = r.ReadUint16()
	m.AgentState = r.ReadUint16()
	m.DepartmentID = r.ReadInt32()
	m.SessionType = r.ReadUint16()

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
		m.NumPeripherals = ff.GetUint16(protocol.TagNumPeripherals)
		m.FltPeripheralID = ff.GetUint32(protocol.TagPeripheralID)
		m.MultilineAgentControl = ff.GetUint16(protocol.TagMultilineAgentControl)
	}

	return nil
}

// HeartbeatReq is sent by the client to maintain the connection.
// Protocol Version 24 - HEARTBEAT_REQ (MessageType = 5)
type HeartbeatReq struct {
	InvokeID uint32 // 4 bytes
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
// Protocol Version 24 - HEARTBEAT_CONF (MessageType = 6)
type HeartbeatConf struct {
	InvokeID uint32 // 4 bytes
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
// Protocol Version 24 - CLOSE_REQ (MessageType = 7)
type CloseReq struct {
	InvokeID uint32 // 4 bytes
	Status   uint32 // 4 bytes
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
// Protocol Version 24 - CLOSE_CONF (MessageType = 8)
type CloseConf struct {
	InvokeID uint32 // 4 bytes
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
