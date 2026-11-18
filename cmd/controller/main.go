package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"ticker-ingester/internal/server"
	"ticker-ingester/internal/config"
	"ticker-ingester/internal/observability"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("load config", "error", err)
		os.Exit(1)
	}

	logger, err := util.NewLogger(cfg.Server)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger.Info("Creating server")

	router := server.NewRoutes()
	handler := server.Chain(
		server.RecoverPanic(),
		server.RequestID(),
		server.Logging(),
	)(router)
	srv := server.NewServer(cfg.Server, logger, handler)

	if err := server.Run(ctx, srv); nil != err {
		logger.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
