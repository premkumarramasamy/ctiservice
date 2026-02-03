package messages

import (
	"ctiservice/internal/protocol"
)

// FailureConf is sent when a request fails.
type FailureConf struct {
	InvokeID uint32 // Matches the request's InvokeID
	Status   uint32 // Error status code
}

func (m *FailureConf) Type() uint32 {
	return protocol.MsgTypeFailureConf
}

func (m *FailureConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.Status)
	return w.Bytes(), w.Error()
}

func (m *FailureConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.Status = r.ReadUint32()
	return r.Error()
}

// FailureEvent is an unsolicited error notification.
type FailureEvent struct {
	Status uint32 // Error status code
}

func (m *FailureEvent) Type() uint32 {
	return protocol.MsgTypeFailureEvent
}

func (m *FailureEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.Status)
	return w.Bytes(), w.Error()
}

func (m *FailureEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.Status = r.ReadUint32()
	return r.Error()
}

// SystemEvent reports system status changes.
type SystemEvent struct {
	PGStatus              uint32 // Peripheral gateway status
	ICMCentralControllerTime uint32 // ICM controller timestamp
	SystemEventID         uint32 // Type of system event
	SystemEventArg1       uint32 // Event-specific argument 1
	SystemEventArg2       uint32 // Event-specific argument 2
	SystemEventArg3       uint32 // Event-specific argument 3
	EventDeviceType       uint16 // Device type involved
	Reserved              uint16 // Reserved
	ICMCentralController  uint32 // ICM central controller status
	Text                  string // Event description (floating)
}

func (m *SystemEvent) Type() uint32 {
	return protocol.MsgTypeSystemEvent
}

func (m *SystemEvent) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.PGStatus)
	w.WriteUint32(m.ICMCentralControllerTime)
	w.WriteUint32(m.SystemEventID)
	w.WriteUint32(m.SystemEventArg1)
	w.WriteUint32(m.SystemEventArg2)
	w.WriteUint32(m.SystemEventArg3)
	w.WriteUint16(m.EventDeviceType)
	w.WriteUint16(m.Reserved)
	w.WriteUint32(m.ICMCentralController)
	return w.Bytes(), w.Error()
}

func (m *SystemEvent) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.PGStatus = r.ReadUint32()
	m.ICMCentralControllerTime = r.ReadUint32()
	m.SystemEventID = r.ReadUint32()
	m.SystemEventArg1 = r.ReadUint32()
	m.SystemEventArg2 = r.ReadUint32()
	m.SystemEventArg3 = r.ReadUint32()
	m.EventDeviceType = r.ReadUint16()
	m.Reserved = r.ReadUint16()
	m.ICMCentralController = r.ReadUint32()

	if err := r.Error(); err != nil {
		return err
	}

	// No standard floating fields defined for SystemEvent, but parse any that exist
	return nil
}

// EventName returns the human-readable name for this system event.
func (m *SystemEvent) EventName() string {
	return protocol.SystemEventName(m.SystemEventID)
}
