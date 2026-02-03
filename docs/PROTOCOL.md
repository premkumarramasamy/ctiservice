# GED-188 Protocol Reference

This document describes the Cisco CTI Server Protocol (GED-188) as implemented in this service.

## Overview

GED-188 is a binary TCP protocol used to communicate with Cisco Unified Contact Center Enterprise (UCCE) CTI Servers. The protocol provides real-time call and agent state events.

**Default Ports:**
- Primary: 42027
- Secondary: 43027

## Message Structure

All messages follow this structure:

```
┌──────────────────┬──────────────────┬─────────────────────────────┐
│  MessageLength   │   MessageType    │        Message Body         │
│    (4 bytes)     │    (4 bytes)     │    (variable length)        │
├──────────────────┴──────────────────┼─────────────────────────────┤
│            Header (8 bytes)         │    Fixed + Floating Parts   │
└─────────────────────────────────────┴─────────────────────────────┘
```

### Byte Order

All multi-byte values are **big-endian** (network byte order).

### Message Header

| Field | Size | Description |
|-------|------|-------------|
| MessageLength | 4 bytes (uint32) | Length of message body (excludes header) |
| MessageType | 4 bytes (uint32) | Message type identifier |

### Message Body

The body consists of:
1. **Fixed Part**: Required fields in predetermined order
2. **Floating Part**: Optional variable-length fields (tag-length-value)

## Data Types

### Primitive Types

| Type | Size | Description |
|------|------|-------------|
| CHAR | 1 byte | Signed character |
| UCHAR | 1 byte | Unsigned character |
| SHORT | 2 bytes | Signed 16-bit integer |
| USHORT | 2 bytes | Unsigned 16-bit integer |
| INT | 4 bytes | Signed 32-bit integer |
| UINT | 4 bytes | Unsigned 32-bit integer |
| BOOL | 2 bytes | Boolean (0 = false, non-zero = true) |
| TIME | 4 bytes | Seconds since Jan 1, 1970 UTC |

### Floating Field Format (Protocol Version 24)

```
┌────────────┬──────────────┬─────────────────┐
│  FieldTag  │ FieldLength  │      Data       │
│ (2 bytes)  │  (2 bytes)   │  (FieldLength)  │
└────────────┴──────────────┴─────────────────┘
```

- **FieldTag**: USHORT identifying the field type (see Field Tags below)
- **FieldLength**: USHORT indicating number of data bytes (max 65535)
- **Data**: Field value, format depends on tag
- **Total overhead**: 4 bytes per floating field

Floating fields are packed contiguously. Message ends when all bytes consumed.

## Message Types

### Session Management

| Message | Type ID | Direction | Description |
|---------|---------|-----------|-------------|
| FAILURE_CONF | 1 | S→C | Request failed |
| FAILURE_EVENT | 2 | S→C | Unsolicited error |
| OPEN_REQ | 3 | C→S | Initialize session |
| OPEN_CONF | 4 | S→C | Session established |
| HEARTBEAT_REQ | 5 | C→S | Keepalive request |
| HEARTBEAT_CONF | 6 | S→C | Keepalive response |
| CLOSE_REQ | 7 | C→S | Close session |
| CLOSE_CONF | 8 | S→C | Session closed |

### Call Events

| Message | Type ID | Description |
|---------|---------|-------------|
| CALL_DELIVERED_EVENT | 9 | Call arrived at device |
| CALL_ESTABLISHED_EVENT | 10 | Call answered |
| CALL_HELD_EVENT | 11 | Call placed on hold |
| CALL_RETRIEVED_EVENT | 12 | Call retrieved from hold |
| CALL_CLEARED_EVENT | 13 | Call terminated |
| CALL_CONNECTION_CLEARED_EVENT | 14 | Party left call |
| CALL_ORIGINATED_EVENT | 15 | Outbound call initiated |
| CALL_FAILED_EVENT | 16 | Call failed |
| CALL_CONFERENCED_EVENT | 17 | Conference created |
| CALL_TRANSFERRED_EVENT | 18 | Call transferred |
| CALL_DIVERTED_EVENT | 19 | Call redirected |
| CALL_SERVICE_INITIATED_EVENT | 20 | Service initiated |
| CALL_QUEUED_EVENT | 21 | Call queued |
| BEGIN_CALL_EVENT | 23 | New call started |
| END_CALL_EVENT | 24 | Call ended |
| CALL_DATA_UPDATE_EVENT | 25 | Call data changed |
| CALL_DEQUEUED_EVENT | 86 | Call removed from queue |

