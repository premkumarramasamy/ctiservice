# Architecture

## Overview

The CTI Service is a Go microservice that connects to Cisco UCCE CTI Servers using the GED-188 protocol over TCP. It receives real-time call and agent events and processes them through configurable handlers.

## System Context

```
┌─────────────────┐         TCP/42027          ┌─────────────────┐
│                 │  ─────────────────────────▶│                 │
│   CTI Service   │      GED-188 Protocol      │   CTI Server    │
│    (This App)   │  ◀─────────────────────────│     (Cisco)     │
│                 │         Events             │                 │
└─────────────────┘                            └─────────────────┘
        │
        │ Structured JSON Logs
        ▼
   ┌─────────┐
   │  stdout │
   └─────────┘
```

## Component Architecture

```
┌──────────────────────────────────────────────────────────────────┐
│                          main.go                                  │
│  - Load configuration                                            │
│  - Setup logger                                                  │
│  - Create client                                                 │
│  - Handle signals                                                │
└─────────────────────────────┬────────────────────────────────────┘
                              │
                              ▼
┌──────────────────────────────────────────────────────────────────┐
│                     internal/client                               │
│  ┌────────────┐  ┌────────────┐  ┌────────────┐  ┌────────────┐ │
│  │  client.go │  │ session.go │  │ heartbeat  │  │  reader.go │ │
│  │            │  │            │  │    .go     │  │            │ │
│  │ - Connect  │  │ - State    │  │            │  │ - Read     │ │
│  │ - Open     │  │   machine  │  │ - Periodic │  │   header   │ │
│  │ - Run loop │  │ - Monitor  │  │   HB send  │  │ - Read     │ │
│  │ - Close    │  │   ID       │  │ - Failure  │  │   body     │ │
│  │ - Reconnect│  │ - Invoke   │  │   detect   │  │ - Parse    │ │
│  └────────────┘  │   ID       │  └────────────┘  └────────────┘ │
│                  └────────────┘                                  │
└─────────────────────────────┬────────────────────────────────────┘
                              │
          ┌───────────────────┼───────────────────┐
          ▼                   ▼                   ▼
┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐
│ internal/protocol│  │ internal/messages│  │ internal/handler│
│                 │  │                 │  │                 │
│ - constants.go  │  │ - session.go    │  │ - handler.go    │
│ - types.go      │  │ - events.go     │  │ - logger.go     │
│ - header.go     │  │ - call_events.go│  │                 │
│ - message.go    │  │ - agent_events  │  │ EventHandler    │
│ - fixed.go      │  │ - registry.go   │  │ interface       │
│ - floating.go   │  │                 │  │                 │
└─────────────────┘  └─────────────────┘  └─────────────────┘
```

## Package Responsibilities

### `cmd/ctiservice`
Entry point. Loads configuration, initializes components, runs the client, handles OS signals.

### `internal/config`
Configuration management. Loads settings from environment variables with sensible defaults.

### `internal/protocol`
Low-level GED-188 protocol implementation:
- **constants.go**: Message type IDs, field tags, status codes, service masks
- **types.go**: Protocol data types (ConnectionID, call types, states)
- **header.go**: 8-byte message header encoding/decoding
- **message.go**: Message interface, Buffer helpers for binary I/O
- **fixed.go**: Fixed field reader/writer with error accumulation
- **floating.go**: Tag-length-value floating field parser/writer

### `internal/messages`
Message type definitions implementing `protocol.Message`:
- **session.go**: OpenReq, OpenConf, HeartbeatReq, HeartbeatConf, CloseReq, CloseConf
- **events.go**: FailureConf, FailureEvent, SystemEvent
- **call_events.go**: All call-related event messages
- **agent_events.go**: AgentStateEvent
- **registry.go**: Message factory by type ID

### `internal/client`
CTI client connection management:
- **client.go**: Main orchestrator - connect, open session, process messages, reconnect
- **session.go**: Session state machine (Disconnected → Connecting → Connected → Opening → Open)
- **heartbeat.go**: Periodic heartbeat sender with 3-strike failure detection
- **reader.go**: TCP stream reader that parses complete messages

### `internal/handler`
Event handling:
- **handler.go**: EventHandler interface
- **logger.go**: JSON structured logging handler

