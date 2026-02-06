package messages

import (
	"ctiservice/internal/protocol"
	"fmt"
)

// Registry provides message parsing by type ID.
type Registry struct{}

// NewRegistry creates a new message registry.
func NewRegistry() *Registry {
	return &Registry{}
}

// Parse creates and decodes a message from its type ID and body data.
func (r *Registry) Parse(msgType uint32, data []byte) (protocol.Message, error) {
	msg := r.Create(msgType)
	if msg == nil {
		return nil, fmt.Errorf("unknown message type: %d (%s)",
			msgType, protocol.MessageTypeName(msgType))
	}

	if err := msg.Decode(data); err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w",
			protocol.MessageTypeName(msgType), err)
	}

	return msg, nil
}

// Create instantiates a new message of the given type.
// Returns nil for unknown message types.
func (r *Registry) Create(msgType uint32) protocol.Message {
	switch msgType {
	// Session management
	case protocol.MsgTypeOpenReq:
		return &OpenReq{}
	case protocol.MsgTypeOpenConf:
		return &OpenConf{}
	case protocol.MsgTypeHeartbeatReq:
		return &HeartbeatReq{}
	case protocol.MsgTypeHeartbeatConf:
		return &HeartbeatConf{}
	case protocol.MsgTypeCloseReq:
		return &CloseReq{}
	case protocol.MsgTypeCloseConf:
		return &CloseConf{}

	// Error messages
	case protocol.MsgTypeFailureConf:
		return &FailureConf{}
	case protocol.MsgTypeFailureEvent:
		return &FailureEvent{}

	// System events
	case protocol.MsgTypeSystemEvent:
		return &SystemEvent{}

	// Call events
	case protocol.MsgTypeBeginCallEvent:
		return &BeginCallEvent{}
	case protocol.MsgTypeEndCallEvent:
		return &EndCallEvent{}
	case protocol.MsgTypeCallDataUpdateEvent:
		return &CallDataUpdateEvent{}
	case protocol.MsgTypeCallDeliveredEvent:
		return &CallDeliveredEvent{}
	case protocol.MsgTypeCallEstablishedEvent:
		return &CallEstablishedEvent{}
	case protocol.MsgTypeCallHeldEvent:
		return &CallHeldEvent{}
	case protocol.MsgTypeCallRetrievedEvent:
		return &CallRetrievedEvent{}
	case protocol.MsgTypeCallClearedEvent:
		return &CallClearedEvent{}
	case protocol.MsgTypeCallConnectionClearedEvent:
		return &CallConnectionClearedEvent{}
	case protocol.MsgTypeCallOriginatedEvent:
		return &CallOriginatedEvent{}
	case protocol.MsgTypeCallFailedEvent:
		return &CallFailedEvent{}
	case protocol.MsgTypeCallConferencedEvent:
		return &CallConferencedEvent{}
	case protocol.MsgTypeCallTransferredEvent:
		return &CallTransferredEvent{}
	case protocol.MsgTypeCallQueuedEvent:
		return &CallQueuedEvent{}
	case protocol.MsgTypeCallDequeuedEvent:
		return &CallDequeuedEvent{}
	case protocol.MsgTypeCallServiceInitiatedEvent:
		return &CallServiceInitiatedEvent{}

	// Agent events
	case protocol.MsgTypeAgentStateEvent:
		return &AgentStateEvent{}
	case protocol.MsgTypeAgentPreCallEvent:
		return &AgentPreCallEvent{}
	case protocol.MsgTypeAgentPreCallAbortEvent:
		return &AgentPreCallAbortEvent{}

	// Call control messages
	case protocol.MsgTypeConsultCallReq:
		return &ConsultCallReq{}
	case protocol.MsgTypeConsultCallConf:
		return &ConsultCallConf{}
	case protocol.MsgTypeConferenceCallReq:
		return &ConferenceCallReq{}
	case protocol.MsgTypeConferenceCallConf:
		return &ConferenceCallConf{}
	case protocol.MsgTypeTransferCallReq:
		return &TransferCallReq{}
	case protocol.MsgTypeTransferCallConf:
		return &TransferCallConf{}
	case protocol.MsgTypeHoldCallReq:
		return &HoldCallReq{}
	case protocol.MsgTypeHoldCallConf:
		return &HoldCallConf{}
	case protocol.MsgTypeRetrieveCallReq:
		return &RetrieveCallReq{}
	case protocol.MsgTypeRetrieveCallConf:
		return &RetrieveCallConf{}

	// Supervisor events
	case protocol.MsgTypeSupervisorAssistEvent:
		return &SupervisorAssistEvent{}

	// Config events
	case protocol.MsgTypeConfigAgentEvent:
		return &ConfigAgentEvent{}
	case protocol.MsgTypeConfigDeviceEvent:
		return &ConfigDeviceEvent{}
	case protocol.MsgTypeConfigCSQEvent:
		return &ConfigCSQEvent{}
	case protocol.MsgTypeConfigBeginEvent:
		return &ConfigBeginEvent{}
	case protocol.MsgTypeConfigEndEvent:
		return &ConfigEndEvent{}
	case protocol.MsgTypeConfigRequestEvent:
		return &ConfigRequestEvent{}

	default:
		// Return a generic message for unknown types
		return &GenericMessage{msgType: msgType}
	}
}

// GenericMessage holds raw data for unknown message types.
type GenericMessage struct {
	msgType uint32
	Data    []byte
}

func (m *GenericMessage) Type() uint32 {
	return m.msgType
}

func (m *GenericMessage) Encode() ([]byte, error) {
	return m.Data, nil
}

func (m *GenericMessage) Decode(data []byte) error {
	m.Data = make([]byte, len(data))
	copy(m.Data, data)
	return nil
}