### Agent Events

| Message | Type ID | Description |
|---------|---------|-------------|
| AGENT_STATE_EVENT | 30 | Agent state changed |
| SYSTEM_EVENT | 31 | System status change |

## Key Message Formats

### OPEN_REQ (Type ID: 3)

**Fixed Part:**

| Field | Type | Description |
|-------|------|-------------|
| InvokeID | UINT | Client-assigned request ID |
| VersionNumber | UINT | Protocol version (24) |
| IdleTimeout | UINT | Idle timeout in seconds |
| PeripheralID | UINT | Peripheral ID (0 = any) |
| ServicesRequested | UINT | Service mask bitmap |
| CallMsgMask | UINT | Call event filter |
| AgentStateMask | UINT | Agent state filter |
| ConfigMsgMask | UINT | Config event filter |
| Reserved1-3 | UINT | Reserved (set to 0) |

**Floating Part:**

| Tag | Field | Description |
|-----|-------|-------------|
| 1 | ClientID | Client identifier string |
| 2 | ClientPassword | Client password (if required) |

### OPEN_CONF (Type ID: 4)

**Fixed Part:**

| Field | Type | Description |
|-------|------|-------------|
| InvokeID | UINT | Matches OPEN_REQ |
| ServiceGranted | UINT | Granted services |
| MonitorID | UINT | Assigned monitor ID |
| PGStatus | UINT | Peripheral gateway status |
| ICMCentralController | UINT | ICM controller status |
| AgentState | USHORT | Current agent state |
| NumPeripherals | USHORT | Connected peripherals |
| MultiLineAgentControl | USHORT | Multi-line flag |
| Reserved | USHORT | Reserved |
| PeripheralID | UINT | Peripheral ID |

### HEARTBEAT_REQ (Type ID: 5)

**Fixed Part:**

| Field | Type | Description |
|-------|------|-------------|
| InvokeID | UINT | Request ID |

### HEARTBEAT_CONF (Type ID: 6)

**Fixed Part:**

| Field | Type | Description |
|-------|------|-------------|
| InvokeID | UINT | Matches request |

### AGENT_STATE_EVENT (Type ID: 30)

**Fixed Part:**

| Field | Type | Description |
|-------|------|-------------|
| MonitorID | UINT | Monitor ID |
| PeripheralID | UINT | Peripheral ID |
| SessionID | UINT | Session ID |
| PeripheralType | USHORT | Peripheral type |
| SkillGroupState | USHORT | Skill group state |
| StateDuration | UINT | Duration in state (seconds) |
| SkillGroupNumber | UINT | Skill group number |
| SkillGroupID | UINT | Skill group ID |
| SkillGroupPriority | USHORT | Skill group priority |
| AgentState | USHORT | Current agent state |
| EventReasonCode | USHORT | Reason for state change |
| MRDID | USHORT | Media routing domain ID |
| NumTasks | UINT | Active tasks |
| AgentMode | USHORT | Agent mode |
| MaxTaskLimit | USHORT | Max tasks allowed |
| ICMAgentID | UINT | ICM agent ID |
| AgentAvailabilityStatus | UINT | Availability status |
| NumFltSkillGroups | USHORT | Floating skill groups |
| Reserved | USHORT | Reserved |

**Floating Part:**

| Tag | Field | Description |
|-----|-------|-------------|
| 3 | AgentExtension | Agent's extension |
| 4 | AgentID | Agent's ID |
| 5 | AgentInstrument | Agent's instrument |

### BEGIN_CALL_EVENT (Type ID: 23)

**Fixed Part:**

| Field | Type | Description |
|-------|------|-------------|
| MonitorID | UINT | Monitor ID |
| PeripheralID | UINT | Peripheral ID |
| PeripheralType | USHORT | Peripheral type |
| ConnectionDeviceType | USHORT | Device type |
| ConnectionCallID | UINT | Call ID |
| ConnectionState | USHORT | Connection state |
| Reserved | USHORT | Reserved |
| ServiceNumber | UINT | Service number |
| ServiceID | UINT | Service ID |
| SkillGroupNumber | UINT | Skill group number |
| SkillGroupID | UINT | Skill group ID |
| SkillGroupPriority | USHORT | Skill group priority |
| CallType | USHORT | Type of call |
| CallingDeviceType | USHORT | Calling device type |
| CalledDeviceType | USHORT | Called device type |
| LastRedirectDeviceType | USHORT | Last redirect type |
| Reserved2 | USHORT | Reserved |

