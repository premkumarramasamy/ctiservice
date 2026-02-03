package messages

import (
	"ctiservice/internal/protocol"
)

// BaseCallEvent contains common fields for call events.
type BaseCallEvent struct {
	MonitorID            uint32 // Monitor ID
	PeripheralID         uint32 // Peripheral ID
	PeripheralType       uint16 // Type of peripheral
	ConnectionDeviceType uint16 // Device type
	ConnectionCallID     uint32 // Call ID
	ConnectionState      uint16 // Current connection state
	Reserved             uint16 // Reserved
}

func (b *BaseCallEvent) readFrom(r *protocol.FixedFieldReader) {
	b.MonitorID = r.ReadUint32()
	b.PeripheralID = r.ReadUint32()
	b.PeripheralType = r.ReadUint16()
	b.ConnectionDeviceType = r.ReadUint16()
	b.ConnectionCallID = r.ReadUint32()
	b.ConnectionState = r.ReadUint16()
	b.Reserved = r.ReadUint16()
}

func (b *BaseCallEvent) writeTo(w *protocol.FixedFieldWriter) {
	w.WriteUint32(b.MonitorID)
	w.WriteUint32(b.PeripheralID)
	w.WriteUint16(b.PeripheralType)
	w.WriteUint16(b.ConnectionDeviceType)
	w.WriteUint32(b.ConnectionCallID)
	w.WriteUint16(b.ConnectionState)
	w.WriteUint16(b.Reserved)
}

// CallVariables holds the 10 call variables from floating fields.
type CallVariables struct {
	Var1  string
	Var2  string
	Var3  string
	Var4  string
	Var5  string
	Var6  string
	Var7  string
	Var8  string
	Var9  string
	Var10 string
}

func (cv *CallVariables) parseFrom(ff *protocol.FloatingFields) {
	cv.Var1 = ff.GetString(protocol.TagCallVariable1)
	cv.Var2 = ff.GetString(protocol.TagCallVariable2)
	cv.Var3 = ff.GetString(protocol.TagCallVariable3)
	cv.Var4 = ff.GetString(protocol.TagCallVariable4)
	cv.Var5 = ff.GetString(protocol.TagCallVariable5)
	cv.Var6 = ff.GetString(protocol.TagCallVariable6)
	cv.Var7 = ff.GetString(protocol.TagCallVariable7)
	cv.Var8 = ff.GetString(protocol.TagCallVariable8)
	cv.Var9 = ff.GetString(protocol.TagCallVariable9)
	cv.Var10 = ff.GetString(protocol.TagCallVariable10)
}

// BeginCallEvent is sent when a new call begins.
type BeginCallEvent struct {
	BaseCallEvent
	ServiceNumber   uint32 // Service number
	ServiceID       uint32 // Service ID
	SkillGroupNumber uint32 // Skill group number
	SkillGroupID    uint32 // Skill group ID
	SkillGroupPriority uint16 // Skill group priority
	CallType        uint16 // Type of call
	CallingDeviceType uint16 // Calling device type
	CalledDeviceType uint16 // Called device type
	LastRedirectDeviceType uint16 // Last redirect device type
	Reserved2       uint16 // Reserved

	// Floating fields
	ANI               string // Automatic Number Identification
	DNIS              string // Dialed Number Identification Service
	CallingDeviceID   string // Calling device ID
	CalledDeviceID    string // Called device ID
	LastRedirectDeviceID string // Last redirect device ID
	UserToUserInfo    string // User-to-user information
	CallVariables     CallVariables
}

func (m *BeginCallEvent) Type() uint32 {
	return protocol.MsgTypeBeginCallEvent
}

