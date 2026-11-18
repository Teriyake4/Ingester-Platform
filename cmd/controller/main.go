package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"ticker-ingester/internal/config"
	"ticker-ingester/internal/database/postgres"
	"ticker-ingester/internal/database/redis"
	"ticker-ingester/internal/observability"
	"ticker-ingester/internal/server"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("error loading config", "error", err)
		os.Exit(1)
	}

	logger, err := observability.NewLogger(cfg.Server)
	if err != nil {
		slog.Error("creating logger", "error", err)
		os.Exit(1)
	}
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	slog.Info("connecting to Postgres")
	pg := postgres.NewPGConn(ctx, cfg.Postgres)
	slog.Info("successfully connected to Postgres")

	slog.Info("connecting to Redis")
	rdb := redis.NewRedisConn(ctx, cfg.Redis)
	slog.Info("successfully connected to Redis")

	slog.Info("creating server")

	router := server.NewRoutes(pg, rdb)
	handler := server.Chain(
		server.RecoverPanic(),
		server.RequestID(),
		server.Logging(),
	)(router)
	srv := server.NewServer(cfg.Server, logger, handler)

	if err := server.Run(ctx, srv); nil != err {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}
