package messages

import (
	"ctiservice/internal/protocol"
)

// BeginCallEvent is sent when a new call begins.
// Protocol Version 24 - BEGIN_CALL_EVENT (MessageType = 23)
type BeginCallEvent struct {
	// Fixed Part
	MonitorID              uint32 // Monitor ID
	PeripheralID           uint32 // Peripheral ID
	PeripheralType         uint16 // Type of peripheral (USHORT)
	NumCTIClients          uint16 // Number of CTI clients (USHORT)
	NumNamedVariables      uint16 // Number of named variables (USHORT)
	NumNamedArrays         uint16 // Number of named arrays (USHORT)
	CallType               uint16 // Type of call (USHORT)
	ConnectionDeviceIDType uint16 // Device ID type (USHORT)
	ConnectionCallID       uint32 // Call ID (UINT)
	CalledPartyDisposition uint16 // Called party disposition (USHORT)

	// Floating fields
	ConnectionDeviceID   string // Tag 31
	ANI                  string // Tag 15
	DNIS                 string // Tag 16
	DialedNumber         string // Tag 40
	CallerEnteredDigits  string // Tag 41
	UserToUserInfo       string // Tag 17
	CallWrapupData       string // Tag 30
	CallVariable1        string // Tag 18
	CallVariable2        string // Tag 19
	CallVariable3        string // Tag 20
	CallVariable4        string // Tag 21
	CallVariable5        string // Tag 22
	CallVariable6        string // Tag 23
	CallVariable7        string // Tag 24
	CallVariable8        string // Tag 25
	CallVariable9        string // Tag 26
	CallVariable10       string // Tag 27
	RouterCallKeyDay     uint32 // Tag 72
	RouterCallKeyCallID  uint32 // Tag 73
	RouterCallKeySeqNum  uint32 // Tag 214
}

func (m *BeginCallEvent) Type() uint32 {
	return protocol.MsgTypeBeginCallEvent
}

func (m *BeginCallEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.NumCTIClients)
	w.WriteUint16(m.NumNamedVariables)
	w.WriteUint16(m.NumNamedArrays)
	w.WriteUint16(m.CallType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.CalledPartyDisposition)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
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
	if m.CallWrapupData != "" {
		fw.WriteString(protocol.TagCallWrapupData, m.CallWrapupData)
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

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *BeginCallEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.NumCTIClients = r.ReadUint16()
	m.NumNamedVariables = r.ReadUint16()
	m.NumNamedArrays = r.ReadUint16()
	m.CallType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.CalledPartyDisposition = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields
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
		m.CallWrapupData = ff.GetString(protocol.TagCallWrapupData)
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
		m.RouterCallKeyDay = ff.GetUint32(protocol.TagRouterCallKeyDay)
		m.RouterCallKeyCallID = ff.GetUint32(protocol.TagRouterCallKeyCallID)
		m.RouterCallKeySeqNum = ff.GetUint32(protocol.TagRouterCallKeySeqNum)
	}

	return nil
}

// EndCallEvent is sent when a call ends.
// Protocol Version 24 - END_CALL_EVENT (MessageType = 24)
type EndCallEvent struct {
	// Fixed Part
	MonitorID              uint32 // Monitor ID
	PeripheralID           uint32 // Peripheral ID
	PeripheralType         uint16 // Type of peripheral
	ConnectionDeviceIDType uint16 // Device ID type
	ConnectionCallID       uint32 // Call ID

	// Floating fields
	ConnectionDeviceID string // Tag 31
}

func (m *EndCallEvent) Type() uint32 {
	return protocol.MsgTypeEndCallEvent
}

func (m *EndCallEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	return w.Bytes(), w.Error()
}

func (m *EndCallEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
	}

	return nil
}

