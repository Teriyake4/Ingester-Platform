package server

import (
	"net/http"
	"ticker-ingester/health"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func NewRoutes(pg *pgxpool.Pool, rdb *redis.Client) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", health.Health)

	// mux.HandleFunc("GET /api/v1/jobs", )

	return mux
}
