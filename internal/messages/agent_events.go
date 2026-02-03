package messages

import (
	"ctiservice/internal/protocol"
)

// AgentStateEvent reports an agent's state change.
// Protocol Version 24 - AGENT_STATE_EVENT (MessageType = 30)
type AgentStateEvent struct {
	// Fixed Part (60 bytes)
	MonitorID               uint32 // Monitor ID (UINT)
	PeripheralID            uint32 // Peripheral ID (UINT)
	SessionID               uint32 // Session ID (UINT)
	PeripheralType          uint16 // Peripheral type (USHORT)
	SkillGroupState         uint16 // Skill group state (USHORT)
	StateDuration           uint32 // Duration in current state (UINT, seconds)
	SkillGroupNumber        uint32 // Skill group number (UINT)
	SkillGroupID            uint32 // Skill group ID (UINT)
	SkillGroupPriority      uint16 // Skill group priority (USHORT)
	AgentState              uint16 // Current agent state (USHORT)
	EventReasonCode         uint16 // Reason code for state change (USHORT)
	MRDID                   int32  // Media routing domain ID (INT)
	NumTasks                uint32 // Number of active tasks (UINT)
	AgentMode               uint16 // Agent mode (USHORT)
	MaxTaskLimit            uint32 // Maximum task limit (UINT)
	ICMAgentID              int32  // ICM agent ID (INT)
	AgentAvailabilityStatus uint32 // Availability status (UINT)
	NumFltSkillGroups       uint16 // Number of floating skill groups (USHORT)
	DepartmentID            int32  // Department ID (INT)

	// Floating fields
	CTIClientSignature string // Tag 28
	AgentID            string // Tag 4 (max 12 bytes)
	AgentExtension     string // Tag 3 (max 16 bytes)
	ActiveTerminal     string // Tag 127 (max 64 bytes)
	AgentInstrument    string // Tag 5 (max 64 bytes)
	Duration           uint32 // Tag 126
	NextAgentState     uint16 // Tag 123
	Direction          uint32 // Tag 128
}

func (m *AgentStateEvent) Type() uint32 {
	return protocol.MsgTypeAgentStateEvent
}

func (m *AgentStateEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint32(m.SessionID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.SkillGroupState)
	w.WriteUint32(m.StateDuration)
	w.WriteUint32(m.SkillGroupNumber)
	w.WriteUint32(m.SkillGroupID)
	w.WriteUint16(m.SkillGroupPriority)
	w.WriteUint16(m.AgentState)
	w.WriteUint16(m.EventReasonCode)
	w.WriteInt32(m.MRDID)
	w.WriteUint32(m.NumTasks)
	w.WriteUint16(m.AgentMode)
	w.WriteUint32(m.MaxTaskLimit)
	w.WriteInt32(m.ICMAgentID)
	w.WriteUint32(m.AgentAvailabilityStatus)
	w.WriteUint16(m.NumFltSkillGroups)
	w.WriteInt32(m.DepartmentID)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.CTIClientSignature != "" {
		fw.WriteString(protocol.TagCTIClientSignature, m.CTIClientSignature)
	}
	if m.AgentID != "" {
		fw.WriteString(protocol.TagAgentID, m.AgentID)
	}
	if m.AgentExtension != "" {
		fw.WriteString(protocol.TagAgentExtension, m.AgentExtension)
	}
	if m.ActiveTerminal != "" {
		fw.WriteString(protocol.TagActiveTerminal, m.ActiveTerminal)
	}
	if m.AgentInstrument != "" {
		fw.WriteString(protocol.TagAgentInstrument, m.AgentInstrument)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *AgentStateEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)

	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.SessionID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.SkillGroupState = r.ReadUint16()
	m.StateDuration = r.ReadUint32()
	m.SkillGroupNumber = r.ReadUint32()
	m.SkillGroupID = r.ReadUint32()
	m.SkillGroupPriority = r.ReadUint16()
	m.AgentState = r.ReadUint16()
	m.EventReasonCode = r.ReadUint16()
	m.MRDID = r.ReadInt32()
	m.NumTasks = r.ReadUint32()
	m.AgentMode = r.ReadUint16()
	m.MaxTaskLimit = r.ReadUint32()
	m.ICMAgentID = r.ReadInt32()
	m.AgentAvailabilityStatus = r.ReadUint32()
	m.NumFltSkillGroups = r.ReadUint16()
	m.DepartmentID = r.ReadInt32()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.CTIClientSignature = ff.GetString(protocol.TagCTIClientSignature)
		m.AgentID = ff.GetString(protocol.TagAgentID)
		m.AgentExtension = ff.GetString(protocol.TagAgentExtension)
		m.ActiveTerminal = ff.GetString(protocol.TagActiveTerminal)
		m.AgentInstrument = ff.GetString(protocol.TagAgentInstrument)
		m.Duration = ff.GetUint32(protocol.TagDuration)
		m.NextAgentState = ff.GetUint16(protocol.TagNextAgentState)
		m.Direction = ff.GetUint32(protocol.TagDirection)
	}

	return nil
}

// StateName returns the human-readable name for the agent state.
func (m *AgentStateEvent) StateName() string {
	return protocol.AgentStateName(m.AgentState)
}