// CallDeliveredEvent is sent when a call arrives at a device.
// Protocol Version 24 - CALL_DELIVERED_EVENT (MessageType = 9)
type CallDeliveredEvent struct {
	// Fixed Part
	MonitorID              uint32
	PeripheralID           uint32
	PeripheralType         uint16
	ConnectionDeviceIDType uint16
	ConnectionCallID       uint32
	LineHandle             uint16
	LineType               uint16
	ServiceNumber          uint32
	ServiceID              uint32
	SkillGroupNumber       uint32
	SkillGroupID           uint32
	SkillGroupPriority     uint16
	AlertingDeviceType     uint16
	CallingDeviceType      uint16
	CalledDeviceType       uint16
	LastRedirectDeviceType uint16
	LocalConnectionState   uint16
	EventCause             uint16
	NumNamedVariables      uint16
	NumNamedArrays         uint16

	// Floating fields
	ConnectionDeviceID   string
	AlertingDeviceID     string
	CallingDeviceID      string
	CalledDeviceID       string
	LastRedirectDeviceID string
	TrunkNumber          uint32
	TrunkGroupNumber     uint32
	SecondaryConnCallID  uint32
	ANI                  string
	DNIS                 string
	DialedNumber         string
	CallerEnteredDigits  string
	UserToUserInfo       string
	CallVariable1        string
	CallVariable2        string
	CallVariable3        string
	CallVariable4        string
	CallVariable5        string
	CallVariable6        string
	CallVariable7        string
	CallVariable8        string
	CallVariable9        string
	CallVariable10       string
	CallWrapupData       string
}

func (m *CallDeliveredEvent) Type() uint32 {
	return protocol.MsgTypeCallDeliveredEvent
}

func (m *CallDeliveredEvent) Encode() ([]byte, error) {
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
	w.WriteUint16(m.AlertingDeviceType)
	w.WriteUint16(m.CallingDeviceType)
	w.WriteUint16(m.CalledDeviceType)
	w.WriteUint16(m.LastRedirectDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	w.WriteUint16(m.NumNamedVariables)
	w.WriteUint16(m.NumNamedArrays)
	return w.Bytes(), w.Error()
}

func (m *CallDeliveredEvent) Decode(data []byte) error {
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
	m.AlertingDeviceType = r.ReadUint16()
	m.CallingDeviceType = r.ReadUint16()
	m.CalledDeviceType = r.ReadUint16()
	m.LastRedirectDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	m.NumNamedVariables = r.ReadUint16()
	m.NumNamedArrays = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.AlertingDeviceID = ff.GetString(protocol.TagAlertingDeviceID)
		m.CallingDeviceID = ff.GetString(protocol.TagCallingDeviceID)
		m.CalledDeviceID = ff.GetString(protocol.TagCalledDeviceID)
		m.LastRedirectDeviceID = ff.GetString(protocol.TagLastRedirectDeviceID)
		m.TrunkNumber = ff.GetUint32(protocol.TagTrunkNumber)
		m.TrunkGroupNumber = ff.GetUint32(protocol.TagTrunkGroupNumber)
		m.SecondaryConnCallID = ff.GetUint32(protocol.TagSecondaryConnCallID)
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
		m.CallWrapupData = ff.GetString(protocol.TagCallWrapupData)
	}

	return nil
}

// CallEstablishedEvent is sent when a call is answered.
// Protocol Version 24 - CALL_ESTABLISHED_EVENT (MessageType = 10)
type CallEstablishedEvent struct {
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
	AnsweringDeviceType    uint16 // Answering device type (USHORT)
	CallingDeviceType      uint16 // Calling device type (USHORT)
	CalledDeviceType       uint16 // Called device type (USHORT)
	LastRedirectDeviceType uint16 // Last redirect device type (USHORT)
	LocalConnectionState   uint16 // Local connection state (USHORT)
	EventCause             uint16 // Event cause (USHORT)

	// Floating fields (per GED-188 v24 spec)
	ConnectionDeviceID   string // Tag 31
	AnsweringDeviceID    string // Tag 33
	CallingDeviceID      string // Tag 12
	CalledDeviceID       string // Tag 13
	LastRedirectDeviceID string // Tag 14
	TrunkNumber          uint32 // Tag 121
	TrunkGroupNumber     uint32 // Tag 122
}

