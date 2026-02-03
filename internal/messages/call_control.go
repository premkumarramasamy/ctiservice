package messages

import (
	"ctiservice/internal/protocol"
)

// ConsultCallReq is sent to initiate a consult call.
// Protocol Version 24 - CONSULT_CALL_REQ (MessageType = 50)
type ConsultCallReq struct {
	// Fixed Part
	InvokeID               uint32 // Client-assigned request ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	ActiveConnectionCallID uint32 // Active call ID to consult from (UINT)
	ActiveConnectionType   uint16 // Active connection device type (USHORT)
	ConsultType            uint16 // Type of consult (USHORT)
	Reserved               uint32 // Reserved (UINT)

	// Floating fields
	ActiveConnectionDeviceID string // Tag 31
	ConsultedDeviceID        string // Tag 45
	ANI                      string // Tag 15
	UserToUserInfo           string // Tag 17
	CallVariable1            string // Tag 18
	CallVariable2            string // Tag 19
	CallVariable3            string // Tag 20
	CallVariable4            string // Tag 21
	CallVariable5            string // Tag 22
	CallVariable6            string // Tag 23
	CallVariable7            string // Tag 24
	CallVariable8            string // Tag 25
	CallVariable9            string // Tag 26
	CallVariable10           string // Tag 27
}

func (m *ConsultCallReq) Type() uint32 {
	return protocol.MsgTypeConsultCallReq
}

func (m *ConsultCallReq) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint32(m.ActiveConnectionCallID)
	w.WriteUint16(m.ActiveConnectionType)
	w.WriteUint16(m.ConsultType)
	w.WriteUint32(m.Reserved)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.ActiveConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ActiveConnectionDeviceID)
	}
	if m.ConsultedDeviceID != "" {
		fw.WriteString(protocol.TagConsultedDeviceID, m.ConsultedDeviceID)
	}
	if m.ANI != "" {
		fw.WriteString(protocol.TagANI, m.ANI)
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

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *ConsultCallReq) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.ActiveConnectionCallID = r.ReadUint32()
	m.ActiveConnectionType = r.ReadUint16()
	m.ConsultType = r.ReadUint16()
	m.Reserved = r.ReadUint32()

	if err := r.Error(); err != nil {
		return err
	}

	// Parse floating fields
	if r.Remaining() > 0 {
		ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
		if err != nil {
			return err
		}
		m.ActiveConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.ConsultedDeviceID = ff.GetString(protocol.TagConsultedDeviceID)
		m.ANI = ff.GetString(protocol.TagANI)
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
	}

	return nil
}

// ConsultCallConf is the server's response to ConsultCallReq.
// Protocol Version 24 - CONSULT_CALL_CONF (MessageType = 51)
type ConsultCallConf struct {
	// Fixed Part
	InvokeID                uint32 // Matches ConsultCallReq InvokeID (UINT)
	NewConnectionCallID     uint32 // New call ID for consult call (UINT)
	NewConnectionDeviceType uint16 // New connection device type (USHORT)
	LineHandle              uint16 // Line handle (USHORT)
	LineType                uint16 // Line type (USHORT)
	Reserved                uint16 // Reserved (USHORT)

	// Floating fields
	NewConnectionDeviceID string // Tag 186
}

func (m *ConsultCallConf) Type() uint32 {
	return protocol.MsgTypeConsultCallConf
}

func (m *ConsultCallConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.NewConnectionCallID)
	w.WriteUint16(m.NewConnectionDeviceType)
	w.WriteUint16(m.LineHandle)
	w.WriteUint16(m.LineType)
	w.WriteUint16(m.Reserved)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.NewConnectionDeviceID != "" {
		fw.WriteString(protocol.TagNewConnectionDeviceID, m.NewConnectionDeviceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *ConsultCallConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.NewConnectionCallID = r.ReadUint32()
	m.NewConnectionDeviceType = r.ReadUint16()
	m.LineHandle = r.ReadUint16()
	m.LineType = r.ReadUint16()
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
		m.NewConnectionDeviceID = ff.GetString(protocol.TagNewConnectionDeviceID)
	}

	return nil
}

// ConferenceCallReq is sent to create a conference call.
// Protocol Version 24 - CONFERENCE_CALL_REQ (MessageType = 48)
type ConferenceCallReq struct {
	// Fixed Part
	InvokeID                  uint32 // Client-assigned request ID (UINT)
	PeripheralID              uint32 // Peripheral ID (UINT)
	ActiveConnectionCallID    uint32 // Active call ID (UINT)
	ActiveConnectionType      uint16 // Active connection device type (USHORT)
	HeldConnectionCallID      uint32 // Held call ID (UINT)
	HeldConnectionType        uint16 // Held connection device type (USHORT)
	Reserved                  uint16 // Reserved (USHORT)

	// Floating fields
	ActiveConnectionDeviceID string // Tag 31
	HeldConnectionDeviceID   string // Tag 34
}

