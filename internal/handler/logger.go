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
				"callType", m.CallType,
				"connectionDeviceIDType", m.ConnectionDeviceIDType,
				"ANI", m.ANI,
				"DNIS", m.DNIS,
				"connectionDeviceID", m.ConnectionDeviceID,
				"dialedNumber", m.DialedNumber,
			)...)

	case *messages.EndCallEvent:
		h.logger.Info("call ended",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"connectionDeviceID", m.ConnectionDeviceID,
			)...)

	case *messages.CallDeliveredEvent:
		h.logger.Info("call delivered",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"ANI", m.ANI,
				"DNIS", m.DNIS,
				"eventCause", m.EventCause,
				"localConnectionState", m.LocalConnectionState,
				"connectionDeviceID", m.ConnectionDeviceID,
			)...)

	case *messages.CallEstablishedEvent:
		h.logger.Info("call established",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"eventCause", m.EventCause,
				"localConnectionState", m.LocalConnectionState,
				"connectionDeviceID", m.ConnectionDeviceID,
				"answeringDeviceID", m.AnsweringDeviceID,
				"callingDeviceID", m.CallingDeviceID,
				"calledDeviceID", m.CalledDeviceID,
			)...)

	case *messages.CallHeldEvent:
		h.logger.Info("call held",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"eventCause", m.EventCause,
				"localConnectionState", m.LocalConnectionState,
				"connectionDeviceID", m.ConnectionDeviceID,
				"holdingDeviceID", m.HoldingDeviceID,
			)...)

	case *messages.CallRetrievedEvent:
		h.logger.Info("call retrieved",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"eventCause", m.EventCause,
				"localConnectionState", m.LocalConnectionState,
				"connectionDeviceID", m.ConnectionDeviceID,
				"retrievingDeviceID", m.RetrievingDeviceID,
			)...)

	case *messages.CallClearedEvent:
		h.logger.Info("call cleared",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
				"connectionDeviceID", m.ConnectionDeviceID,
			)...)

	case *messages.CallConnectionClearedEvent:
		h.logger.Info("call connection cleared",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
				"connectionDeviceID", m.ConnectionDeviceID,
			)...)

	case *messages.CallOriginatedEvent:
		h.logger.Info("call originated",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"callingDeviceID", m.CallingDeviceID,
				"calledDeviceID", m.CalledDeviceID,
				"connectionDeviceID", m.ConnectionDeviceID,
			)...)

	case *messages.CallFailedEvent:
		h.logger.Warn("call failed",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
				"connectionDeviceID", m.ConnectionDeviceID,
			)...)

	case *messages.CallConferencedEvent:
		h.logger.Info("call conferenced",
			append(attrs,
				"primaryCallID", m.PrimaryCallID,
				"secondaryCallID", m.SecondaryCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"numParties", m.NumParties,
				"eventCause", m.EventCause,
				"localConnectionState", m.LocalConnectionState,
				"primaryDeviceID", m.PrimaryDeviceID,
				"secondaryDeviceID", m.SecondaryDeviceID,
				"controllerDeviceID", m.ControllerDeviceID,
				"addedPartyDeviceID", m.AddedPartyDeviceID,
			)...)

	case *messages.CallTransferredEvent:
		h.logger.Info("call transferred",
			append(attrs,
				"primaryCallID", m.PrimaryCallID,
				"secondaryCallID", m.SecondaryCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"numParties", m.NumParties,
				"eventCause", m.EventCause,
				"localConnectionState", m.LocalConnectionState,
				"primaryDeviceID", m.PrimaryDeviceID,
				"secondaryDeviceID", m.SecondaryDeviceID,
				"transferringDeviceID", m.TransferringDeviceID,
				"transferredDeviceID", m.TransferredDeviceID,
			)...)

	case *messages.CallQueuedEvent:
		h.logger.Info("call queued",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
				"connectionDeviceID", m.ConnectionDeviceID,
			)...)

	case *messages.CallDequeuedEvent:
		h.logger.Info("call dequeued",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"eventCause", m.EventCause,
				"connectionDeviceID", m.ConnectionDeviceID,
			)...)

	case *messages.CallDataUpdateEvent:
		h.logger.Info("call data updated",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"callType", m.CallType,
				"newConnectionCallID", m.NewConnectionCallID,
				"calledPartyDisposition", m.CalledPartyDisposition,
				"campaignID", m.CampaignID,
				"queryRuleID", m.QueryRuleID,
				"connectionDeviceID", m.ConnectionDeviceID,
				"newConnectionDeviceID", m.NewConnectionDeviceID,
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

	case *messages.CallServiceInitiatedEvent:
		h.logger.Info("call service initiated",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"serviceNumber", m.ServiceNumber,
				"serviceID", m.ServiceID,
				"skillGroupNumber", m.SkillGroupNumber,
				"connectionDeviceID", m.ConnectionDeviceID,
				"callingDeviceID", m.CallingDeviceID,
			)...)

	case *messages.AgentPreCallEvent:
		h.logger.Info("agent pre-call notification",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"serviceNumber", m.ServiceNumber,
				"skillGroupNumber", m.SkillGroupNumber,
				"callType", m.CallType,
				"ANI", m.ANI,
				"DNIS", m.DNIS,
				"connectionDeviceID", m.ConnectionDeviceID,
				"preCallInvokeID", m.PreCallInvokeID,
			)...)

	case *messages.AgentPreCallAbortEvent:
		h.logger.Warn("agent pre-call aborted",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"eventCause", m.EventCause,
				"connectionDeviceID", m.ConnectionDeviceID,
				"preCallInvokeID", m.PreCallInvokeID,
			)...)

	case *messages.SupervisorAssistEvent:
		h.logger.Info("supervisor assist event",
			append(attrs,
				"callID", m.ConnectionCallID,
				"monitorID", m.MonitorID,
				"peripheralID", m.PeripheralID,
				"action", m.ActionName(),
				"actionCode", m.SupervisorAction,
				"eventCause", m.EventCause,
				"agentID", m.AgentID,
				"agentExtension", m.AgentExtension,
				"connectionDeviceID", m.ConnectionDeviceID,
			)...)

	case *messages.ConfigAgentEvent:
		h.logger.Info("config agent event",
			append(attrs,
				"peripheralID", m.PeripheralID,
				"operation", m.OperationName(),
				"numRecords", m.NumRecords,
				"agentID", m.AgentID,
				"agentExtension", m.AgentExtension,
				"loginID", m.LoginID,
				"firstName", m.FirstName,
				"lastName", m.LastName,
			)...)

	case *messages.ConfigDeviceEvent:
		h.logger.Info("config device event",
			append(attrs,
				"peripheralID", m.PeripheralID,
				"operation", m.OperationName(),
				"numRecords", m.NumRecords,
				"extension", m.Extension,
				"skillGroupID", m.SkillGroupID,
				"serviceID", m.ServiceID,
			)...)

	case *messages.ConfigCSQEvent:
		h.logger.Info("config CSQ event",
			append(attrs,
				"peripheralID", m.PeripheralID,
				"operation", m.OperationName(),
				"numRecords", m.NumRecords,
				"csqID", m.CSQID,
				"skillGroupID", m.SkillGroupID,
				"serviceID", m.ServiceID,
			)...)

	case *messages.ConfigBeginEvent:
		h.logger.Info("config begin",
			append(attrs,
				"peripheralID", m.PeripheralID,
				"configType", m.ConfigType,
			)...)

	case *messages.ConfigEndEvent:
		h.logger.Info("config end",
			append(attrs,
				"peripheralID", m.PeripheralID,
				"configType", m.ConfigType,
				"numRecords", m.NumRecords,
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
