package server

import "log/slog"

type config struct {
	Host      string     `env:"HOST,required" envDefault:""`
	Port      string     `env:"PORT" envDefault:"3000"`
	Debug     bool       `env:"DEBUG" envDefault:"false"`
	LogLevel  slog.Level `env:"LOG_LEVEL" envDefault:"info"`
	LogFormat string     `env:"LOG_FORMAT" envDefault:"text"`
	DBURL     string     `env:"DB_URL" envDefault:""`
}
