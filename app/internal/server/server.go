package server

import (
	"context"
	"database/sql"
	"io"
	"log/slog"
	"net"
	"net/http"

	"github.com/caarlos0/env/v11"
	_ "github.com/mattn/go-sqlite3"
)

func Run(ctx context.Context, args []string, out io.Writer, getenv func(string) string) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	logger := slog.New(logHandler(out, cfg))

	db, err := sql.Open("sqlite3", cfg.DatabaseSourceName)
	if err != nil {
		return err
	}
	defer db.Close()

	handlers := &Handlers{
		config: cfg,
		logger: logger,
		db:     db,
	}

	srv := http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: handlers.NewRouter(),
	}

	logger.Info("listening on " + srv.Addr)
	return srv.ListenAndServe()
}

func logHandler(out io.Writer, cfg config) slog.Handler {
	opts := &slog.HandlerOptions{Level: cfg.LogLevel}
	if cfg.LogFormat == "json" {
		return slog.NewJSONHandler(out, opts)
	}
	return slog.NewTextHandler(out, opts)
}