func (m *CallEstablishedEvent) Type() uint32 {
	return protocol.MsgTypeCallEstablishedEvent
}

func (m *CallEstablishedEvent) Encode() ([]byte, error) {
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
	w.WriteUint16(m.AnsweringDeviceType)
	w.WriteUint16(m.CallingDeviceType)
	w.WriteUint16(m.CalledDeviceType)
	w.WriteUint16(m.LastRedirectDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields (per GED-188 v24 spec)
	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}
	if m.AnsweringDeviceID != "" {
		fw.WriteString(protocol.TagAnsweringDeviceID, m.AnsweringDeviceID)
	}
	if m.CallingDeviceID != "" {
		fw.WriteString(protocol.TagCallingDeviceID, m.CallingDeviceID)
	}
	if m.CalledDeviceID != "" {
		fw.WriteString(protocol.TagCalledDeviceID, m.CalledDeviceID)
	}
	if m.LastRedirectDeviceID != "" {
		fw.WriteString(protocol.TagLastRedirectDeviceID, m.LastRedirectDeviceID)
	}
	if m.TrunkNumber != 0 {
		fw.WriteUint32(protocol.TagTrunkNumber, m.TrunkNumber)
	}
	if m.TrunkGroupNumber != 0 {
		fw.WriteUint32(protocol.TagTrunkGroupNumber, m.TrunkGroupNumber)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *CallEstablishedEvent) Decode(data []byte) error {
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
	m.AnsweringDeviceType = r.ReadUint16()
	m.CallingDeviceType = r.ReadUint16()
	m.CalledDeviceType = r.ReadUint16()
	m.LastRedirectDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields (per GED-188 v24 spec)
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.AnsweringDeviceID = ff.GetString(protocol.TagAnsweringDeviceID)
		m.CallingDeviceID = ff.GetString(protocol.TagCallingDeviceID)
		m.CalledDeviceID = ff.GetString(protocol.TagCalledDeviceID)
		m.LastRedirectDeviceID = ff.GetString(protocol.TagLastRedirectDeviceID)
		m.TrunkNumber = ff.GetUint32(protocol.TagTrunkNumber)
		m.TrunkGroupNumber = ff.GetUint32(protocol.TagTrunkGroupNumber)
	}

	return nil
}

// CallHeldEvent is sent when a call is placed on hold.
// Protocol Version 24 - CALL_HELD_EVENT (MessageType = 11)
type CallHeldEvent struct {
	// Fixed Part
	MonitorID              uint32 // Monitor ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	PeripheralType         uint16 // Peripheral type (USHORT)
	ConnectionDeviceIDType uint16 // Device ID type (USHORT)
	ConnectionCallID       uint32 // Call ID (UINT)
	HoldingDeviceType      uint16 // Holding device type (USHORT)
	LocalConnectionState   uint16 // Local connection state (USHORT)
	EventCause             uint16 // Event cause (USHORT)

	// Floating fields
	ConnectionDeviceID string // Tag 31
	HoldingDeviceID    string // Tag 34
}

func (m *CallHeldEvent) Type() uint32 {
	return protocol.MsgTypeCallHeldEvent
}

func (m *CallHeldEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.HoldingDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}
	if m.HoldingDeviceID != "" {
		fw.WriteString(protocol.TagHoldingDeviceID, m.HoldingDeviceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *CallHeldEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.HoldingDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.HoldingDeviceID = ff.GetString(protocol.TagHoldingDeviceID)
	}

	return nil
}

