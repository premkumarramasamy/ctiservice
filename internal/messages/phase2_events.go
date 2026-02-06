package messages

import (
	"ctiservice/internal/protocol"
)

// CallServiceInitiatedEvent is sent when telecommunications service (dial tone)
// is initiated at the agent's teleset.
// Protocol Version 24 - CALL_SERVICE_INITIATED_EVENT (MessageType = 20)
type CallServiceInitiatedEvent struct {
	// Fixed Part
	MonitorID              uint32 // Monitor ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	PeripheralType         uint16 // Peripheral type (USHORT)
	ConnectionDeviceIDType uint16 // Device ID type (USHORT)
	ConnectionCallID       uint32 // Call ID (UINT)
	LineHandle             uint16 // Line handle (USHORT)
	LineType               uint16 // Line type (USHORT)
	ServiceNumber          uint32 // Service number (UINT)
	ServiceID              uint32 // Service ID (UINT)
	SkillGroupNumber       uint32 // Skill group number (UINT)
	SkillGroupID           uint32 // Skill group ID (UINT)
	SkillGroupPriority     uint16 // Skill group priority (USHORT)
	CallingDeviceType      uint16 // Calling device type (USHORT)
	LocalConnectionState   uint16 // Local connection state (USHORT)
	EventCause             uint16 // Event cause (USHORT)

	// Floating fields
	ConnectionDeviceID string // Tag 31 - Connection device ID
	CallingDeviceID    string // Tag 12 - Calling device ID (optional)
	CallReferenceID    string // Tag 248 - Call reference ID (optional)
}

func (m *CallServiceInitiatedEvent) Type() uint32 {
	return protocol.MsgTypeCallServiceInitiatedEvent
}

func (m *CallServiceInitiatedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.LineHandle)
	w.WriteUint16(m.LineType)
	w.WriteUint32(m.ServiceNumber)
	w.WriteUint32(m.ServiceID)
	w.WriteUint32(m.SkillGroupNumber)
	w.WriteUint32(m.SkillGroupID)
	w.WriteUint16(m.SkillGroupPriority)
	w.WriteUint16(m.CallingDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)

	if err := w.Error(); err != nil {
		return nil, err
	}

	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}
	if m.CallingDeviceID != "" {
		fw.WriteString(protocol.TagCallingDeviceID, m.CallingDeviceID)
	}
	if m.CallReferenceID != "" {
		fw.WriteString(protocol.TagCallReferenceID, m.CallReferenceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *CallServiceInitiatedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.LineHandle = r.ReadUint16()
	m.LineType = r.ReadUint16()
	m.ServiceNumber = r.ReadUint32()
	m.ServiceID = r.ReadUint32()
	m.SkillGroupNumber = r.ReadUint32()
	m.SkillGroupID = r.ReadUint32()
	m.SkillGroupPriority = r.ReadUint16()
	m.CallingDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.CallingDeviceID = ff.GetString(protocol.TagCallingDeviceID)
		m.CallReferenceID = ff.GetString(protocol.TagCallReferenceID)
	}

	return nil
}

