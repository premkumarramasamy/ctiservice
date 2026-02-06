package messages

import (
	"ctiservice/internal/protocol"
)

// ConfigOperation represents the type of configuration change.
type ConfigOperation uint16

const (
	ConfigOperationAdd    ConfigOperation = 1
	ConfigOperationUpdate ConfigOperation = 2
	ConfigOperationDelete ConfigOperation = 3
)

// ConfigOperationName returns a human-readable name for a config operation.
func ConfigOperationName(op ConfigOperation) string {
	switch op {
	case ConfigOperationAdd:
		return "Add"
	case ConfigOperationUpdate:
		return "Update"
	case ConfigOperationDelete:
		return "Delete"
	default:
		return "Unknown"
	}
}

// AgentConfigRecord represents a single agent configuration record.
type AgentConfigRecord struct {
	RecordType    uint16 // Record type (Tag 183)
	AgentType     uint16 // Agent type (Tag 189)
	LoginID       string // Login ID (Tag 190)
	LastName      string // Last name (Tag 138)
	FirstName     string // First name (Tag 137)
	Extension     string // Extension (Tag 173/3)
	NumCSQ        uint16 // Number of CSQs (Tag 191)
	CSQIDs        []uint32 // CSQ IDs (Tag 62, repeated)
}

// ConfigAgentEvent is sent when agent configuration changes.
// Contains information about agent additions, updates, or deletions.
// Protocol Version 24 - CONFIG_AGENT_EVENT (MessageType = 237)
type ConfigAgentEvent struct {
	// Fixed Part
	PeripheralID    uint32 // Peripheral ID (UINT)
	ConfigOperation uint16 // Configuration operation (USHORT)
	NumRecords      uint16 // Number of records (USHORT)

	// Floating fields contain agent records
	AgentID       string // Tag 4 - Agent ID
	AgentExtension string // Tag 3 - Agent extension
	LoginID       string // Tag 190 - Login ID
	LastName      string // Tag 138 - Last name
	FirstName     string // Tag 137 - First name
	SkillGroupID  uint32 // Tag 10 - Skill group ID
	ICMAgentID    uint32 // ICM Agent ID from fixed fields or floating
}

func (m *ConfigAgentEvent) Type() uint32 {
	return protocol.MsgTypeConfigAgentEvent
}

func (m *ConfigAgentEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.ConfigOperation)
	w.WriteUint16(m.NumRecords)

	if err := w.Error(); err != nil {
		return nil, err
	}

	fw := protocol.NewFloatingFieldWriter()
	if m.AgentID != "" {
		fw.WriteString(protocol.TagAgentID, m.AgentID)
	}
	if m.AgentExtension != "" {
		fw.WriteString(protocol.TagAgentExtension, m.AgentExtension)
	}
	if m.LoginID != "" {
		fw.WriteString(protocol.TagLoginID, m.LoginID)
	}
	if m.LastName != "" {
		fw.WriteString(protocol.TagLastName, m.LastName)
	}
	if m.FirstName != "" {
		fw.WriteString(protocol.TagFirstName, m.FirstName)
	}
	if m.SkillGroupID != 0 {
		fw.WriteUint32(protocol.TagSkillGroupID, m.SkillGroupID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *ConfigAgentEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.PeripheralID = r.ReadUint32()
	m.ConfigOperation = r.ReadUint16()
	m.NumRecords = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.AgentID = ff.GetString(protocol.TagAgentID)
		m.AgentExtension = ff.GetString(protocol.TagAgentExtension)
		m.LoginID = ff.GetString(protocol.TagLoginID)
		m.LastName = ff.GetString(protocol.TagLastName)
		m.FirstName = ff.GetString(protocol.TagFirstName)
		m.SkillGroupID = ff.GetUint32(protocol.TagSkillGroupID)
	}

	return nil
}

// OperationName returns a human-readable name for the config operation.
func (m *ConfigAgentEvent) OperationName() string {
	return ConfigOperationName(ConfigOperation(m.ConfigOperation))
}

// ConfigDeviceEvent is sent when device configuration changes.
// Contains information about device additions, updates, or deletions.
// Protocol Version 24 - CONFIG_DEVICE_EVENT (MessageType = 238)
type ConfigDeviceEvent struct {
	// Fixed Part
	PeripheralID    uint32 // Peripheral ID (UINT)
	ConfigOperation uint16 // Configuration operation (USHORT)
	NumRecords      uint16 // Number of records (USHORT)

	// Floating fields contain device records
	DeviceID        string // Device identifier
	DeviceType      uint16 // Device type
	Extension       string // Extension (Tag 3)
	SkillGroupID    uint32 // Tag 10 - Associated skill group ID
	ServiceID       uint32 // Tag 8 - Associated service ID
}

func (m *ConfigDeviceEvent) Type() uint32 {
	return protocol.MsgTypeConfigDeviceEvent
}

func (m *ConfigDeviceEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.ConfigOperation)
	w.WriteUint16(m.NumRecords)

	if err := w.Error(); err != nil {
		return nil, err
	}

	fw := protocol.NewFloatingFieldWriter()
	if m.Extension != "" {
		fw.WriteString(protocol.TagAgentExtension, m.Extension)
	}
	if m.SkillGroupID != 0 {
		fw.WriteUint32(protocol.TagSkillGroupID, m.SkillGroupID)
	}
	if m.ServiceID != 0 {
		fw.WriteUint32(protocol.TagServiceID, m.ServiceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *ConfigDeviceEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.PeripheralID = r.ReadUint32()
	m.ConfigOperation = r.ReadUint16()
	m.NumRecords = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.Extension = ff.GetString(protocol.TagAgentExtension)
		m.SkillGroupID = ff.GetUint32(protocol.TagSkillGroupID)
		m.ServiceID = ff.GetUint32(protocol.TagServiceID)
	}

	return nil
}