// CallRetrievedEvent is sent when a call is retrieved from hold.
// Protocol Version 24 - CALL_RETRIEVED_EVENT (MessageType = 12)
type CallRetrievedEvent struct {
	// Fixed Part
	MonitorID              uint32 // Monitor ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	PeripheralType         uint16 // Peripheral type (USHORT)
	ConnectionDeviceIDType uint16 // Device ID type (USHORT)
	ConnectionCallID       uint32 // Call ID (UINT)
	RetrievingDeviceType   uint16 // Retrieving device type (USHORT)
	LocalConnectionState   uint16 // Local connection state (USHORT)
	EventCause             uint16 // Event cause (USHORT)

	// Floating fields
	ConnectionDeviceID string // Tag 31
	RetrievingDeviceID string // Tag 35
}

func (m *CallRetrievedEvent) Type() uint32 {
	return protocol.MsgTypeCallRetrievedEvent
}

func (m *CallRetrievedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.RetrievingDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}
	if m.RetrievingDeviceID != "" {
		fw.WriteString(protocol.TagRetrievingDeviceID, m.RetrievingDeviceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *CallRetrievedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.RetrievingDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.RetrievingDeviceID = ff.GetString(protocol.TagRetrievingDeviceID)
	}

	return nil
}

// CallClearedEvent is sent when a call is terminated.
type CallClearedEvent struct {
	MonitorID              uint32
	PeripheralID           uint32
	PeripheralType         uint16
	ConnectionDeviceIDType uint16
	ConnectionCallID       uint32
	LocalConnectionState   uint16
	EventCause             uint16

	ConnectionDeviceID string
}

func (m *CallClearedEvent) Type() uint32 {
	return protocol.MsgTypeCallClearedEvent
}

func (m *CallClearedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	return w.Bytes(), w.Error()
}

func (m *CallClearedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
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
	}

	return nil
}

// CallConnectionClearedEvent is sent when a party leaves a call.
type CallConnectionClearedEvent struct {
	MonitorID              uint32
	PeripheralID           uint32
	PeripheralType         uint16
	ConnectionDeviceIDType uint16
	ConnectionCallID       uint32
	ReleasingDeviceType    uint16
	LocalConnectionState   uint16
	EventCause             uint16

	ConnectionDeviceID string
	ReleasingDeviceID  string
}

func (m *CallConnectionClearedEvent) Type() uint32 {
	return protocol.MsgTypeCallConnectionClearedEvent
}

func (m *CallConnectionClearedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.ReleasingDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	return w.Bytes(), w.Error()
}

func (m *CallConnectionClearedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.ReleasingDeviceType = r.ReadUint16()
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
		m.ReleasingDeviceID = ff.GetString(protocol.TagReleasingDeviceID)
	}

	return nil
}

// CallOriginatedEvent is sent when an outbound call is initiated.
type CallOriginatedEvent struct {
	MonitorID              uint32
	PeripheralID           uint32
	PeripheralType         uint16
	ConnectionDeviceIDType uint16
	ConnectionCallID       uint32
	LineHandle             uint16
	LineType               uint16
	ServiceNumber          uint32
	ServiceID              uint32
	SkillGroupNumber       uint32
	SkillGroupID           uint32
	SkillGroupPriority     uint16
	CallingDeviceType      uint16
	CalledDeviceType       uint16
	LocalConnectionState   uint16
	EventCause             uint16

	ConnectionDeviceID string
	CallingDeviceID    string
	CalledDeviceID     string
}

func (m *CallOriginatedEvent) Type() uint32 {
	return protocol.MsgTypeCallOriginatedEvent
}

func (m *CallOriginatedEvent) Encode() ([]byte, error) {
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
	w.WriteUint16(m.CalledDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	return w.Bytes(), w.Error()
}

func (m *CallOriginatedEvent) Decode(data []byte) error {
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
	m.CalledDeviceType = r.ReadUint16()
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
		m.CalledDeviceID = ff.GetString(protocol.TagCalledDeviceID)
	}

	return nil
}

// CallFailedEvent is sent when a call fails.
type CallFailedEvent struct {
	MonitorID              uint32
	PeripheralID           uint32
	PeripheralType         uint16
	ConnectionDeviceIDType uint16
	ConnectionCallID       uint32
	FailingDeviceType      uint16
	CalledDeviceType       uint16
	LocalConnectionState   uint16
	EventCause             uint16

	ConnectionDeviceID string
	FailingDeviceID    string
	CalledDeviceID     string
}

