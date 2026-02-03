# Development Guide

## Prerequisites

- Go 1.21 or later (for `log/slog`)
- Access to a Cisco UCCE CTI Server (for integration testing)

## Project Setup

```bash
# Clone/navigate to project
cd ctiservice

# Download dependencies
go mod download

# Build
go build ./...

# Run tests
go test ./...

# Run with race detector
go run -race ./cmd/ctiservice
```

## Code Organization

```
ctiservice/
├── cmd/ctiservice/          # Application entry point
├── internal/                # Private packages
│   ├── config/              # Configuration
│   ├── protocol/            # Binary protocol layer
│   ├── messages/            # Message definitions
│   ├── client/              # Connection management
│   └── handler/             # Event processing
├── docs/                    # Documentation
└── go.mod
```

## Adding New Message Types

### Step 1: Add Constant

In `internal/protocol/constants.go`:

```go
const (
    // ... existing constants
    MsgTypeNewMessage uint32 = 123
)
```

Update `MessageTypeName()` function.

### Step 2: Create Message Struct

In `internal/messages/`, create or add to appropriate file:

```go
type NewMessage struct {
    // Fixed fields (in order)
    InvokeID uint32
    SomeField uint16
    // ...

    // Floating fields
    SomeString string
}

func (m *NewMessage) Type() uint32 {
    return protocol.MsgTypeNewMessage
}

func (m *NewMessage) Encode() ([]byte, error) {
    w := protocol.NewFixedFieldWriter()
    w.WriteUint32(m.InvokeID)
    w.WriteUint16(m.SomeField)

    if err := w.Error(); err != nil {
        return nil, err
    }

    // Add floating fields if any
    fw := protocol.NewFloatingFieldWriter()
    if m.SomeString != "" {
        fw.WriteString(protocol.TagSomeField, m.SomeString)
    }

    fixed := w.Bytes()
    floating := fw.Bytes()
    result := make([]byte, len(fixed)+len(floating))
    copy(result, fixed)
    copy(result[len(fixed):], floating)

    return result, nil
}

func (m *NewMessage) Decode(data []byte) error {
    r := protocol.NewFixedFieldReader(data)

    m.InvokeID = r.ReadUint32()
    m.SomeField = r.ReadUint16()

    if err := r.Error(); err != nil {
        return err
    }

    // Parse floating fields
    if r.Remaining() > 0 {
        ff, err := protocol.ParseFloatingFields(r.RemainingBytes())
        if err != nil {
            return err
        }
        m.SomeString = ff.GetString(protocol.TagSomeField)
    }

    return nil
}
```

### Step 3: Register in Registry

In `internal/messages/registry.go`, add to `Create()`:

```go
case protocol.MsgTypeNewMessage:
    return &NewMessage{}
```

### Step 4: Handle in Client (if needed)

For request/response messages, add handling in `internal/client/client.go`.

### Step 5: Add to Logger

In `internal/handler/logger.go`, add case to `Handle()`:

```go
case *messages.NewMessage:
    h.logger.Info("new message received",
        append(attrs,
            "invokeID", m.InvokeID,
            "someField", m.SomeField,
        )...)
```

## Adding New Event Handlers

### Implement EventHandler Interface

```go
package myhandler

import (
    "ctiservice/internal/protocol"
    "ctiservice/internal/messages"
)

type MyHandler struct {
    // dependencies
}

func New( /* deps */ ) *MyHandler {
    return &MyHandler{ /* ... */ }
}

func (h *MyHandler) Handle(msg protocol.Message) {
    switch m := msg.(type) {
    case *messages.AgentStateEvent:
        // Process agent state
    case *messages.BeginCallEvent:
        // Process call start
    }
}
```

### Wire Into Main

```go
// In main.go
myHandler := myhandler.New()
logHandler := handler.NewLogHandler(logger)

// Use MultiHandler to chain
multiHandler := handler.NewMultiHandler(logHandler, myHandler)

client := client.New(cfg, logger, multiHandler.Handle)
```

## Adding Message Queue Output

Example Kafka handler:

```go
package kafka

import (
    "encoding/json"
    "ctiservice/internal/protocol"
    "ctiservice/internal/messages"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaHandler struct {
    producer *kafka.Producer
    topic    string
}

func New(brokers, topic string) (*KafkaHandler, error) {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": brokers,
    })
    if err != nil {
        return nil, err
    }
    return &KafkaHandler{producer: p, topic: topic}, nil
}

func (h *KafkaHandler) Handle(msg protocol.Message) {
    event := struct {
        Type string      `json:"type"`
        Data interface{} `json:"data"`
    }{
        Type: protocol.MessageTypeName(msg.Type()),
        Data: msg,
    }

    data, _ := json.Marshal(event)

    h.producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{
            Topic:     &h.topic,
            Partition: kafka.PartitionAny,
        },
        Value: data,
    }, nil)
}

func (h *KafkaHandler) Close() {
    h.producer.Close()
}
```