func (m *BeginCallEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint32(m.ServiceNumber)
	w.WriteUint32(m.ServiceID)
	w.WriteUint32(m.SkillGroupNumber)
	w.WriteUint32(m.SkillGroupID)
	w.WriteUint16(m.SkillGroupPriority)
	w.WriteUint16(m.CallType)
	w.WriteUint16(m.CallingDeviceType)
	w.WriteUint16(m.CalledDeviceType)
	w.WriteUint16(m.LastRedirectDeviceType)
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *BeginCallEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.ServiceNumber = r.ReadUint32()
	m.ServiceID = r.ReadUint32()
	m.SkillGroupNumber = r.ReadUint32()
	m.SkillGroupID = r.ReadUint32()
	m.SkillGroupPriority = r.ReadUint16()
	m.CallType = r.ReadUint16()
	m.CallingDeviceType = r.ReadUint16()
	m.CalledDeviceType = r.ReadUint16()
	m.LastRedirectDeviceType = r.ReadUint16()
	m.Reserved2 = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ANI = ff.GetString(protocol.TagANI)
		m.DNIS = ff.GetString(protocol.TagDNIS)
		m.CallingDeviceID = ff.GetString(protocol.TagCallingDeviceID)
		m.CalledDeviceID = ff.GetString(protocol.TagCalledDeviceID)
		m.LastRedirectDeviceID = ff.GetString(protocol.TagLastRedirectDeviceID)
		m.UserToUserInfo = ff.GetString(protocol.TagUserToUserInfo)
		m.CallVariables.parseFrom(ff)
	}

	return nil
}

// EndCallEvent is sent when a call ends.
type EndCallEvent struct {
	BaseCallEvent
}

func (m *EndCallEvent) Type() uint32 {
	return protocol.MsgTypeEndCallEvent
}

func (m *EndCallEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	return w.Bytes(), w.Error()
}

func (m *EndCallEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	return r.Error()
}

// CallDataUpdateEvent is sent when call data changes.
type CallDataUpdateEvent struct {
	BaseCallEvent
	NumCTIClients   uint16 // Number of CTI clients
	NumNamedVars    uint16 // Number of named variables
	NumNamedArrays  uint16 // Number of named arrays
	CallType        uint16 // Call type
	CallDisposition uint32 // Call disposition
	Reserved2       uint32 // Reserved

	// Floating fields
	ANI            string
	DNIS           string
	UserToUserInfo string
	CallVariables  CallVariables
}

func (m *CallDataUpdateEvent) Type() uint32 {
	return protocol.MsgTypeCallDataUpdateEvent
}

func (m *CallDataUpdateEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint16(m.NumCTIClients)
	w.WriteUint16(m.NumNamedVars)
	w.WriteUint16(m.NumNamedArrays)
	w.WriteUint16(m.CallType)
	w.WriteUint32(m.CallDisposition)
	w.WriteUint32(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallDataUpdateEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.NumCTIClients = r.ReadUint16()
	m.NumNamedVars = r.ReadUint16()
	m.NumNamedArrays = r.ReadUint16()
	m.CallType = r.ReadUint16()
	m.CallDisposition = r.ReadUint32()
	m.Reserved2 = r.ReadUint32()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ANI = ff.GetString(protocol.TagANI)
		m.DNIS = ff.GetString(protocol.TagDNIS)
		m.UserToUserInfo = ff.GetString(protocol.TagUserToUserInfo)
		m.CallVariables.parseFrom(ff)
	}

	return nil
}

// CallDeliveredEvent is sent when a call arrives at a device.
type CallDeliveredEvent struct {
	BaseCallEvent
	LineHandle       uint16 // Line handle
	LineType         uint16 // Line type
	ServiceNumber    uint32 // Service number
	ServiceID        uint32 // Service ID
	SkillGroupNumber uint32 // Skill group number
	SkillGroupID     uint32 // Skill group ID
	SkillGroupPriority uint16 // Skill group priority
	AlertingDeviceType uint16 // Alerting device type
	CallingDeviceType uint16 // Calling device type
	CalledDeviceType uint16 // Called device type
	LastRedirectDeviceType uint16 // Last redirect device type
	LocalConnectionState uint16 // Local connection state
	EventCause       uint16 // Event cause
	Reserved2        uint16 // Reserved

	// Floating fields
	ANI               string
	DNIS              string
	CallingDeviceID   string
	CalledDeviceID    string
	LastRedirectDeviceID string
	AlertingDeviceID  string
}