func (m *CallFailedEvent) Type() uint32 {
	return protocol.MsgTypeCallFailedEvent
}

func (m *CallFailedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.FailingDeviceType)
	w.WriteUint16(m.CalledDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	return w.Bytes(), w.Error()
}

func (m *CallFailedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.FailingDeviceType = r.ReadUint16()
	m.CalledDeviceType = r.ReadUint16()
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
		m.FailingDeviceID = ff.GetString(protocol.TagFailingDeviceID)
		m.CalledDeviceID = ff.GetString(protocol.TagCalledDeviceID)
	}

	return nil
}

// ConnectedParty represents a party in a conference or transferred call.
type ConnectedParty struct {
	CallID         uint32 // Call ID of this party
	DeviceIDType   uint16 // Device ID type
	DeviceID       string // Device identifier
}

// CallConferencedEvent is sent when a conference call is created.
// Protocol Version 24 - CALL_CONFERENCED_EVENT (MessageType = 17)
type CallConferencedEvent struct {
	// Fixed Part
	MonitorID            uint32 // Monitor ID (UINT)
	PeripheralID         uint32 // Peripheral ID (UINT)
	PeripheralType       uint16 // Peripheral type (USHORT)
	PrimaryDeviceIDType  uint16 // Primary device ID type (USHORT)
	PrimaryCallID        uint32 // Primary call ID (UINT)
	LineHandle           uint16 // Line handle (USHORT)
	LineType             uint16 // Line type (USHORT)
	SkillGroupNumber     uint32 // Skill group number (UINT)
	SkillGroupID         uint32 // Skill group ID (UINT)
	SkillGroupPriority   uint16 // Skill group priority (USHORT)
	NumParties           uint16 // Number of parties (USHORT)
	SecondaryDeviceIDType uint16 // Secondary device ID type (USHORT)
	SecondaryCallID      uint32 // Secondary call ID (UINT)
	ControllerDeviceType uint16 // Controller device type (USHORT)
	AddedPartyDeviceType uint16 // Added party device type (USHORT)
	LocalConnectionState uint16 // Local connection state (USHORT)
	EventCause           uint16 // Event cause (USHORT)

	// Floating fields (per GED-188 v24 spec)
	PrimaryDeviceID    string // Tag 46
	SecondaryDeviceID  string // Tag 47
	ControllerDeviceID string // Tag 42
	AddedPartyDeviceID string // Tag 43
	// ConnectedParties contains repeating party info (up to NumParties)
	ConnectedParties []ConnectedParty
}

func (m *CallConferencedEvent) Type() uint32 {
	return protocol.MsgTypeCallConferencedEvent
}