## Testing

### Unit Testing Protocol Encoding

```go
func TestOpenReqEncode(t *testing.T) {
    msg := &messages.OpenReq{
        InvokeID:          1,
        VersionNumber:     24,
        IdleTimeout:       120,
        ServicesRequested: 0x11,
        ClientID:          "test",
    }

    data, err := msg.Encode()
    if err != nil {
        t.Fatal(err)
    }

    // Verify fixed part
    buf := protocol.NewBuffer(data)
    invokeID, _ := buf.ReadUint32()
    if invokeID != 1 {
        t.Errorf("expected invokeID 1, got %d", invokeID)
    }
    // ... more assertions
}
```

### Integration Testing with Mock Server

```go
func TestClientConnect(t *testing.T) {
    // Start mock CTI server
    ln, _ := net.Listen("tcp", "127.0.0.1:0")
    defer ln.Close()

    go func() {
        conn, _ := ln.Accept()
        defer conn.Close()

        // Read OPEN_REQ
        reader := client.NewReader(conn)
        msg, _ := reader.ReadMessage()

        openReq := msg.(*messages.OpenReq)

        // Send OPEN_CONF
        conf := &messages.OpenConf{
            InvokeID:       openReq.InvokeID,
            ServiceGranted: openReq.ServicesRequested,
            MonitorID:      12345,
        }
        data, _ := protocol.EncodeMessage(conf)
        conn.Write(data)

        // Keep connection open
        time.Sleep(time.Second)
    }()

    // Create client
    cfg := config.DefaultConfig()
    cfg.ServerHost = "127.0.0.1"
    _, port, _ := net.SplitHostPort(ln.Addr().String())
    cfg.ServerPort, _ = strconv.Atoi(port)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    c := client.New(cfg, slog.Default(), nil)

    // Run briefly
    go c.Run(ctx)
    time.Sleep(500 * time.Millisecond)

    if c.State() != client.StateOpen {
        t.Errorf("expected StateOpen, got %v", c.State())
    }
}
```

## Debugging

### Enable Debug Logging

```bash
export CTI_LOG_LEVEL=debug
```

### Capture Raw Messages

Modify `reader.go` to log raw bytes:

```go
func (r *Reader) ReadMessage() (protocol.Message, error) {
    // ... read header and body

    // Debug: log raw message
    log.Printf("RAW: type=%d len=%d data=%x",
        header.MessageType, header.MessageLength, body)

    // ... parse and return
}
```

### Network Capture

```bash
# Capture CTI traffic
tcpdump -i any -w cti.pcap port 42027

# View in Wireshark
wireshark cti.pcap
```

## Performance Considerations

### Buffer Pooling

For high-volume scenarios, consider pooling message buffers:

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 4096)
    },
}

func (r *Reader) ReadMessage() (protocol.Message, error) {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf)
    // ... use buf
}
```

### Batch Processing

For handlers that write to databases/queues:

```go
type BatchHandler struct {
    batch   []protocol.Message
    mu      sync.Mutex
    ticker  *time.Ticker
    flushFn func([]protocol.Message)
}

func (h *BatchHandler) Handle(msg protocol.Message) {
    h.mu.Lock()
    h.batch = append(h.batch, msg)
    if len(h.batch) >= 100 {
        h.flush()
    }
    h.mu.Unlock()
}

func (h *BatchHandler) flush() {
    if len(h.batch) == 0 {
        return
    }
    h.flushFn(h.batch)
    h.batch = h.batch[:0]
}
```

## Common Issues

### Connection Refused
- Verify CTI Server is running
- Check firewall rules for port 42027/43027
- Verify CTI_SERVER_HOST and CTI_SERVER_PORT

### OPEN_REQ Rejected
- Check client credentials
- Verify ServicesRequested mask is valid
- Check peripheral ID

### Heartbeat Failures
- Network latency may require longer interval
- Increase CTI_HEARTBEAT_INTERVAL
- Ensure CTI_IDLE_TIMEOUT > 4 * HeartbeatInterval

### Parse Errors
- Protocol version mismatch
- Check CTI Server version compatibility
- Enable debug logging to see raw bytes

## Future Enhancements

Potential improvements for this codebase:

1. **Metrics**: Add Prometheus metrics for events/sec, latency, errors
2. **Tracing**: OpenTelemetry integration for distributed tracing
3. **TLS**: Support encrypted connections if CTI Server supports it
4. **Multiple Servers**: Connection pooling for multiple CTI Servers
5. **Failover**: Active-passive failover between primary/secondary servers
6. **Call Control**: Implement request messages (MAKE_CALL_REQ, etc.)
7. **Configuration**: Support config file (YAML/JSON) in addition to env vars
8. **Health Check**: HTTP endpoint for Kubernetes readiness/liveness probes
