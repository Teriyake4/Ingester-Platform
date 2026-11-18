package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Server ServerConfig
	Worker WorkerConfig
	Postgres PostgresConfig
	Redis RedisConfig
	ClickHouse ClickHouseConfig
}

type ServerConfig struct {
	LogLevel string
	Host string
	Port string
	ReadTimeout int
	WriteTimeout int
}

type WorkerConfig struct {
	Concurrency int
}

type PostgresConfig struct {
	URL string
}

type RedisConfig struct {
	URL string
}

type ClickHouseConfig struct {
	URL string
}


func LoadConfig() (Config, error) {
	logLevel := getEnv("LOG_LEVEL", "info")
	switch strings.ToLower(logLevel) {
		case "debug", "info", "warn", "error":
		default:
			return Config{}, fmt.Errorf("invalid LOG_LEVEL: %q", logLevel)
	}

	readTimeoutStr := getEnv("READ_TIMEOUT", "10")
	readTimeout, err := strconv.Atoi(readTimeoutStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid READ_TIMEOUT: %q", readTimeoutStr)
	}

	writeTimeoutStr := getEnv("WRITE_TIMEOUT", "10")
	writeTimeout, err := strconv.Atoi(writeTimeoutStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid WRITE_TIMEOUT: %q", writeTimeoutStr)
	}


	postgresURL, err := getRequiredEnv("POSTGRES_URL")
	if err != nil {
		return Config{}, err
	}
	redisURL, err := getRequiredEnv("REDIS_URL")
	if err != nil {
		return Config{}, err
	}
	clickHouseURL, err := getRequiredEnv("CLICKHOUSE_URL")
	if err != nil {
		return Config{}, err
	}

	return Config {
		Server: ServerConfig{
			LogLevel: logLevel,
			Host: getEnv("HOST", ""),
			Port: getEnv("PORT", "1116"),
			ReadTimeout: readTimeout,
			WriteTimeout: writeTimeout,
		},
		Worker: WorkerConfig{Concurrency: 3},
		Postgres: PostgresConfig{URL: postgresURL},
		Redis: RedisConfig{URL: redisURL},
		ClickHouse: ClickHouseConfig{URL: clickHouseURL},
	}, nil
}

func getRequiredEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return "", fmt.Errorf("required environment variable %q is not set", key)
	}
	return value, nil
}

func getEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}
	return value
}
