package server

import (
	"database/sql"
	"log/slog"
	"net/http"
)

type server struct {
	config Config
	logger *slog.Logger
	db     *sql.DB
	mux    *http.ServeMux
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func NewServer(cfg Config, logger *slog.Logger, db *sql.DB) *server {
	srv := &server{
		config: cfg,
		logger: logger,
		db:     db,
		mux:    http.NewServeMux(),
	}
	srv.routes()
	return srv
}