// OperationName returns a human-readable name for the config operation.
func (m *ConfigDeviceEvent) OperationName() string {
	return ConfigOperationName(ConfigOperation(m.ConfigOperation))
}

// ConfigCSQEvent is sent when CSQ (Contact Service Queue) configuration changes.
// Protocol Version 24 - CONFIG_CSQ_EVENT (MessageType = 236)
type ConfigCSQEvent struct {
	// Fixed Part
	PeripheralID    uint32 // Peripheral ID (UINT)
	ConfigOperation uint16 // Configuration operation (USHORT)
	NumRecords      uint16 // Number of records (USHORT)

	// Floating fields
	CSQID           uint32 // Tag 62 - CSQ ID
	SkillGroupID    uint32 // Tag 10 - Skill group ID
	SkillGroupNumber uint32 // Tag 9 - Skill group number
	ServiceID       uint32 // Tag 8 - Service ID
	ServiceNumber   uint32 // Tag 7 - Service number
}

func (m *ConfigCSQEvent) Type() uint32 {
	return protocol.MsgTypeConfigCSQEvent
}

func (m *ConfigCSQEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.ConfigOperation)
	w.WriteUint16(m.NumRecords)

	if err := w.Error(); err != nil {
		return nil, err
	}

	fw := protocol.NewFloatingFieldWriter()
	if m.CSQID != 0 {
		fw.WriteUint32(protocol.TagCSQID, m.CSQID)
	}
	if m.SkillGroupID != 0 {
		fw.WriteUint32(protocol.TagSkillGroupID, m.SkillGroupID)
	}
	if m.SkillGroupNumber != 0 {
		fw.WriteUint32(protocol.TagSkillGroupNumber, m.SkillGroupNumber)
	}
	if m.ServiceID != 0 {
		fw.WriteUint32(protocol.TagServiceID, m.ServiceID)
	}
	if m.ServiceNumber != 0 {
		fw.WriteUint32(protocol.TagServiceNumber, m.ServiceNumber)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *ConfigCSQEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.PeripheralID = r.ReadUint32()
	m.ConfigOperation = r.ReadUint16()
	m.NumRecords = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.CSQID = ff.GetUint32(protocol.TagCSQID)
		m.SkillGroupID = ff.GetUint32(protocol.TagSkillGroupID)
		m.SkillGroupNumber = ff.GetUint32(protocol.TagSkillGroupNumber)
		m.ServiceID = ff.GetUint32(protocol.TagServiceID)
		m.ServiceNumber = ff.GetUint32(protocol.TagServiceNumber)
	}

	return nil
}

// OperationName returns a human-readable name for the config operation.
func (m *ConfigCSQEvent) OperationName() string {
	return ConfigOperationName(ConfigOperation(m.ConfigOperation))
}

// ConfigBeginEvent signals the start of configuration data transmission.
// Protocol Version 24 - CONFIG_BEGIN_EVENT (MessageType = 233)
type ConfigBeginEvent struct {
	// Fixed Part
	PeripheralID uint32 // Peripheral ID (UINT)
	ConfigType   uint16 // Type of configuration being sent (USHORT)
}

func (m *ConfigBeginEvent) Type() uint32 {
	return protocol.MsgTypeConfigBeginEvent
}

func (m *ConfigBeginEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.ConfigType)
	return w.Bytes(), w.Error()
}

func (m *ConfigBeginEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.PeripheralID = r.ReadUint32()
	m.ConfigType = r.ReadUint16()
	return r.Error()
}

// ConfigEndEvent signals the end of configuration data transmission.
// Protocol Version 24 - CONFIG_END_EVENT (MessageType = 234)
type ConfigEndEvent struct {
	// Fixed Part
	PeripheralID uint32 // Peripheral ID (UINT)
	ConfigType   uint16 // Type of configuration that was sent (USHORT)
	NumRecords   uint32 // Total number of records sent (UINT)
}

func (m *ConfigEndEvent) Type() uint32 {
	return protocol.MsgTypeConfigEndEvent
}

func (m *ConfigEndEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.ConfigType)
	w.WriteUint32(m.NumRecords)
	return w.Bytes(), w.Error()
}

func (m *ConfigEndEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.PeripheralID = r.ReadUint32()
	m.ConfigType = r.ReadUint16()
	m.NumRecords = r.ReadUint32()
	return r.Error()
}

// ConfigRequestEvent is used to request configuration data.
// Protocol Version 24 - CONFIG_REQUEST_EVENT (MessageType = 232)
type ConfigRequestEvent struct {
	// Fixed Part
	PeripheralID uint32 // Peripheral ID (UINT)
	ConfigType   uint16 // Type of configuration requested (USHORT)
}

func (m *ConfigRequestEvent) Type() uint32 {
	return protocol.MsgTypeConfigRequestEvent
}

func (m *ConfigRequestEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.ConfigType)
	return w.Bytes(), w.Error()
}

func (m *ConfigRequestEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.PeripheralID = r.ReadUint32()
	m.ConfigType = r.ReadUint16()
	return r.Error()
}