// AgentPreCallEvent is sent when a call is routed to an Enterprise Agent.
// This provides advance notification before the call arrives.
// Protocol Version 24 - AGENT_PRE_CALL_EVENT (MessageType = 87)
type AgentPreCallEvent struct {
	// Fixed Part
	MonitorID              uint32 // Monitor ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	PeripheralType         uint16 // Peripheral type (USHORT)
	ConnectionDeviceIDType uint16 // Device ID type (USHORT)
	ConnectionCallID       uint32 // Call ID (UINT)
	ServiceNumber          uint32 // Service number (UINT)
	ServiceID              uint32 // Service ID (UINT)
	SkillGroupNumber       uint32 // Skill group number (UINT)
	SkillGroupID           uint32 // Skill group ID (UINT)
	SkillGroupPriority     uint16 // Skill group priority (USHORT)
	NumCTIClients          uint16 // Number of CTI clients (USHORT)
	NumNamedVariables      uint16 // Number of named variables (USHORT)
	NumNamedArrays         uint16 // Number of named arrays (USHORT)
	CallType               uint16 // Call type (USHORT)

	// Floating fields
	ConnectionDeviceID  string // Tag 31 - Connection device ID
	ANI                 string // Tag 15 - Automatic Number Identification
	DNIS                string // Tag 16 - Dialed Number Identification Service
	DialedNumber        string // Tag 40 - Dialed number
	CallerEnteredDigits string // Tag 41 - Caller entered digits
	UserToUserInfo      string // Tag 17 - User to user info
	CallVariable1       string // Tag 18
	CallVariable2       string // Tag 19
	CallVariable3       string // Tag 20
	CallVariable4       string // Tag 21
	CallVariable5       string // Tag 22
	CallVariable6       string // Tag 23
	CallVariable7       string // Tag 24
	CallVariable8       string // Tag 25
	CallVariable9       string // Tag 26
	CallVariable10      string // Tag 27
	CallTypeID          uint32 // Tag 250 - Call type ID
	PreCallInvokeID     uint32 // Tag 249 - Pre-call invoke ID
	RouterCallKeyDay    uint32 // Tag 72
	RouterCallKeyCallID uint32 // Tag 73
	RouterCallKeySeqNum uint32 // Tag 214
}

func (m *AgentPreCallEvent) Type() uint32 {
	return protocol.MsgTypeAgentPreCallEvent
}

func (m *AgentPreCallEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint32(m.ServiceNumber)
	w.WriteUint32(m.ServiceID)
	w.WriteUint32(m.SkillGroupNumber)
	w.WriteUint32(m.SkillGroupID)
	w.WriteUint16(m.SkillGroupPriority)
	w.WriteUint16(m.NumCTIClients)
	w.WriteUint16(m.NumNamedVariables)
	w.WriteUint16(m.NumNamedArrays)
	w.WriteUint16(m.CallType)

	if err := w.Error(); err != nil {
		return nil, err
	}

	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}
	if m.ANI != "" {
		fw.WriteString(protocol.TagANI, m.ANI)
	}
	if m.DNIS != "" {
		fw.WriteString(protocol.TagDNIS, m.DNIS)
	}
	if m.DialedNumber != "" {
		fw.WriteString(protocol.TagDialedNumber, m.DialedNumber)
	}
	if m.CallerEnteredDigits != "" {
		fw.WriteString(protocol.TagCallerEnteredDigits, m.CallerEnteredDigits)
	}
	if m.UserToUserInfo != "" {
		fw.WriteString(protocol.TagUserToUserInfo, m.UserToUserInfo)
	}
	if m.CallVariable1 != "" {
		fw.WriteString(protocol.TagCallVariable1, m.CallVariable1)
	}
	if m.CallVariable2 != "" {
		fw.WriteString(protocol.TagCallVariable2, m.CallVariable2)
	}
	if m.CallVariable3 != "" {
		fw.WriteString(protocol.TagCallVariable3, m.CallVariable3)
	}
	if m.CallVariable4 != "" {
		fw.WriteString(protocol.TagCallVariable4, m.CallVariable4)
	}
	if m.CallVariable5 != "" {
		fw.WriteString(protocol.TagCallVariable5, m.CallVariable5)
	}
	if m.CallVariable6 != "" {
		fw.WriteString(protocol.TagCallVariable6, m.CallVariable6)
	}
	if m.CallVariable7 != "" {
		fw.WriteString(protocol.TagCallVariable7, m.CallVariable7)
	}
	if m.CallVariable8 != "" {
		fw.WriteString(protocol.TagCallVariable8, m.CallVariable8)
	}
	if m.CallVariable9 != "" {
		fw.WriteString(protocol.TagCallVariable9, m.CallVariable9)
	}
	if m.CallVariable10 != "" {
		fw.WriteString(protocol.TagCallVariable10, m.CallVariable10)
	}
	if m.CallTypeID != 0 {
		fw.WriteUint32(protocol.TagCallTypeID, m.CallTypeID)
	}
	if m.PreCallInvokeID != 0 {
		fw.WriteUint32(protocol.TagPreCallInvokeID, m.PreCallInvokeID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *AgentPreCallEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.ServiceNumber = r.ReadUint32()
	m.ServiceID = r.ReadUint32()
	m.SkillGroupNumber = r.ReadUint32()
	m.SkillGroupID = r.ReadUint32()
	m.SkillGroupPriority = r.ReadUint16()
	m.NumCTIClients = r.ReadUint16()
	m.NumNamedVariables = r.ReadUint16()
	m.NumNamedArrays = r.ReadUint16()
	m.CallType = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.ANI = ff.GetString(protocol.TagANI)
		m.DNIS = ff.GetString(protocol.TagDNIS)
		m.DialedNumber = ff.GetString(protocol.TagDialedNumber)
		m.CallerEnteredDigits = ff.GetString(protocol.TagCallerEnteredDigits)
		m.UserToUserInfo = ff.GetString(protocol.TagUserToUserInfo)
		m.CallVariable1 = ff.GetString(protocol.TagCallVariable1)
		m.CallVariable2 = ff.GetString(protocol.TagCallVariable2)
		m.CallVariable3 = ff.GetString(protocol.TagCallVariable3)
		m.CallVariable4 = ff.GetString(protocol.TagCallVariable4)
		m.CallVariable5 = ff.GetString(protocol.TagCallVariable5)
		m.CallVariable6 = ff.GetString(protocol.TagCallVariable6)
		m.CallVariable7 = ff.GetString(protocol.TagCallVariable7)
		m.CallVariable8 = ff.GetString(protocol.TagCallVariable8)
		m.CallVariable9 = ff.GetString(protocol.TagCallVariable9)
		m.CallVariable10 = ff.GetString(protocol.TagCallVariable10)
		m.CallTypeID = ff.GetUint32(protocol.TagCallTypeID)
		m.PreCallInvokeID = ff.GetUint32(protocol.TagPreCallInvokeID)
		m.RouterCallKeyDay = ff.GetUint32(protocol.TagRouterCallKeyDay)
		m.RouterCallKeyCallID = ff.GetUint32(protocol.TagRouterCallKeyCallID)
		m.RouterCallKeySeqNum = ff.GetUint32(protocol.TagRouterCallKeySeqNum)
	}

	return nil
}

