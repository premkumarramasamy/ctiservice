// Package protocol implements the GED-188 CTI protocol encoding/decoding.
package protocol

// Message Type IDs as defined in GED-188 protocol specification.
const (
	// Error messages
	MsgTypeFailureConf  uint32 = 1
	MsgTypeFailureEvent uint32 = 2

	// Session management
	MsgTypeOpenReq      uint32 = 3
	MsgTypeOpenConf     uint32 = 4
	MsgTypeHeartbeatReq uint32 = 5
	MsgTypeHeartbeatConf uint32 = 6
	MsgTypeCloseReq     uint32 = 7
	MsgTypeCloseConf    uint32 = 8

	// Call events
	MsgTypeCallDeliveredEvent         uint32 = 9
	MsgTypeCallEstablishedEvent       uint32 = 10
	MsgTypeCallHeldEvent              uint32 = 11
	MsgTypeCallRetrievedEvent         uint32 = 12
	MsgTypeCallClearedEvent           uint32 = 13
	MsgTypeCallConnectionClearedEvent uint32 = 14
	MsgTypeCallOriginatedEvent        uint32 = 15
	MsgTypeCallFailedEvent            uint32 = 16
	MsgTypeCallConferencedEvent       uint32 = 17
	MsgTypeCallTransferredEvent       uint32 = 18
	MsgTypeCallDivertedEvent          uint32 = 19
	MsgTypeCallServiceInitiatedEvent  uint32 = 20
	MsgTypeCallQueuedEvent            uint32 = 21
	MsgTypeBeginCallEvent             uint32 = 23
	MsgTypeEndCallEvent               uint32 = 24
	MsgTypeCallDataUpdateEvent        uint32 = 25
	MsgTypeSetCallDataReq             uint32 = 26
	MsgTypeSetCallDataConf            uint32 = 27

	// Agent state
	MsgTypeAgentStateEvent     uint32 = 30
	MsgTypeSystemEvent         uint32 = 31
	MsgTypeControlFailureConf  uint32 = 35
	MsgTypeQueryAgentStateReq  uint32 = 36
	MsgTypeQueryAgentStateConf uint32 = 37
	MsgTypeSetAgentStateReq    uint32 = 38
	MsgTypeSetAgentStateConf   uint32 = 39

	// Call control
	MsgTypeAlternateCallReq       uint32 = 40
	MsgTypeAlternateCallConf      uint32 = 41
	MsgTypeAnswerCallReq          uint32 = 42
	MsgTypeAnswerCallConf         uint32 = 43
	MsgTypeClearCallReq           uint32 = 44
	MsgTypeClearCallConf          uint32 = 45
	MsgTypeClearConnectionReq     uint32 = 46
	MsgTypeClearConnectionConf    uint32 = 47
	MsgTypeConferenceCallReq      uint32 = 48
	MsgTypeConferenceCallConf     uint32 = 49
	MsgTypeConsultCallReq         uint32 = 50
	MsgTypeConsultCallConf        uint32 = 51
	MsgTypeHoldCallReq            uint32 = 54
	MsgTypeHoldCallConf           uint32 = 55
	MsgTypeMakeCallReq            uint32 = 56
	MsgTypeMakeCallConf           uint32 = 57
	MsgTypeReconnectCallReq       uint32 = 60
	MsgTypeReconnectCallConf      uint32 = 61
	MsgTypeRetrieveCallReq        uint32 = 62
	MsgTypeRetrieveCallConf       uint32 = 63
	MsgTypeTransferCallReq        uint32 = 64
	MsgTypeTransferCallConf       uint32 = 65

	// Device queries
	MsgTypeQueryDeviceInfoReq  uint32 = 78
	MsgTypeQueryDeviceInfoConf uint32 = 79
	MsgTypeSnapshotCallReq     uint32 = 82
	MsgTypeSnapshotCallConf    uint32 = 83
	MsgTypeSnapshotDeviceReq   uint32 = 84
	MsgTypeSnapshotDeviceConf  uint32 = 85
	MsgTypeCallDequeuedEvent   uint32 = 86

	// DTMF
	MsgTypeSendDTMFSignalReq  uint32 = 91
	MsgTypeSendDTMFSignalConf uint32 = 92

	// RTP events
	MsgTypeRTPStartedEvent uint32 = 116
	MsgTypeRTPStoppedEvent uint32 = 117

	// Supervisor
	MsgTypeSupervisorAssistReq   uint32 = 118
	MsgTypeSupervisorAssistConf  uint32 = 119
	MsgTypeSupervisorAssistEvent uint32 = 120
	MsgTypeSuperviseCallReq      uint32 = 124
	MsgTypeSuperviseCallConf     uint32 = 125

	// Bad call
	MsgTypeBadCallReq  uint32 = 139
	MsgTypeBadCallConf uint32 = 140

	// Queue statistics
	MsgTypeQueryQueueStatisticsReq  uint32 = 223
	MsgTypeQueryQueueStatisticsConf uint32 = 224
	MsgTypeQuerySummaryStatisticsReq  uint32 = 225
	MsgTypeQuerySummaryStatisticsConf uint32 = 226

	// Configuration
	MsgTypeConfigRequestKeyEvent         uint32 = 230
	MsgTypeConfigKeyEvent                uint32 = 231
	MsgTypeConfigRequestEvent            uint32 = 232
	MsgTypeConfigBeginEvent              uint32 = 233
	MsgTypeConfigEndEvent                uint32 = 234
	MsgTypeConfigApplicationEvent        uint32 = 235
	MsgTypeConfigCSQEvent                uint32 = 236
	MsgTypeConfigAgentEvent              uint32 = 237
	MsgTypeConfigDeviceEvent             uint32 = 238
	MsgTypeQueryAgentQueueStatisticsReq  uint32 = 239
	MsgTypeQueryAgentQueueStatisticsConf uint32 = 240
	MsgTypeTeamConfigReq                 uint32 = 242
	MsgTypeTeamConfigEvent               uint32 = 243
	MsgTypeTeamConfigConf                uint32 = 244
)

