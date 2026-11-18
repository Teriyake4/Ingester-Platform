package server

import (
	"net/http"
	"ticker-ingester/health"
)

func NewRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", health.Health)

	// mux.HandleFunc("GET /api/v1/jobs", )

	return mux
}
