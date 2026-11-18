package server

import (
	"context"
	"log/slog"
	"net/http"
	"ticker-ingester/internal/config"
	"time"
)

func NewServer(cfg config.ServerConfig, logger *slog.Logger, handler http.Handler) *http.Server {
	return &http.Server{
		Addr: cfg.Host + ":" + cfg.Port,
		Handler: handler,
		ReadTimeout: time.Duration(cfg.ReadTimeout),
		WriteTimeout: time.Duration(cfg.WriteTimeout),
	}
}

func Run(ctx context.Context, srv *http.Server) error {
	return nil
}
