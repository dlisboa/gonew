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
	_ "github.com/mattn/go-sqlite3"

	"github.com/dlisboa/gonew/app/internal/database"
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
		return fmt.Errorf("cannot parse config: %w", err)
	}

	logger := newLogger(out, cfg)

	db, err := setupDB(cfg)
	if err != nil {
		return fmt.Errorf("cannot setup db: %w", err)
	}
	defer db.Close()

	queries := database.New(db)
	handler := server.NewServer(cfg, logger, queries)

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

func setupDB(cfg server.Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.DSN)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