**Floating Part:**

| Tag | Field | Description |
|-----|-------|-------------|
| 15 | ANI | Automatic Number ID |
| 16 | DNIS | Dialed Number ID |
| 12 | CallingDeviceID | Calling device |
| 13 | CalledDeviceID | Called device |
| 14 | LastRedirectDeviceID | Last redirect device |
| 17 | UserToUserInfo | UUI data |
| 18-27 | CallVariable1-10 | Call variables |

## Field Tags

| Tag | Name | Type |
|-----|------|------|
| 1 | ClientID | STRING |
| 2 | ClientPassword | STRING |
| 3 | AgentExtension | STRING |
| 4 | AgentID | STRING |
| 5 | AgentInstrument | STRING |
| 6 | PeripheralID | UINT |
| 7 | ServiceNumber | UINT |
| 8 | ServiceID | UINT |
| 9 | SkillGroupNumber | UINT |
| 10 | SkillGroupID | UINT |
| 11 | SkillGroupPriority | USHORT |
| 12 | CallingDeviceID | STRING |
| 13 | CalledDeviceID | STRING |
| 14 | LastRedirectDeviceID | STRING |
| 15 | ANI | STRING |
| 16 | DNIS | STRING |
| 17 | UserToUserInfo | STRING |
| 18-27 | CallVariable1-10 | STRING |
| 28 | CTIClientSignature | STRING |
| 29 | CTIClientTimestamp | UINT |
| 30 | CallWrapupData | STRING |
| 82 | NamedVariable | NAMEDVAR |
| 83 | NamedArray | NAMEDARRAY |

## Agent States

| Value | State | Description |
|-------|-------|-------------|
| 0 | LoggedOut | Not logged in |
| 1 | LoggedIn | Logged in (transitional) |
| 2 | NotReady | Not ready for calls |
| 3 | Ready | Ready for calls |
| 4 | Talking | On a call |
| 5 | WorkNotReady | Wrap-up, not ready |
| 6 | WorkReady | Wrap-up, will be ready |
| 7 | Hold | Call on hold |
| 8 | Reserved | Reserved for call |
| 9 | Unknown | Unknown state |

## Service Masks

| Value | Service | Description |
|-------|---------|-------------|
| 0x00000001 | CLIENT_EVENTS | Events for client's device |
| 0x00000002 | CALL_CONTROL | Call control capabilities |
| 0x00000010 | ALL_EVENTS | All events (bridge mode) |
| 0x00000020 | SUPERVISOR | Supervisor capabilities |

## Session Flow

```
Client                                Server
   │                                    │
   │──── TCP Connect ─────────────────▶│
   │                                    │
   │──── OPEN_REQ ────────────────────▶│
   │     (InvokeID, Services, ClientID) │
   │                                    │
   │◀─── OPEN_CONF ────────────────────│
   │     (MonitorID, ServiceGranted)    │
   │                                    │
   │◀─── Events ───────────────────────│
   │     (AgentState, CallEvents, etc.) │
   │                                    │
   │──── HEARTBEAT_REQ ───────────────▶│ (every 30s)
   │◀─── HEARTBEAT_CONF ───────────────│
   │                                    │
   │──── CLOSE_REQ ───────────────────▶│
   │◀─── CLOSE_CONF ───────────────────│
   │                                    │
   │──── TCP Close ───────────────────▶│
```

## Error Handling

### FAILURE_CONF
Sent in response to a failed request. Contains InvokeID matching the request.

### FAILURE_EVENT
Unsolicited error notification. May indicate session will be terminated.

### Heartbeat Failure
If 3 consecutive HEARTBEAT_REQ messages receive no HEARTBEAT_CONF, the client should:
1. Close the TCP connection
2. Wait before reconnecting
3. Attempt to re-establish the session

## References

- [Cisco DevNet CTI Protocol](https://developer.cisco.com/site/cti-protocol/)
- [CTI Server Message Reference Guide](https://www.cisco.com/c/en/us/td/docs/voice_ip_comm/cust_contact/contact_center/icm_enterprise/icm_enterprise_12_6_1/reference/guide/ucce_b_cti-servermessage-reference-guide-for-1261.html)