// Service mask bits for ServicesRequested field in OPEN_REQ.
const (
	ServiceClientEvents uint32 = 0x00000001 // Agent mode - receive events for client's device
	ServiceCallControl  uint32 = 0x00000002 // Call control capabilities
	ServiceAllEvents    uint32 = 0x00000010 // Bridge mode - receive all events
	ServiceSupervisor   uint32 = 0x00000020 // Supervisor capabilities
)

// Agent state values.
const (
	AgentStateLoggedOut   uint16 = 0
	AgentStateLoggedIn    uint16 = 1
	AgentStateNotReady    uint16 = 2
	AgentStateReady       uint16 = 3
	AgentStateTalking     uint16 = 4
	AgentStateWorkNotReady uint16 = 5
	AgentStateWorkReady   uint16 = 6
	AgentStateHold        uint16 = 7
	AgentStateReserved    uint16 = 8
	AgentStateUnknown     uint16 = 9
)

// AgentStateName returns a human-readable name for an agent state.
func AgentStateName(state uint16) string {
	switch state {
	case AgentStateLoggedOut:
		return "LoggedOut"
	case AgentStateLoggedIn:
		return "LoggedIn"
	case AgentStateNotReady:
		return "NotReady"
	case AgentStateReady:
		return "Ready"
	case AgentStateTalking:
		return "Talking"
	case AgentStateWorkNotReady:
		return "WorkNotReady"
	case AgentStateWorkReady:
		return "WorkReady"
	case AgentStateHold:
		return "Hold"
	case AgentStateReserved:
		return "Reserved"
	default:
		return "Unknown"
	}
}

// Status codes for failure messages.
const (
	StatusSuccess             uint32 = 0
	StatusInvalidRequest      uint32 = 1
	StatusInvalidState        uint32 = 2
	StatusInvalidSession      uint32 = 3
	StatusInvalidService      uint32 = 4
	StatusInvalidCallID       uint32 = 5
	StatusInvalidDeviceID     uint32 = 6
	StatusResourceBusy        uint32 = 7
	StatusResourceUnavailable uint32 = 8
	StatusProtocolError       uint32 = 9
	StatusInternalError       uint32 = 10
)