func (m *CallDeliveredEvent) Type() uint32 {
	return protocol.MsgTypeCallDeliveredEvent
}

func (m *CallDeliveredEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
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
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallDeliveredEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
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
	m.Reserved2 = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ANI = ff.GetString(protocol.TagANI)
		m.DNIS = ff.GetString(protocol.TagDNIS)
		m.CallingDeviceID = ff.GetString(protocol.TagCallingDeviceID)
		m.CalledDeviceID = ff.GetString(protocol.TagCalledDeviceID)
		m.LastRedirectDeviceID = ff.GetString(protocol.TagLastRedirectDeviceID)
	}

	return nil
}

// CallEstablishedEvent is sent when a call is answered.
type CallEstablishedEvent struct {
	BaseCallEvent
	LineHandle           uint16
	LineType             uint16
	ServiceNumber        uint32
	ServiceID            uint32
	SkillGroupNumber     uint32
	SkillGroupID         uint32
	SkillGroupPriority   uint16
	AnsweringDeviceType  uint16
	CallingDeviceType    uint16
	CalledDeviceType     uint16
	LastRedirectDeviceType uint16
	LocalConnectionState uint16
	EventCause           uint16
	Reserved2            uint16

	// Floating fields
	ANI               string
	DNIS              string
	CallingDeviceID   string
	CalledDeviceID    string
	AnsweringDeviceID string
}

func (m *CallEstablishedEvent) Type() uint32 {
	return protocol.MsgTypeCallEstablishedEvent
}

func (m *CallEstablishedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
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
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallEstablishedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
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
	m.Reserved2 = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ANI = ff.GetString(protocol.TagANI)
		m.DNIS = ff.GetString(protocol.TagDNIS)
		m.CallingDeviceID = ff.GetString(protocol.TagCallingDeviceID)
		m.CalledDeviceID = ff.GetString(protocol.TagCalledDeviceID)
	}

	return nil
}

// CallHeldEvent is sent when a call is placed on hold.
type CallHeldEvent struct {
	BaseCallEvent
	HoldingDeviceType uint16
	LocalConnectionState uint16
	EventCause        uint16
	Reserved2         uint16

	HoldingDeviceID string
}

func (m *CallHeldEvent) Type() uint32 {
	return protocol.MsgTypeCallHeldEvent
}

