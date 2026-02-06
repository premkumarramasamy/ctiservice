// Package config handles configuration for the CTI service.
package config

import (
	"ctiservice/internal/protocol"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds the configuration for the CTI client.
type Config struct {
	// Server connection
	ServerHost string
	ServerPort int

	// Session settings
	ClientID          string
	PeripheralID      uint32
	ServicesRequested uint32
	IdleTimeout       time.Duration

	// Event subscription masks for OPEN_REQ
	CallMsgMask       uint32 // Bitmask for call events to receive
	AgentStateMask    uint32 // Bitmask for agent state events to receive
	ConfigMsgMask     uint32 // Bitmask for config events to receive

	// Heartbeat settings
	HeartbeatInterval time.Duration

	// Reconnection settings
	ReconnectDelay       time.Duration
	ReconnectMaxAttempts int // 0 = infinite

	// Logging
	LogLevel string
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		ServerHost:           "localhost",
		ServerPort:           42027,
		ClientID:             "CTIService",
		PeripheralID:         0,
		ServicesRequested:    protocol.ServiceAllEvents | protocol.ServiceClientEvents,
		IdleTimeout:          120 * time.Second,
		CallMsgMask:          protocol.CallMaskAll,       // Subscribe to all call events
		AgentStateMask:       protocol.AgentMaskAll,      // Subscribe to all agent events
		ConfigMsgMask:        protocol.ConfigMaskAll,     // Subscribe to all config events
		HeartbeatInterval:    30 * time.Second,
		ReconnectDelay:       10 * time.Second,
		ReconnectMaxAttempts: 0,
		LogLevel:             "info",
	}
}

// LoadFromEnv loads configuration from environment variables.
func LoadFromEnv() (*Config, error) {
	cfg := DefaultConfig()

	if v := os.Getenv("CTI_SERVER_HOST"); v != "" {
		cfg.ServerHost = v
	}

	if v := os.Getenv("CTI_SERVER_PORT"); v != "" {
		port, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_SERVER_PORT: %w", err)
		}
		cfg.ServerPort = port
	}

	if v := os.Getenv("CTI_CLIENT_ID"); v != "" {
		cfg.ClientID = v
	}

	if v := os.Getenv("CTI_PERIPHERAL_ID"); v != "" {
		id, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_PERIPHERAL_ID: %w", err)
		}
		cfg.PeripheralID = uint32(id)
	}

	if v := os.Getenv("CTI_SERVICES_REQUESTED"); v != "" {
		services, err := strconv.ParseUint(v, 0, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_SERVICES_REQUESTED: %w", err)
		}
		cfg.ServicesRequested = uint32(services)
	}

	if v := os.Getenv("CTI_CALL_MSG_MASK"); v != "" {
		mask, err := strconv.ParseUint(v, 0, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_CALL_MSG_MASK: %w", err)
		}
		cfg.CallMsgMask = uint32(mask)
	}

	if v := os.Getenv("CTI_AGENT_STATE_MASK"); v != "" {
		mask, err := strconv.ParseUint(v, 0, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_AGENT_STATE_MASK: %w", err)
		}
		cfg.AgentStateMask = uint32(mask)
	}

	if v := os.Getenv("CTI_CONFIG_MSG_MASK"); v != "" {
		mask, err := strconv.ParseUint(v, 0, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_CONFIG_MSG_MASK: %w", err)
		}
		cfg.ConfigMsgMask = uint32(mask)
	}

	if v := os.Getenv("CTI_IDLE_TIMEOUT"); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_IDLE_TIMEOUT: %w", err)
		}
		cfg.IdleTimeout = d
	}

	if v := os.Getenv("CTI_HEARTBEAT_INTERVAL"); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_HEARTBEAT_INTERVAL: %w", err)
		}
		cfg.HeartbeatInterval = d
	}

	if v := os.Getenv("CTI_RECONNECT_DELAY"); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_RECONNECT_DELAY: %w", err)
		}
		cfg.ReconnectDelay = d
	}

	if v := os.Getenv("CTI_RECONNECT_MAX_ATTEMPTS"); v != "" {
		attempts, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid CTI_RECONNECT_MAX_ATTEMPTS: %w", err)
		}
		cfg.ReconnectMaxAttempts = attempts
	}

	if v := os.Getenv("CTI_LOG_LEVEL"); v != "" {
		cfg.LogLevel = v
	}

	return cfg, nil
}

// Validate checks if the configuration is valid.
func (c *Config) Validate() error {
	if c.ServerHost == "" {
		return fmt.Errorf("server host is required")
	}
	if c.ServerPort <= 0 || c.ServerPort > 65535 {
		return fmt.Errorf("invalid server port: %d", c.ServerPort)
	}
	if c.HeartbeatInterval < time.Second {
		return fmt.Errorf("heartbeat interval too short: %v", c.HeartbeatInterval)
	}
	if c.IdleTimeout < c.HeartbeatInterval*4 {
		return fmt.Errorf("idle timeout should be at least 4x heartbeat interval")
	}
	return nil
}

// String returns a string representation of the config (for logging).
func (c *Config) String() string {
	return fmt.Sprintf(
		"Config{ServerHost=%s, ServerPort=%d, ClientID=%s, HeartbeatInterval=%v, IdleTimeout=%v}",
		c.ServerHost, c.ServerPort, c.ClientID, c.HeartbeatInterval, c.IdleTimeout,
	)
}
