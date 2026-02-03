package messages

import (
	"ctiservice/internal/protocol"
)

// AgentStateEvent reports an agent's state change.
type AgentStateEvent struct {
	MonitorID              uint32 // Monitor ID
	PeripheralID           uint32 // Peripheral ID
	SessionID              uint32 // Session ID
	PeripheralType         uint16 // Type of peripheral
	SkillGroupState        uint16 // Skill group state
	StateDuration          uint32 // Duration in current state (seconds)
	SkillGroupNumber       uint32 // Skill group number
	SkillGroupID           uint32 // Skill group ID
	SkillGroupPriority     uint16 // Skill group priority
	AgentState             uint16 // Current agent state
	EventReasonCode        uint16 // Reason code for state change
	MRDID                  uint16 // Media routing domain ID
	NumTasks               uint32 // Number of active tasks
	AgentMode              uint16 // Agent mode
	MaxTaskLimit           uint16 // Maximum task limit
	ICMAgentID             uint32 // ICM agent ID
	AgentAvailabilityStatus uint32 // Availability status
	NumFltSkillGroups      uint16 // Number of skill groups (floating)
	Reserved               uint16 // Reserved

	// Floating fields
	AgentExtension  string // Agent's extension
	AgentID         string // Agent's ID
	AgentInstrument string // Agent's instrument
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
	w.WriteUint16(m.MRDID)
	w.WriteUint32(m.NumTasks)
	w.WriteUint16(m.AgentMode)
	w.WriteUint16(m.MaxTaskLimit)
	w.WriteUint32(m.ICMAgentID)
	w.WriteUint32(m.AgentAvailabilityStatus)
	w.WriteUint16(m.NumFltSkillGroups)
	w.WriteUint16(m.Reserved)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.AgentExtension != "" {
		fw.WriteString(protocol.TagAgentExtension, m.AgentExtension)
	}
	if m.AgentID != "" {
		fw.WriteString(protocol.TagAgentID, m.AgentID)
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
	m.MRDID = r.ReadUint16()
	m.NumTasks = r.ReadUint32()
	m.AgentMode = r.ReadUint16()
	m.MaxTaskLimit = r.ReadUint16()
	m.ICMAgentID = r.ReadUint32()
	m.AgentAvailabilityStatus = r.ReadUint32()
	m.NumFltSkillGroups = r.ReadUint16()
	m.Reserved = r.ReadUint16()

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

// StateName returns the human-readable name for the agent state.
func (m *AgentStateEvent) StateName() string {
	return protocol.AgentStateName(m.AgentState)
}
