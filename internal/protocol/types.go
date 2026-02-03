package protocol

// FloatingField represents a variable-length field in the floating part of a message.
type FloatingField struct {
	Tag    uint16
	Length uint8
	Data   []byte
}

// ConnectionID uniquely identifies a call connection.
type ConnectionID struct {
	CallID         uint32
	DeviceIDType   uint16
	DeviceIDLength uint16
	DeviceID       string
}

// ConnectionDeviceIDType constants.
const (
	DeviceIDTypeDynamic  uint16 = 0
	DeviceIDTypeStatic   uint16 = 1
	DeviceIDTypeExternal uint16 = 2
)

// CallType values.
const (
	CallTypeInbound       uint16 = 1
	CallTypeOutbound      uint16 = 2
	CallTypeOutboundPrivate uint16 = 3
	CallTypeInternal      uint16 = 4
)

// CallTypeName returns a human-readable name for a call type.
func CallTypeName(callType uint16) string {
	switch callType {
	case CallTypeInbound:
		return "Inbound"
	case CallTypeOutbound:
		return "Outbound"
	case CallTypeOutboundPrivate:
		return "OutboundPrivate"
	case CallTypeInternal:
		return "Internal"
	default:
		return "Unknown"
	}
}

// LocalConnectionState values.
const (
	ConnectionStateNull       uint16 = 0
	ConnectionStateInitiated  uint16 = 1
	ConnectionStateAlerting   uint16 = 2
	ConnectionStateConnected  uint16 = 3
	ConnectionStateHeld       uint16 = 4
	ConnectionStateQueued     uint16 = 5
	ConnectionStateFailed     uint16 = 6
)

// ConnectionStateName returns a human-readable name for a connection state.
func ConnectionStateName(state uint16) string {
	switch state {
	case ConnectionStateNull:
		return "Null"
	case ConnectionStateInitiated:
		return "Initiated"
	case ConnectionStateAlerting:
		return "Alerting"
	case ConnectionStateConnected:
		return "Connected"
	case ConnectionStateHeld:
		return "Held"
	case ConnectionStateQueued:
		return "Queued"
	case ConnectionStateFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}

// EventCause values.
const (
	CauseNone                 uint16 = 0
	CauseActiveMonitor        uint16 = 1
	CauseAlternate            uint16 = 2
	CauseBusy                 uint16 = 3
	CauseCallCancelled        uint16 = 4
	CauseCallForwardAlways    uint16 = 5
	CauseCallForwardBusy      uint16 = 6
	CauseCallForwardNoAnswer  uint16 = 7
	CauseCallNotAnswered      uint16 = 8
	CauseCallPickup           uint16 = 9
	CauseConference           uint16 = 10
	CauseConsult              uint16 = 11
	CauseDestNotObtainable    uint16 = 12
	CauseDoNotDisturb         uint16 = 13
	CauseIncompatibleDest     uint16 = 14
	CauseNetworkCongestion    uint16 = 15
	CauseNetworkNotObtainable uint16 = 16
	CauseNewCall              uint16 = 17
	CauseNoAvailableAgents    uint16 = 18
	CauseNormalClearing       uint16 = 19
	CauseOverflow             uint16 = 20
	CausePark                 uint16 = 21
	CauseRecall               uint16 = 22
	CauseRedirect             uint16 = 23
	CauseReorderTone          uint16 = 24
	CauseResourceNotAvailable uint16 = 25
	CauseSilentMonitor        uint16 = 26
	CauseTransfer             uint16 = 27
)

// PeripheralType values.
const (
	PeripheralTypeACD       uint16 = 1
	PeripheralTypePBX       uint16 = 2
	PeripheralTypeVRU       uint16 = 3
	PeripheralTypeVoiceMail uint16 = 4
)

// SystemEventID values.
const (
	SystemEventCentralControllerOnline  uint32 = 1
	SystemEventCentralControllerOffline uint32 = 2
	SystemEventPeripheralOnline         uint32 = 3
	SystemEventPeripheralOffline        uint32 = 4
	SystemEventCTIServerOffline         uint32 = 5
	SystemEventCTIServerOnline          uint32 = 6
	SystemEventHalfHourChange           uint32 = 7
	SystemEventInstrumentOutOfService   uint32 = 8
	SystemEventInstrumentBackInService  uint32 = 9
)

// SystemEventName returns a human-readable name for a system event.
func SystemEventName(eventID uint32) string {
	switch eventID {
	case SystemEventCentralControllerOnline:
		return "CentralControllerOnline"
	case SystemEventCentralControllerOffline:
		return "CentralControllerOffline"
	case SystemEventPeripheralOnline:
		return "PeripheralOnline"
	case SystemEventPeripheralOffline:
		return "PeripheralOffline"
	case SystemEventCTIServerOffline:
		return "CTIServerOffline"
	case SystemEventCTIServerOnline:
		return "CTIServerOnline"
	case SystemEventHalfHourChange:
		return "HalfHourChange"
	case SystemEventInstrumentOutOfService:
		return "InstrumentOutOfService"
	case SystemEventInstrumentBackInService:
		return "InstrumentBackInService"
	default:
		return "Unknown"
	}
}
