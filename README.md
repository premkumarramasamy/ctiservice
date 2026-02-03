# CTI Service

A Go microservice for receiving Cisco CTI server events using the GED-188 protocol.

## Overview

This service connects to a Cisco Unified Contact Center Enterprise (UCCE) CTI Server and receives real-time call and agent state events. It implements the GED-188 protocol (Protocol Version 24) for TCP socket communication.

## Features

- **GED-188 Protocol Implementation**: Full binary protocol encoding/decoding
- **Session Management**: OPEN/CLOSE handshake with automatic reconnection
- **Heartbeat Monitoring**: Configurable keepalive with failure detection
- **Event Processing**: Call events, agent state events, system events
- **Structured Logging**: JSON output for easy parsing and monitoring
- **Graceful Shutdown**: Clean session termination on SIGINT/SIGTERM

## Quick Start

### Build

```bash
go build -o ctiservice ./cmd/ctiservice
```

### Configure

Set environment variables:

```bash
# Required
export CTI_SERVER_HOST=192.168.1.100
export CTI_SERVER_PORT=42027

# Optional (defaults shown)
export CTI_CLIENT_ID=CTIService
export CTI_PERIPHERAL_ID=0
export CTI_SERVICES_REQUESTED=0x11        # ServiceAllEvents | ServiceClientEvents
export CTI_HEARTBEAT_INTERVAL=30s
export CTI_IDLE_TIMEOUT=120s
export CTI_RECONNECT_DELAY=10s
export CTI_LOG_LEVEL=info                 # debug, info, warn, error
```

### Run

```bash
./ctiservice
```

## Configuration Reference

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `CTI_SERVER_HOST` | localhost | CTI server hostname or IP |
| `CTI_SERVER_PORT` | 42027 | CTI server port |
| `CTI_CLIENT_ID` | CTIService | Client identifier sent in OPEN_REQ |
| `CTI_PERIPHERAL_ID` | 0 | Peripheral ID (0 = any) |
| `CTI_SERVICES_REQUESTED` | 0x11 | Service mask bitmap |
| `CTI_HEARTBEAT_INTERVAL` | 30s | Heartbeat send interval |
| `CTI_IDLE_TIMEOUT` | 120s | Server idle timeout (should be 4x heartbeat) |
| `CTI_RECONNECT_DELAY` | 10s | Wait time before reconnection attempt |
| `CTI_RECONNECT_MAX_ATTEMPTS` | 0 | Max reconnect attempts (0 = infinite) |
| `CTI_LOG_LEVEL` | info | Logging level |

## Service Mask Values

| Service | Value | Description |
|---------|-------|-------------|
| `ServiceClientEvents` | 0x01 | Agent mode - events for client's device |
| `ServiceCallControl` | 0x02 | Call control capabilities |
| `ServiceAllEvents` | 0x10 | Bridge mode - receive all events |
| `ServiceSupervisor` | 0x20 | Supervisor capabilities |

## Supported Events

### Call Events
- `BEGIN_CALL_EVENT` - New call started
- `END_CALL_EVENT` - Call ended
- `CALL_DELIVERED_EVENT` - Call arrived at device
- `CALL_ESTABLISHED_EVENT` - Call answered
- `CALL_HELD_EVENT` - Call placed on hold
- `CALL_RETRIEVED_EVENT` - Call retrieved from hold
- `CALL_CLEARED_EVENT` - Call terminated
- `CALL_ORIGINATED_EVENT` - Outbound call initiated
- `CALL_FAILED_EVENT` - Call failed
- `CALL_CONFERENCED_EVENT` - Conference created
- `CALL_TRANSFERRED_EVENT` - Call transferred
- `CALL_QUEUED_EVENT` - Call placed in queue
- `CALL_DEQUEUED_EVENT` - Call removed from queue
- `CALL_DATA_UPDATE_EVENT` - Call data changed

### Agent Events
- `AGENT_STATE_EVENT` - Agent state changed (Ready, NotReady, Talking, etc.)

### System Events
- `SYSTEM_EVENT` - System status changes (peripheral online/offline, etc.)

## Project Structure

```
ctiservice/
├── cmd/ctiservice/
│   └── main.go              # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   ├── protocol/            # GED-188 protocol implementation
│   ├── messages/            # Message type definitions
│   ├── client/              # CTI client and session management
│   └── handler/             # Event handlers
├── docs/
│   ├── ARCHITECTURE.md      # Architecture documentation
│   ├── PROTOCOL.md          # GED-188 protocol reference
│   └── DEVELOPMENT.md       # Development guide
└── README.md
```

## Documentation

- [Architecture](docs/ARCHITECTURE.md) - System design and components
- [Protocol Reference](docs/PROTOCOL.md) - GED-188 protocol details
- [Development Guide](docs/DEVELOPMENT.md) - Contributing and extending

## References

- [Cisco DevNet CTI Protocol Documentation](https://developer.cisco.com/site/cti-protocol/documentation/)
- [CTI Server Message Reference Guide (Protocol Version 24)](https://www.cisco.com/c/en/us/td/docs/voice_ip_comm/cust_contact/contact_center/icm_enterprise/icm_enterprise_12_6_1/reference/guide/ucce_b_cti-servermessage-reference-guide-for-1261.html)

## License

Internal use only.