## Data Flow

### Connection Establishment

```
1. Client.Run() called
   │
2. client.connect()
   │  - Dial TCP to CTI_SERVER_HOST:CTI_SERVER_PORT
   │  - Create Reader
   │  - State: Disconnected → Connecting → Connected
   │
3. client.open()
   │  - Send OPEN_REQ with InvokeID, VersionNumber, ServicesRequested
   │  - Wait for OPEN_CONF (30s timeout)
   │  - Store MonitorID, ServiceGranted
   │  - State: Connected → Opening → Open
   │
4. Start heartbeat goroutine
   │
5. client.processMessages() loop
   │  - Read message from TCP
   │  - Dispatch to handleMessage()
   │  - Call EventHandler for events
   │
6. On error or context cancel:
   - Send CLOSE_REQ
   - Close TCP connection
   - State: Open → Closing → Disconnected
   - If not canceled, wait ReconnectDelay and goto step 2
```

### Message Processing

```
TCP Stream
    │
    ▼
┌─────────────────┐
│ Reader.ReadMessage()
│  1. Read 8-byte header
│  2. Parse MessageLength, MessageType
│  3. Read MessageLength bytes (body)
│  4. Registry.Parse(type, body)
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│ Registry.Parse()
│  1. Create message struct by type ID
│  2. Call msg.Decode(body)
│  3. Return typed message
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│ client.handleMessage()
│  - HeartbeatConf → heartbeat.Confirm()
│  - CloseConf → update state
│  - FailureConf/Event → log error
│  - Others → call EventHandler
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│ LogHandler.Handle()
│  - Extract fields based on message type
│  - Log structured JSON
└─────────────────┘
```

## State Machine

```
                    ┌──────────────┐
                    │ Disconnected │◀─────────────────┐
                    └──────┬───────┘                  │
                           │ connect()                │
                           ▼                          │
                    ┌──────────────┐                  │
                    │  Connecting  │                  │
                    └──────┬───────┘                  │
                           │ TCP connected            │
                           ▼                          │
                    ┌──────────────┐                  │
                    │  Connected   │                  │
                    └──────┬───────┘                  │
                           │ send OPEN_REQ            │
                           ▼                          │
                    ┌──────────────┐                  │
                    │   Opening    │──────────────────┤
                    └──────┬───────┘  FAILURE_CONF    │
                           │ OPEN_CONF                │
                           ▼                          │
                    ┌──────────────┐                  │
          ┌────────│     Open     │──────────────────┤
          │        └──────────────┘  error/timeout   │
          │                                          │
          │ CLOSE_REQ                                │
          ▼                                          │
   ┌──────────────┐                                  │
   │   Closing    │──────────────────────────────────┘
   └──────────────┘  CLOSE_CONF / timeout
```

## Error Handling

### Reconnection Strategy
1. On any connection/session error, close the TCP connection
2. Wait `CTI_RECONNECT_DELAY` (default 10s)
3. Attempt to reconnect
4. Repeat indefinitely (or until `CTI_RECONNECT_MAX_ATTEMPTS`)

### Heartbeat Failure
1. Send HEARTBEAT_REQ every `CTI_HEARTBEAT_INTERVAL`
2. Track unconfirmed heartbeats
3. If 3 consecutive heartbeats go unconfirmed, trigger reconnect

### Message Parse Errors
- Unknown message types are captured in `GenericMessage` with raw bytes
- Parse errors are logged but don't terminate the session

## Thread Safety

- `Client.mu` protects connection and reader access
- `Session` uses RWMutex for state access
- `Heartbeat` uses Mutex for unconfirmed counter
- All state transitions are synchronized

## Extension Points

### Adding New Event Handlers
1. Implement `handler.EventHandler` interface
2. Use `handler.MultiHandler` to chain handlers
3. Pass to `client.New()`

### Adding New Message Types
1. Add constant to `protocol/constants.go`
2. Create struct in `messages/` implementing `protocol.Message`
3. Register in `messages/registry.go`

### Adding Message Queue Output
1. Create new handler implementing `EventHandler`
2. Connect to message queue in handler constructor
3. Serialize messages in `Handle()` method