func (m *CallConferencedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.PrimaryDeviceIDType)
	w.WriteUint32(m.PrimaryCallID)
	w.WriteUint16(m.LineHandle)
	w.WriteUint16(m.LineType)
	w.WriteUint32(m.SkillGroupNumber)
	w.WriteUint32(m.SkillGroupID)
	w.WriteUint16(m.SkillGroupPriority)
	w.WriteUint16(m.NumParties)
	w.WriteUint16(m.SecondaryDeviceIDType)
	w.WriteUint32(m.SecondaryCallID)
	w.WriteUint16(m.ControllerDeviceType)
	w.WriteUint16(m.AddedPartyDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields (per GED-188 v24 spec)
	fw := protocol.NewFloatingFieldWriter()
	if m.PrimaryDeviceID != "" {
		fw.WriteString(protocol.TagPrimaryDeviceID, m.PrimaryDeviceID)
	}
	if m.SecondaryDeviceID != "" {
		fw.WriteString(protocol.TagSecondaryDeviceID, m.SecondaryDeviceID)
	}
	if m.ControllerDeviceID != "" {
		fw.WriteString(protocol.TagControllerDeviceID, m.ControllerDeviceID)
	}
	if m.AddedPartyDeviceID != "" {
		fw.WriteString(protocol.TagAddedPartyDeviceID, m.AddedPartyDeviceID)
	}
	// Note: ConnectedParty fields (repeating) encoding not implemented

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *CallConferencedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.PrimaryDeviceIDType = r.ReadUint16()
	m.PrimaryCallID = r.ReadUint32()
	m.LineHandle = r.ReadUint16()
	m.LineType = r.ReadUint16()
	m.SkillGroupNumber = r.ReadUint32()
	m.SkillGroupID = r.ReadUint32()
	m.SkillGroupPriority = r.ReadUint16()
	m.NumParties = r.ReadUint16()
	m.SecondaryDeviceIDType = r.ReadUint16()
	m.SecondaryCallID = r.ReadUint32()
	m.ControllerDeviceType = r.ReadUint16()
	m.AddedPartyDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields (per GED-188 v24 spec)
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.PrimaryDeviceID = ff.GetString(protocol.TagPrimaryDeviceID)
		m.SecondaryDeviceID = ff.GetString(protocol.TagSecondaryDeviceID)
		m.ControllerDeviceID = ff.GetString(protocol.TagControllerDeviceID)
		m.AddedPartyDeviceID = ff.GetString(protocol.TagAddedPartyDeviceID)
		// Note: ConnectedParty fields (repeating) parsing not implemented
	}

	return nil
}

// CallTransferredEvent is sent when a call is transferred.
// Protocol Version 24 - CALL_TRANSFERRED_EVENT (MessageType = 18)
type CallTransferredEvent struct {
	// Fixed Part
	MonitorID              uint32 // Monitor ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	PeripheralType         uint16 // Peripheral type (USHORT)
	PrimaryDeviceIDType    uint16 // Primary device ID type (USHORT)
	PrimaryCallID          uint32 // Primary call ID (UINT)
	LineHandle             uint16 // Line handle (USHORT)
	LineType               uint16 // Line type (USHORT)
	SkillGroupNumber       uint32 // Skill group number (UINT)
	SkillGroupID           uint32 // Skill group ID (UINT)
	SkillGroupPriority     uint16 // Skill group priority (USHORT)
	NumParties             uint16 // Number of parties (USHORT)
	SecondaryDeviceIDType  uint16 // Secondary device ID type (USHORT)
	SecondaryCallID        uint32 // Secondary call ID (UINT)
	TransferringDeviceType uint16 // Transferring device type (USHORT)
	TransferredDeviceType  uint16 // Transferred device type (USHORT)
	LocalConnectionState   uint16 // Local connection state (USHORT)
	EventCause             uint16 // Event cause (USHORT)

	// Floating fields (per GED-188 v24 spec)
	PrimaryDeviceID      string // Tag 46
	SecondaryDeviceID    string // Tag 47
	TransferringDeviceID string // Tag 38
	TransferredDeviceID  string // Tag 39
	// ConnectedParties contains repeating party info (up to NumParties)
	ConnectedParties []ConnectedParty
}

func (m *CallTransferredEvent) Type() uint32 {
	return protocol.MsgTypeCallTransferredEvent
}