// AgentPreCallAbortEvent is sent when a call that was previously announced
// through an AGENT_PRE_CALL_EVENT cannot be routed as intended.
// Protocol Version 24 - AGENT_PRE_CALL_ABORT_EVENT (MessageType = 88)
type AgentPreCallAbortEvent struct {
	// Fixed Part
	MonitorID              uint32 // Monitor ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	PeripheralType         uint16 // Peripheral type (USHORT)
	ConnectionDeviceIDType uint16 // Device ID type (USHORT)
	ConnectionCallID       uint32 // Call ID (UINT)
	EventCause             uint16 // Event cause (USHORT)

	// Floating fields
	ConnectionDeviceID string // Tag 31 - Connection device ID
	PreCallInvokeID    uint32 // Tag 249 - Pre-call invoke ID (matches the original event)
}

func (m *AgentPreCallAbortEvent) Type() uint32 {
	return protocol.MsgTypeAgentPreCallAbortEvent
}

func (m *AgentPreCallAbortEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.EventCause)

	if err := w.Error(); err != nil {
		return nil, err
	}

	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}
	if m.PreCallInvokeID != 0 {
		fw.WriteUint32(protocol.TagPreCallInvokeID, m.PreCallInvokeID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *AgentPreCallAbortEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.EventCause = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.PreCallInvokeID = ff.GetUint32(protocol.TagPreCallInvokeID)
	}

	return nil
}

// SupervisorAction represents the type of supervisor action in SUPERVISOR_ASSIST_EVENT.
type SupervisorAction uint16