func (m *CallHeldEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint16(m.HoldingDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallHeldEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.HoldingDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	m.Reserved2 = r.ReadUint16()
	return r.Error()
}

// CallRetrievedEvent is sent when a call is retrieved from hold.
type CallRetrievedEvent struct {
	BaseCallEvent
	RetrievingDeviceType uint16
	LocalConnectionState uint16
	EventCause          uint16
	Reserved2           uint16

	RetrievingDeviceID string
}

func (m *CallRetrievedEvent) Type() uint32 {
	return protocol.MsgTypeCallRetrievedEvent
}

func (m *CallRetrievedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint16(m.RetrievingDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallRetrievedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.RetrievingDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	m.Reserved2 = r.ReadUint16()
	return r.Error()
}

// CallClearedEvent is sent when a call is terminated.
type CallClearedEvent struct {
	BaseCallEvent
	LocalConnectionState uint16
	EventCause          uint16
}

func (m *CallClearedEvent) Type() uint32 {
	return protocol.MsgTypeCallClearedEvent
}

func (m *CallClearedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	return w.Bytes(), w.Error()
}

func (m *CallClearedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	return r.Error()
}

// CallConnectionClearedEvent is sent when a party leaves a call.
type CallConnectionClearedEvent struct {
	BaseCallEvent
	ReleasingDeviceType uint16
	LocalConnectionState uint16
	EventCause          uint16
	Reserved2           uint16

	ReleasingDeviceID string
}

func (m *CallConnectionClearedEvent) Type() uint32 {
	return protocol.MsgTypeCallConnectionClearedEvent
}

func (m *CallConnectionClearedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint16(m.ReleasingDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallConnectionClearedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.ReleasingDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	m.Reserved2 = r.ReadUint16()
	return r.Error()
}

// CallOriginatedEvent is sent when an outbound call is initiated.
type CallOriginatedEvent struct {
	BaseCallEvent
	LineHandle         uint16
	LineType           uint16
	ServiceNumber      uint32
	ServiceID          uint32
	SkillGroupNumber   uint32
	SkillGroupID       uint32
	SkillGroupPriority uint16
	CallingDeviceType  uint16
	CalledDeviceType   uint16
	LocalConnectionState uint16
	EventCause         uint16
	Reserved2          uint16

	CallingDeviceID string
	CalledDeviceID  string
}

func (m *CallOriginatedEvent) Type() uint32 {
	return protocol.MsgTypeCallOriginatedEvent
}

func (m *CallOriginatedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
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
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallOriginatedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
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
	m.Reserved2 = r.ReadUint16()

	if err := r.Error(); err != nil {
		return err
	}

	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.CallingDeviceID = ff.GetString(protocol.TagCallingDeviceID)
		m.CalledDeviceID = ff.GetString(protocol.TagCalledDeviceID)
	}

	return nil
}

// CallFailedEvent is sent when a call fails.
type CallFailedEvent struct {
	BaseCallEvent
	FailingDeviceType uint16
	CalledDeviceType  uint16
	LocalConnectionState uint16
	EventCause        uint16

	FailingDeviceID string
	CalledDeviceID  string
}

func (m *CallFailedEvent) Type() uint32 {
	return protocol.MsgTypeCallFailedEvent
}

func (m *CallFailedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint16(m.FailingDeviceType)
	w.WriteUint16(m.CalledDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	return w.Bytes(), w.Error()
}

func (m *CallFailedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.FailingDeviceType = r.ReadUint16()
	m.CalledDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	return r.Error()
}

// CallConferencedEvent is sent when a conference call is created.
type CallConferencedEvent struct {
	BaseCallEvent
	PrimaryCallID        uint32
	PrimaryDeviceType    uint16
	PrimaryDeviceIDType  uint16
	SecondaryCallID      uint32
	SecondaryDeviceType  uint16
	SecondaryDeviceIDType uint16
	ControllerDeviceType uint16
	LocalConnectionState uint16
	EventCause           uint16
	Reserved2            uint16

	PrimaryDeviceID    string
	SecondaryDeviceID  string
	ControllerDeviceID string
}

func (m *CallConferencedEvent) Type() uint32 {
	return protocol.MsgTypeCallConferencedEvent
}

func (m *CallConferencedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint32(m.PrimaryCallID)
	w.WriteUint16(m.PrimaryDeviceType)
	w.WriteUint16(m.PrimaryDeviceIDType)
	w.WriteUint32(m.SecondaryCallID)
	w.WriteUint16(m.SecondaryDeviceType)
	w.WriteUint16(m.SecondaryDeviceIDType)
	w.WriteUint16(m.ControllerDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallConferencedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.PrimaryCallID = r.ReadUint32()
	m.PrimaryDeviceType = r.ReadUint16()
	m.PrimaryDeviceIDType = r.ReadUint16()
	m.SecondaryCallID = r.ReadUint32()
	m.SecondaryDeviceType = r.ReadUint16()
	m.SecondaryDeviceIDType = r.ReadUint16()
	m.ControllerDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	m.Reserved2 = r.ReadUint16()
	return r.Error()
}

// CallTransferredEvent is sent when a call is transferred.
type CallTransferredEvent struct {
	BaseCallEvent
	PrimaryCallID         uint32
	PrimaryDeviceType     uint16
	PrimaryDeviceIDType   uint16
	SecondaryCallID       uint32
	SecondaryDeviceType   uint16
	SecondaryDeviceIDType uint16
	TransferringDeviceType uint16
	TransferredDeviceType uint16
	LocalConnectionState  uint16
	EventCause            uint16

	PrimaryDeviceID      string
	SecondaryDeviceID    string
	TransferringDeviceID string
	TransferredDeviceID  string
}

func (m *CallTransferredEvent) Type() uint32 {
	return protocol.MsgTypeCallTransferredEvent
}

func (m *CallTransferredEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint32(m.PrimaryCallID)
	w.WriteUint16(m.PrimaryDeviceType)
	w.WriteUint16(m.PrimaryDeviceIDType)
	w.WriteUint32(m.SecondaryCallID)
	w.WriteUint16(m.SecondaryDeviceType)
	w.WriteUint16(m.SecondaryDeviceIDType)
	w.WriteUint16(m.TransferringDeviceType)
	w.WriteUint16(m.TransferredDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	return w.Bytes(), w.Error()
}

func (m *CallTransferredEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.PrimaryCallID = r.ReadUint32()
	m.PrimaryDeviceType = r.ReadUint16()
	m.PrimaryDeviceIDType = r.ReadUint16()
	m.SecondaryCallID = r.ReadUint32()
	m.SecondaryDeviceType = r.ReadUint16()
	m.SecondaryDeviceIDType = r.ReadUint16()
	m.TransferringDeviceType = r.ReadUint16()
	m.TransferredDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	return r.Error()
}

// CallQueuedEvent is sent when a call is queued.
type CallQueuedEvent struct {
	BaseCallEvent
	ServiceNumber    uint32
	ServiceID        uint32
	SkillGroupNumber uint32
	SkillGroupID     uint32
	SkillGroupPriority uint16
	QueueDeviceType  uint16
	CallingDeviceType uint16
	CalledDeviceType uint16
	LastRedirectDeviceType uint16
	LocalConnectionState uint16
	EventCause       uint16
	Reserved2        uint16

	QueueDeviceID        string
	CallingDeviceID      string
	CalledDeviceID       string
	LastRedirectDeviceID string
}

func (m *CallQueuedEvent) Type() uint32 {
	return protocol.MsgTypeCallQueuedEvent
}

func (m *CallQueuedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint32(m.ServiceNumber)
	w.WriteUint32(m.ServiceID)
	w.WriteUint32(m.SkillGroupNumber)
	w.WriteUint32(m.SkillGroupID)
	w.WriteUint16(m.SkillGroupPriority)
	w.WriteUint16(m.QueueDeviceType)
	w.WriteUint16(m.CallingDeviceType)
	w.WriteUint16(m.CalledDeviceType)
	w.WriteUint16(m.LastRedirectDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallQueuedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.ServiceNumber = r.ReadUint32()
	m.ServiceID = r.ReadUint32()
	m.SkillGroupNumber = r.ReadUint32()
	m.SkillGroupID = r.ReadUint32()
	m.SkillGroupPriority = r.ReadUint16()
	m.QueueDeviceType = r.ReadUint16()
	m.CallingDeviceType = r.ReadUint16()
	m.CalledDeviceType = r.ReadUint16()
	m.LastRedirectDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	m.Reserved2 = r.ReadUint16()
	return r.Error()
}

// CallDequeuedEvent is sent when a call is removed from a queue.
type CallDequeuedEvent struct {
	BaseCallEvent
	ServiceNumber   uint32
	ServiceID       uint32
	QueueDeviceType uint16
	LocalConnectionState uint16
	EventCause      uint16
	Reserved2       uint16

	QueueDeviceID string
}

func (m *CallDequeuedEvent) Type() uint32 {
	return protocol.MsgTypeCallDequeuedEvent
}

func (m *CallDequeuedEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	m.BaseCallEvent.writeTo(w)
	w.WriteUint32(m.ServiceNumber)
	w.WriteUint32(m.ServiceID)
	w.WriteUint16(m.QueueDeviceType)
	w.WriteUint16(m.LocalConnectionState)
	w.WriteUint16(m.EventCause)
	w.WriteUint16(m.Reserved2)
	return w.Bytes(), w.Error()
}

func (m *CallDequeuedEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.BaseCallEvent.readFrom(r)
	m.ServiceNumber = r.ReadUint32()
	m.ServiceID = r.ReadUint32()
	m.QueueDeviceType = r.ReadUint16()
	m.LocalConnectionState = r.ReadUint16()
	m.EventCause = r.ReadUint16()
	m.Reserved2 = r.ReadUint16()
	return r.Error()
}
