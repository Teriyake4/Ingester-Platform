package main

import (
	"log/slog"
	"os"
	"ticker-ingester/internal/config"
	"ticker-ingester/internal/util"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("load config", "error", err)
		os.Exit(1)
	}

	logger, err := util.NewLogger(cfg.Server)

	logger.Info("Starting Server")
}
