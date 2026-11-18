package redis

import (
	"context"
	"log/slog"
	"ticker-ingester/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisConn(ctx context.Context, cfg config.RedisConfig) *redis.Client {
	opt, err := redis.ParseURL(cfg.URL)
	if err != nil {
		slog.Error("error parsing redis URL", "error", err)
		return nil
	}
	rdb := redis.NewClient(opt)
	if err := rdb.Ping(ctx).Err(); err != nil {
		slog.Error("error connecting to Redis", "error", err)
		return nil
	}

	return rdb
}
