package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/dlisboa/gonew/app/internal/server"
)

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatalf("main: %s. Exiting.", err)
	}
}

func run(out io.Writer) error {
	// ctx := context.Background()

	cfg := server.Config{}
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	logger := newLogger(out, cfg)

	db, err := sql.Open("sqlite3", cfg.DSN)
	if err != nil {
		return fmt.Errorf("sql open: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("db ping: %w", err)
	}
	defer db.Close()

	handler := server.NewServer(cfg, logger, db)

	srv := http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: handler,
	}

	logger.Info("listening on " + srv.Addr)
	return srv.ListenAndServe()
}

func newLogger(w io.Writer, cfg server.Config) *slog.Logger {
	opts := &slog.HandlerOptions{Level: cfg.LogLevel}
	return slog.New(slog.NewTextHandler(w, opts))
}