func (m *ConferenceCallReq) Type() uint32 {
	return protocol.MsgTypeConferenceCallReq
}

func (m *ConferenceCallReq) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint32(m.ActiveConnectionCallID)
	w.WriteUint16(m.ActiveConnectionType)
	w.WriteUint32(m.HeldConnectionCallID)
	w.WriteUint16(m.HeldConnectionType)
	w.WriteUint16(m.Reserved)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.ActiveConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ActiveConnectionDeviceID)
	}
	if m.HeldConnectionDeviceID != "" {
		fw.WriteString(protocol.TagHoldingDeviceID, m.HeldConnectionDeviceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *ConferenceCallReq) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.ActiveConnectionCallID = r.ReadUint32()
	m.ActiveConnectionType = r.ReadUint16()
	m.HeldConnectionCallID = r.ReadUint32()
	m.HeldConnectionType = r.ReadUint16()
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
		m.ActiveConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.HeldConnectionDeviceID = ff.GetString(protocol.TagHoldingDeviceID)
	}

	return nil
}

// ConferenceCallConf is the server's response to ConferenceCallReq.
// Protocol Version 24 - CONFERENCE_CALL_CONF (MessageType = 49)
type ConferenceCallConf struct {
	// Fixed Part
	InvokeID                uint32 // Matches ConferenceCallReq InvokeID (UINT)
	NewConnectionCallID     uint32 // New conference call ID (UINT)
	NewConnectionDeviceType uint16 // New connection device type (USHORT)
	LineHandle              uint16 // Line handle (USHORT)
	LineType                uint16 // Line type (USHORT)
	Reserved                uint16 // Reserved (USHORT)

	// Floating fields
	NewConnectionDeviceID string // Tag 186
}

func (m *ConferenceCallConf) Type() uint32 {
	return protocol.MsgTypeConferenceCallConf
}

func (m *ConferenceCallConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.NewConnectionCallID)
	w.WriteUint16(m.NewConnectionDeviceType)
	w.WriteUint16(m.LineHandle)
	w.WriteUint16(m.LineType)
	w.WriteUint16(m.Reserved)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.NewConnectionDeviceID != "" {
		fw.WriteString(protocol.TagNewConnectionDeviceID, m.NewConnectionDeviceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *ConferenceCallConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.NewConnectionCallID = r.ReadUint32()
	m.NewConnectionDeviceType = r.ReadUint16()
	m.LineHandle = r.ReadUint16()
	m.LineType = r.ReadUint16()
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
		m.NewConnectionDeviceID = ff.GetString(protocol.TagNewConnectionDeviceID)
	}

	return nil
}

// TransferCallReq is sent to transfer a call.
// Protocol Version 24 - TRANSFER_CALL_REQ (MessageType = 64)
type TransferCallReq struct {
	// Fixed Part
	InvokeID                  uint32 // Client-assigned request ID (UINT)
	PeripheralID              uint32 // Peripheral ID (UINT)
	ActiveConnectionCallID    uint32 // Active call ID (UINT)
	ActiveConnectionType      uint16 // Active connection device type (USHORT)
	HeldConnectionCallID      uint32 // Held call ID (UINT)
	HeldConnectionType        uint16 // Held connection device type (USHORT)
	Reserved                  uint16 // Reserved (USHORT)

	// Floating fields
	ActiveConnectionDeviceID string // Tag 31
	HeldConnectionDeviceID   string // Tag 34
}

func (m *TransferCallReq) Type() uint32 {
	return protocol.MsgTypeTransferCallReq
}

func (m *TransferCallReq) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint32(m.ActiveConnectionCallID)
	w.WriteUint16(m.ActiveConnectionType)
	w.WriteUint32(m.HeldConnectionCallID)
	w.WriteUint16(m.HeldConnectionType)
	w.WriteUint16(m.Reserved)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.ActiveConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ActiveConnectionDeviceID)
	}
	if m.HeldConnectionDeviceID != "" {
		fw.WriteString(protocol.TagHoldingDeviceID, m.HeldConnectionDeviceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *TransferCallReq) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.ActiveConnectionCallID = r.ReadUint32()
	m.ActiveConnectionType = r.ReadUint16()
	m.HeldConnectionCallID = r.ReadUint32()
	m.HeldConnectionType = r.ReadUint16()
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
		m.ActiveConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
		m.HeldConnectionDeviceID = ff.GetString(protocol.TagHoldingDeviceID)
	}

	return nil
}

// TransferCallConf is the server's response to TransferCallReq.
// Protocol Version 24 - TRANSFER_CALL_CONF (MessageType = 65)
type TransferCallConf struct {
	// Fixed Part
	InvokeID                uint32 // Matches TransferCallReq InvokeID (UINT)
	NewConnectionCallID     uint32 // New transferred call ID (UINT)
	NewConnectionDeviceType uint16 // New connection device type (USHORT)
	LineHandle              uint16 // Line handle (USHORT)
	LineType                uint16 // Line type (USHORT)
	Reserved                uint16 // Reserved (USHORT)

	// Floating fields
	NewConnectionDeviceID string // Tag 186
}