func (m *CallTransferredEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.PrimaryDeviceIDType)
	w.WriteUint32(m.PrimaryCallID)
	w.WriteUint16(m.LineHandle)
	w.WriteUint16(m.LineType)
	w.WriteUint32(m.SkillGroupNumber)
	w.WriteUint32(m.SkillGroupID)
	w.WriteUint16(m.SkillGroupPriority)
	w.WriteUint16(m.NumParties)
	w.WriteUint16(m.SecondaryDeviceIDType)
	w.WriteUint32(m.SecondaryCallID)
	w.WriteUint16(m.TransferringDeviceType)
	w.WriteUint16(m.TransferredDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields (per GED-188 v24 spec)
	fw := protocol.NewFloatingFieldWriter()
	if m.PrimaryDeviceID != "" {
		fw.WriteString(protocol.TagPrimaryDeviceID, m.PrimaryDeviceID)
	}
	if m.SecondaryDeviceID != "" {
		fw.WriteString(protocol.TagSecondaryDeviceID, m.SecondaryDeviceID)
	}
	if m.TransferringDeviceID != "" {
		fw.WriteString(protocol.TagTransferringDeviceID, m.TransferringDeviceID)
	}
	if m.TransferredDeviceID != "" {
		fw.WriteString(protocol.TagTransferredDeviceID, m.TransferredDeviceID)
	}
	// Note: ConnectedParty fields (repeating) encoding not implemented

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *CallTransferredEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.PrimaryDeviceIDType = r.ReadUint16()
	m.PrimaryCallID = r.ReadUint32()
	m.LineHandle = r.ReadUint16()
	m.LineType = r.ReadUint16()
	m.SkillGroupNumber = r.ReadUint32()
	m.SkillGroupID = r.ReadUint32()
	m.SkillGroupPriority = r.ReadUint16()
	m.NumParties = r.ReadUint16()
	m.SecondaryDeviceIDType = r.ReadUint16()
	m.SecondaryCallID = r.ReadUint32()
	m.TransferringDeviceType = r.ReadUint16()
	m.TransferredDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields (per GED-188 v24 spec)
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.PrimaryDeviceID = ff.GetString(protocol.TagPrimaryDeviceID)
		m.SecondaryDeviceID = ff.GetString(protocol.TagSecondaryDeviceID)
		m.TransferringDeviceID = ff.GetString(protocol.TagTransferringDeviceID)
		m.TransferredDeviceID = ff.GetString(protocol.TagTransferredDeviceID)
		// Note: ConnectedParty fields (repeating) parsing not implemented
	}

	return nil
}

// CallQueuedEvent is sent when a call is queued.
type CallQueuedEvent struct {
	MonitorID              uint32
	PeripheralID           uint32
	PeripheralType         uint16
	ConnectionDeviceIDType uint16
	ConnectionCallID       uint32
	LocalConnectionState   uint16
	EventCause             uint16

	ConnectionDeviceID string
}

func (m *CallQueuedEvent) Type() uint32 {
	return protocol.MsgTypeCallQueuedEvent
}

func (m *CallQueuedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	return w.Bytes(), w.Error()
}

func (m *CallQueuedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
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
	}

	return nil
}

// CallDequeuedEvent is sent when a call is removed from a queue.
type CallDequeuedEvent struct {
	MonitorID              uint32
	PeripheralID           uint32
	PeripheralType         uint16
	ConnectionDeviceIDType uint16
	ConnectionCallID       uint32
	LocalConnectionState   uint16
	EventCause             uint16

	ConnectionDeviceID string
}

func (m *CallDequeuedEvent) Type() uint32 {
	return protocol.MsgTypeCallDequeuedEvent
}

func (m *CallDequeuedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	return w.Bytes(), w.Error()
}

func (m *CallDequeuedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
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
	}

	return nil
}

