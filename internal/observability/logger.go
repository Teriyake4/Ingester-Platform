package observability

import (
	"fmt"
	"log/slog"
	"os"
	"ticker-ingester/internal/config"
)

func NewLogger(cfg config.ServerConfig) (*slog.Logger, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(cfg.LogLevel)); err != nil {
		return nil, fmt.Errorf("invalid log level %q: %w", cfg.LogLevel, err)
	}
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	return slog.New(handler), nil
}