// Floating field tag IDs.
const (
	TagClientID           uint16 = 1
	TagClientPassword     uint16 = 2
	TagAgentExtension     uint16 = 3
	TagAgentID            uint16 = 4
	TagAgentInstrument    uint16 = 5
	TagPeripheralID       uint16 = 6
	TagServiceNumber      uint16 = 7
	TagServiceID          uint16 = 8
	TagSkillGroupNumber   uint16 = 9
	TagSkillGroupID       uint16 = 10
	TagSkillGroupPriority uint16 = 11
	TagCallingDeviceID    uint16 = 12
	TagCalledDeviceID     uint16 = 13
	TagLastRedirectDeviceID uint16 = 14
	TagANI                uint16 = 15
	TagDNIS               uint16 = 16
	TagUserToUserInfo     uint16 = 17
	TagCallVariable1      uint16 = 18
	TagCallVariable2      uint16 = 19
	TagCallVariable3      uint16 = 20
	TagCallVariable4      uint16 = 21
	TagCallVariable5      uint16 = 22
	TagCallVariable6      uint16 = 23
	TagCallVariable7      uint16 = 24
	TagCallVariable8      uint16 = 25
	TagCallVariable9      uint16 = 26
	TagCallVariable10     uint16 = 27
	TagCTIClientSignature uint16 = 28
	TagCTIClientTimestamp uint16 = 29
	TagCallWrapupData     uint16 = 30
	TagNamedVariable      uint16 = 82
	TagNamedArray         uint16 = 83
)

// Header size in bytes.
const HeaderSize = 8

// Maximum message body size.
const MaxMessageSize = 65536

// MessageTypeName returns a human-readable name for a message type.
func MessageTypeName(msgType uint32) string {
	switch msgType {
	case MsgTypeFailureConf:
		return "FAILURE_CONF"
	case MsgTypeFailureEvent:
		return "FAILURE_EVENT"
	case MsgTypeOpenReq:
		return "OPEN_REQ"
	case MsgTypeOpenConf:
		return "OPEN_CONF"
	case MsgTypeHeartbeatReq:
		return "HEARTBEAT_REQ"
	case MsgTypeHeartbeatConf:
		return "HEARTBEAT_CONF"
	case MsgTypeCloseReq:
		return "CLOSE_REQ"
	case MsgTypeCloseConf:
		return "CLOSE_CONF"
	case MsgTypeCallDeliveredEvent:
		return "CALL_DELIVERED_EVENT"
	case MsgTypeCallEstablishedEvent:
		return "CALL_ESTABLISHED_EVENT"
	case MsgTypeCallHeldEvent:
		return "CALL_HELD_EVENT"
	case MsgTypeCallRetrievedEvent:
		return "CALL_RETRIEVED_EVENT"
	case MsgTypeCallClearedEvent:
		return "CALL_CLEARED_EVENT"
	case MsgTypeCallConnectionClearedEvent:
		return "CALL_CONNECTION_CLEARED_EVENT"
	case MsgTypeCallOriginatedEvent:
		return "CALL_ORIGINATED_EVENT"
	case MsgTypeCallFailedEvent:
		return "CALL_FAILED_EVENT"
	case MsgTypeCallConferencedEvent:
		return "CALL_CONFERENCED_EVENT"
	case MsgTypeCallTransferredEvent:
		return "CALL_TRANSFERRED_EVENT"
	case MsgTypeCallDivertedEvent:
		return "CALL_DIVERTED_EVENT"
	case MsgTypeCallServiceInitiatedEvent:
		return "CALL_SERVICE_INITIATED_EVENT"
	case MsgTypeCallQueuedEvent:
		return "CALL_QUEUED_EVENT"
	case MsgTypeBeginCallEvent:
		return "BEGIN_CALL_EVENT"
	case MsgTypeEndCallEvent:
		return "END_CALL_EVENT"
	case MsgTypeCallDataUpdateEvent:
		return "CALL_DATA_UPDATE_EVENT"
	case MsgTypeAgentStateEvent:
		return "AGENT_STATE_EVENT"
	case MsgTypeSystemEvent:
		return "SYSTEM_EVENT"
	case MsgTypeCallDequeuedEvent:
		return "CALL_DEQUEUED_EVENT"
	case MsgTypeRTPStartedEvent:
		return "RTP_STARTED_EVENT"
	case MsgTypeRTPStoppedEvent:
		return "RTP_STOPPED_EVENT"
	default:
		return "UNKNOWN"
	}
}
