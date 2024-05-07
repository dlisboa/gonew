package server

import "log/slog"

type config struct {
	Host               string     `env:"HOST,required" envDefault:""`
	Port               string     `env:"PORT" envDefault:"3000"`
	Debug              bool       `env:"DEBUG" envDefault:"false"`
	LogLevel           slog.Level `env:"LOG_LEVEL" envDefault:"info"`
	LogFormat          string     `env:"LOG_FORMAT" envDefault:"text"`
	DatabaseSourceName string     `env:"DATABASE_SOURCE_NAME" envDefault:""`
	DatabaseURL        string     `env:"DATABASE_URL" envDefault:""`
}
