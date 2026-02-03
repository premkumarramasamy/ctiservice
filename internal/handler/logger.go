package handler

import (
	"ctiservice/internal/messages"
	"ctiservice/internal/protocol"
	"log/slog"
)

// LogHandler logs all received events.
type LogHandler struct {
	logger *slog.Logger
}

// NewLogHandler creates a new logging event handler.
func NewLogHandler(logger *slog.Logger) *LogHandler {
	return &LogHandler{logger: logger}
}

// Handle logs the received message with appropriate details.
func (h *LogHandler) Handle(msg protocol.Message) {
	msgType := msg.Type()
	msgName := protocol.MessageTypeName(msgType)

	// Create base attributes
	attrs := []any{
		"messageType", msgName,
		"messageTypeID", msgType,
	}

	switch m := msg.(type) {
	case *messages.BeginCallEvent:
		h.logger.Info("call started",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"callType", protocol.CallTypeName(m.CallType),
				"connectionState", protocol.ConnectionStateName(m.ConnectionState),
				"ANI", m.ANI,
				"DNIS", m.DNIS,
				"callingDeviceID", m.CallingDeviceID,
				"calledDeviceID", m.CalledDeviceID,
				"serviceNumber", m.ServiceNumber,
				"skillGroupNumber", m.SkillGroupNumber,
			)...)

	case *messages.EndCallEvent:
		h.logger.Info("call ended",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"connectionState", protocol.ConnectionStateName(m.ConnectionState),
			)...)

	case *messages.CallDeliveredEvent:
		h.logger.Info("call delivered",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"ANI", m.ANI,
				"DNIS", m.DNIS,
				"eventCause", m.EventCause,
			)...)

	case *messages.CallEstablishedEvent:
		h.logger.Info("call established",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"ANI", m.ANI,
				"DNIS", m.DNIS,
				"eventCause", m.EventCause,
			)...)

	case *messages.CallHeldEvent:
		h.logger.Info("call held",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
			)...)

	case *messages.CallRetrievedEvent:
		h.logger.Info("call retrieved",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
			)...)

	case *messages.CallClearedEvent:
		h.logger.Info("call cleared",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
			)...)

	case *messages.CallConnectionClearedEvent:
		h.logger.Info("call connection cleared",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
			)...)

	case *messages.CallOriginatedEvent:
		h.logger.Info("call originated",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"callingDeviceID", m.CallingDeviceID,
				"calledDeviceID", m.CalledDeviceID,
			)...)

	case *messages.CallFailedEvent:
		h.logger.Warn("call failed",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
			)...)

	case *messages.CallConferencedEvent:
		h.logger.Info("call conferenced",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"primaryCallID", m.PrimaryCallID,
				"secondaryCallID", m.SecondaryCallID,
			)...)

	case *messages.CallTransferredEvent:
		h.logger.Info("call transferred",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"primaryCallID", m.PrimaryCallID,
				"secondaryCallID", m.SecondaryCallID,
			)...)

	case *messages.CallQueuedEvent:
		h.logger.Info("call queued",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"serviceNumber", m.ServiceNumber,
				"skillGroupNumber", m.SkillGroupNumber,
			)...)

	case *messages.CallDequeuedEvent:
		h.logger.Info("call dequeued",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"serviceNumber", m.ServiceNumber,
			)...)

	case *messages.CallDataUpdateEvent:
		h.logger.Info("call data updated",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"ANI", m.ANI,
				"DNIS", m.DNIS,
			)...)

	case *messages.AgentStateEvent:
		h.logger.Info("agent state changed",
			append(attrs,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"sessionID", m.SessionID,
				"agentState", m.StateName(),
				"agentStateCode", m.AgentState,
				"reasonCode", m.EventReasonCode,
				"stateDuration", m.StateDuration,
				"agentID", m.AgentID,
				"agentExtension", m.AgentExtension,
				"skillGroupNumber", m.SkillGroupNumber,
				"icmAgentID", m.ICMAgentID,
			)...)

	case *messages.SystemEvent:
		h.logger.Info("system event",
			append(attrs,
				"eventID", m.SystemEventID,
				"eventName", m.EventName(),
				"pgStatus", m.PGStatus,
				"arg1", m.SystemEventArg1,
				"arg2", m.SystemEventArg2,
				"arg3", m.SystemEventArg3,
			)...)

	case *messages.FailureConf:
		h.logger.Error("failure confirmation",
			append(attrs,
				"invokeID", m.InvokeID,
				"status", m.Status,
			)...)

	case *messages.FailureEvent:
		h.logger.Error("failure event",
			append(attrs,
				"status", m.Status,
			)...)

	case *messages.GenericMessage:
		h.logger.Debug("unknown message received",
			append(attrs,
				"dataLength", len(m.Data),
			)...)

	default:
		h.logger.Info("event received", attrs...)
	}
}
