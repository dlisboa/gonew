package server

import (
	"log/slog"
	"net/http"

	"github.com/dlisboa/gonew/app/internal/database"
)

type server struct {
	config Config
	logger *slog.Logger
	db     *database.Queries
	mux    *http.ServeMux
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func NewServer(cfg Config, logger *slog.Logger, db *database.Queries) *server {
	srv := &server{
		config: cfg,
		logger: logger,
		db:     db,
		mux:    http.NewServeMux(),
	}
	srv.routes()
	return srv
}