const (
	SupervisorActionNone      SupervisorAction = 0
	SupervisorActionMonitor   SupervisorAction = 1 // Silent monitoring
	SupervisorActionCoach     SupervisorAction = 2 // Coach (whisper to agent)
	SupervisorActionBarge     SupervisorAction = 3 // Barge into call
	SupervisorActionIntercept SupervisorAction = 4 // Intercept call
)

// SupervisorActionName returns a human-readable name for a supervisor action.
func SupervisorActionName(action SupervisorAction) string {
	switch action {
	case SupervisorActionNone:
		return "None"
	case SupervisorActionMonitor:
		return "Monitor"
	case SupervisorActionCoach:
		return "Coach"
	case SupervisorActionBarge:
		return "Barge"
	case SupervisorActionIntercept:
		return "Intercept"
	default:
		return "Unknown"
	}
}

// SupervisorAssistEvent is sent when a supervisor performs an action on an agent's call.
// This includes silent monitoring, coaching (whisper), barging, and intercepting.
// Protocol Version 24 - SUPERVISOR_ASSIST_EVENT (MessageType = 120)
type SupervisorAssistEvent struct {
	// Fixed Part
	MonitorID              uint32 // Monitor ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	PeripheralType         uint16 // Peripheral type (USHORT)
	ConnectionDeviceIDType uint16 // Device ID type (USHORT)
	ConnectionCallID       uint32 // Call ID (UINT)
	SupervisorAction       uint16 // Supervisor action type (USHORT)
	EventCause             uint16 // Event cause (USHORT)

	// Floating fields
	ConnectionDeviceID       string // Tag 31 - Connection device ID
	AgentConnectionDeviceID  string // Tag 31 - Agent's connection device ID
	AgentConnectionCallID    uint32 // Tag 193 - Agent's connection call ID
	AgentPeripheralID        uint32 // Tag 194 - Agent's peripheral ID
	AgentPeripheralNumber    uint32 // Tag 195 - Agent's peripheral number
	AgentID                  string // Tag 4 - Agent ID
	AgentExtension           string // Tag 3 - Agent extension
}

func (m *SupervisorAssistEvent) Type() uint32 {
	return protocol.MsgTypeSupervisorAssistEvent
}

func (m *SupervisorAssistEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.SupervisorAction)
	w.WriteUint16(m.EventCause)

	if err := w.Error(); err != nil {
		return nil, err
	}

	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}
	if m.AgentID != "" {
		fw.WriteString(protocol.TagAgentID, m.AgentID)
	}
	if m.AgentExtension != "" {
		fw.WriteString(protocol.TagAgentExtension, m.AgentExtension)
	}
	if m.AgentConnectionCallID != 0 {
		fw.WriteUint32(protocol.TagAgentConnectionCallID, m.AgentConnectionCallID)
	}
	if m.AgentPeripheralID != 0 {
		fw.WriteUint32(protocol.TagAgentPeripheralID, m.AgentPeripheralID)
	}
	if m.AgentPeripheralNumber != 0 {
		fw.WriteUint32(protocol.TagAgentPeripheralNumber, m.AgentPeripheralNumber)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *SupervisorAssistEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.SupervisorAction = r.ReadUint16()
	m.EventCause = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.AgentID = ff.GetString(protocol.TagAgentID)
		m.AgentExtension = ff.GetString(protocol.TagAgentExtension)
		m.AgentConnectionCallID = ff.GetUint32(protocol.TagAgentConnectionCallID)
		m.AgentPeripheralID = ff.GetUint32(protocol.TagAgentPeripheralID)
		m.AgentPeripheralNumber = ff.GetUint32(protocol.TagAgentPeripheralNumber)
	}

	return nil
}

// ActionName returns a human-readable name for the supervisor action.
func (m *SupervisorAssistEvent) ActionName() string {
	return SupervisorActionName(SupervisorAction(m.SupervisorAction))
}