func (m *TransferCallConf) Type() uint32 {
	return protocol.MsgTypeTransferCallConf
}

func (m *TransferCallConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.NewConnectionCallID)
	w.WriteUint16(m.NewConnectionDeviceType)
	w.WriteUint16(m.LineHandle)
	w.WriteUint16(m.LineType)
	w.WriteUint16(m.Reserved)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.NewConnectionDeviceID != "" {
		fw.WriteString(protocol.TagNewConnectionDeviceID, m.NewConnectionDeviceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *TransferCallConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.NewConnectionCallID = r.ReadUint32()
	m.NewConnectionDeviceType = r.ReadUint16()
	m.LineHandle = r.ReadUint16()
	m.LineType = r.ReadUint16()
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
		m.NewConnectionDeviceID = ff.GetString(protocol.TagNewConnectionDeviceID)
	}

	return nil
}

// HoldCallReq is sent to place a call on hold.
// Protocol Version 24 - HOLD_CALL_REQ (MessageType = 54)
type HoldCallReq struct {
	// Fixed Part
	InvokeID               uint32 // Client-assigned request ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	ConnectionCallID       uint32 // Call ID to hold (UINT)
	ConnectionDeviceIDType uint16 // Connection device type (USHORT)
	Reserved               uint16 // Reserved (USHORT)

	// Floating fields
	ConnectionDeviceID string // Tag 31
}

func (m *HoldCallReq) Type() uint32 {
	return protocol.MsgTypeHoldCallReq
}

func (m *HoldCallReq) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint16(m.Reserved)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *HoldCallReq) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.ConnectionCallID = r.ReadUint32()
	m.ConnectionDeviceIDType = r.ReadUint16()
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
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
	}

	return nil
}

// HoldCallConf is the server's response to HoldCallReq.
// Protocol Version 24 - HOLD_CALL_CONF (MessageType = 55)
type HoldCallConf struct {
	// Fixed Part
	InvokeID uint32 // Matches HoldCallReq InvokeID (UINT)
}

func (m *HoldCallConf) Type() uint32 {
	return protocol.MsgTypeHoldCallConf
}

func (m *HoldCallConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	return w.Bytes(), w.Error()
}

func (m *HoldCallConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	return r.Error()
}

// RetrieveCallReq is sent to retrieve a held call.
// Protocol Version 24 - RETRIEVE_CALL_REQ (MessageType = 62)
type RetrieveCallReq struct {
	// Fixed Part
	InvokeID               uint32 // Client-assigned request ID (UINT)
	PeripheralID           uint32 // Peripheral ID (UINT)
	ConnectionCallID       uint32 // Call ID to retrieve (UINT)
	ConnectionDeviceIDType uint16 // Connection device type (USHORT)
	Reserved               uint16 // Reserved (USHORT)

	// Floating fields
	ConnectionDeviceID string // Tag 31
}

func (m *RetrieveCallReq) Type() uint32 {
	return protocol.MsgTypeRetrieveCallReq
}

func (m *RetrieveCallReq) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	w.WriteUint32(m.PeripheralID)
	w.WriteUint32(m.ConnectionCallID)
	w.WriteUint16(m.ConnectionDeviceIDType)
	w.WriteUint16(m.Reserved)

	if err := w.Error(); err != nil {
		return nil, err
	}

	// Add floating fields
	fw := protocol.NewFloatingFieldWriter()
	if m.ConnectionDeviceID != "" {
		fw.WriteString(protocol.TagConnectionDeviceID, m.ConnectionDeviceID)
	}

	fixed := w.Bytes()
	floating := fw.Bytes()
	result := make([]byte, len(fixed)+len(floating))
	copy(result, fixed)
	copy(result[len(fixed):], floating)

	return result, nil
}

func (m *RetrieveCallReq) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	m.PeripheralID = r.ReadUint32()
	m.ConnectionCallID = r.ReadUint32()
	m.ConnectionDeviceIDType = r.ReadUint16()
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
		m.ConnectionDeviceID = ff.GetString(protocol.TagConnectionDeviceID)
	}

	return nil
}

// RetrieveCallConf is the server's response to RetrieveCallReq.
// Protocol Version 24 - RETRIEVE_CALL_CONF (MessageType = 63)
type RetrieveCallConf struct {
	// Fixed Part
	InvokeID uint32 // Matches RetrieveCallReq InvokeID (UINT)
}

func (m *RetrieveCallConf) Type() uint32 {
	return protocol.MsgTypeRetrieveCallConf
}

func (m *RetrieveCallConf) Encode() ([]byte, error) {
	w := protocol.NewFixedFieldWriter()
	w.WriteUint32(m.InvokeID)
	return w.Bytes(), w.Error()
}

func (m *RetrieveCallConf) Decode(data []byte) error {
	r := protocol.NewFixedFieldReader(data)
	m.InvokeID = r.ReadUint32()
	return r.Error()
}