// CallDataUpdateEvent is sent when call data changes.
// Protocol Version 24 - CALL_DATA_UPDATE_EVENT (MessageType = 25)
type CallDataUpdateEvent struct {
	// Fixed Part
	MonitorID                  uint32 // Monitor ID (UINT)
	PeripheralID               uint32 // Peripheral ID (UINT)
	PeripheralType             uint16 // Peripheral type (USHORT)
	NumCTIClients              uint16 // Number of CTI clients (USHORT)
	NumNamedVariables          uint16 // Number of named variables (USHORT)
	NumNamedArrays             uint16 // Number of named arrays (USHORT)
	CallType                   uint16 // Type of call (USHORT)
	ConnectionDeviceIDType     uint16 // Device ID type (USHORT)
	ConnectionCallID           uint32 // Call ID (UINT)
	NewConnectionDeviceIDType  uint16 // New connection device ID type (USHORT)
	NewConnectionCallID        uint32 // New connection call ID (UINT)
	CalledPartyDisposition     uint16 // Called party disposition (USHORT)
	CampaignID                 uint32 // Campaign ID (UINT)
	QueryRuleID                uint32 // Query rule ID (UINT)

	// Floating fields
	ConnectionDeviceID    string // Tag 31
	NewConnectionDeviceID string // Tag 186
	ANI                   string // Tag 15
	DNIS                  string // Tag 16
	DialedNumber          string // Tag 40
	CallerEnteredDigits   string // Tag 41
	UserToUserInfo        string // Tag 17
	CallWrapupData        string // Tag 30
	CallVariable1         string // Tag 18
	CallVariable2         string // Tag 19
	CallVariable3         string // Tag 20
	CallVariable4         string // Tag 21
	CallVariable5         string // Tag 22
	CallVariable6         string // Tag 23
	CallVariable7         string // Tag 24
	CallVariable8         string // Tag 25
	CallVariable9         string // Tag 26
	CallVariable10        string // Tag 27
	RouterCallKeyDay      uint32 // Tag 72
	RouterCallKeyCallID   uint32 // Tag 73
	RouterCallKeySeqNum   uint32 // Tag 214
}

func (m *CallDataUpdateEvent) Type() uint32 {
	return protocol.MsgTypeCallDataUpdateEvent
}

func (m *CallDataUpdateEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.MonitorID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint16(m.PeripheralType)
	w.WriteUint16(m.NumCTIClients)
	w.WriteUint16(m.NumNamedVariables)
	w.WriteUint16(m.NumNamedArrays)
	w.WriteUint16(m.CallType)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.NewConnectionDeviceIDType)
	w.WriteUint32(m.NewConnectionCallID)
	w.WriteUint16(m.CalledPartyDisposition)
	w.WriteUint32(m.CampaignID)
	w.WriteUint32(m.QueryRuleID)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}
	if m.NewConnectionDeviceID != "" {
		fw.WriteString(protocol.TagNewConnectionDeviceID, m.NewConnectionDeviceID)
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
	if m.CallWrapupData != "" {
		fw.WriteString(protocol.TagCallWrapupData, m.CallWrapupData)
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

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *CallDataUpdateEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.MonitorID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.PeripheralType = r.ReadUint16()
	m.NumCTIClients = r.ReadUint16()
	m.NumNamedVariables = r.ReadUint16()
	m.NumNamedArrays = r.ReadUint16()
	m.CallType = r.ReadUint16()
	m.ConnectionDeviceIDType = r.ReadUint16()
	m.ConnectionCallID = r.ReadUint32()
	m.NewConnectionDeviceIDType = r.ReadUint16()
	m.NewConnectionCallID = r.ReadUint32()
	m.CalledPartyDisposition = r.ReadUint16()
	m.CampaignID = r.ReadUint32()
	m.QueryRuleID = r.ReadUint32()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.NewConnectionDeviceID = ff.GetString(protocol.TagNewConnectionDeviceID)
		m.ANI = ff.GetString(protocol.TagANI)
		m.DNIS = ff.GetString(protocol.TagDNIS)
		m.DialedNumber = ff.GetString(protocol.TagDialedNumber)
		m.CallerEnteredDigits = ff.GetString(protocol.TagCallerEnteredDigits)
		m.UserToUserInfo = ff.GetString(protocol.TagUserToUserInfo)
		m.CallWrapupData = ff.GetString(protocol.TagCallWrapupData)
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
		m.RouterCallKeyDay = ff.GetUint32(protocol.TagRouterCallKeyDay)
		m.RouterCallKeyCallID = ff.GetUint32(protocol.TagRouterCallKeyCallID)
		m.RouterCallKeySeqNum = ff.GetUint32(protocol.TagRouterCallKeySeqNum)
	}

	return nil
}
