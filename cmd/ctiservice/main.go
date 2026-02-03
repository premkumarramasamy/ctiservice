// Package main is the entry point for the CTI service.
package main

import (
	"context"
	"ctiservice/internal/client"
	"ctiservice/internal/config"
	"ctiservice/internal/handler"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load configuration
	cfg, err := config.LoadFromEnv()
	if err != nil {
		slog.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		slog.Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	// Setup logger
	logLevel := parseLogLevel(cfg.LogLevel)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

	logger.Info("starting CTI service", "config", cfg.String())

	// Create event handler
	eventHandler := handler.NewLogHandler(logger.With("component", "events"))

	// Create context with cancellation on SIGINT/SIGTERM
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Create and run the client
	ctiClient := client.New(cfg, logger.With("component", "client"), eventHandler.Handle)

	logger.Info("connecting to CTI server",
		"host", cfg.ServerHost,
		"port", cfg.ServerPort)

	if err := ctiClient.Run(ctx); err != nil {
		if ctx.Err() != nil {
			logger.Info("shutting down gracefully")
		} else {
			logger.Error("client error", "error", err)
			os.Exit(1)
		}
	}

	logger.Info("CTI service stopped")
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
