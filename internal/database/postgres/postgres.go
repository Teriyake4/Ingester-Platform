package postgres

import (
	"context"
	"log/slog"
	"ticker-ingester/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPGConn(ctx context.Context, cfg config.PostgresConfig) *pgxpool.Pool {
	pg, err := pgxpool.New(ctx, cfg.URL)
	if err != nil {
		slog.Error("error creating Postgres connection pool", "error", err)
		return nil
	}
	if err := pg.Ping(ctx); err != nil {
		slog.Error("error connecting to Postgres", "error", err)
		return nil
	}

	return pg
}
