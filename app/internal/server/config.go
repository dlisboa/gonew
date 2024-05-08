package server

import "log/slog"

type Config struct {
	Host     string     `env:"HOST" envDefault:""`
	Port     string     `env:"PORT" envDefault:"3000"`
	LogLevel slog.Level `env:"LOG_LEVEL" envDefault:"info"`
	DSN      string     `env:"DSN" envDefault:""`
}

type config = Config
