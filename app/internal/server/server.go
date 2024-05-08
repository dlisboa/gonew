package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	config Config
	logger *slog.Logger
	db     *sql.DB
	mux    *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func NewServer(cfg Config, logger *slog.Logger, db *sql.DB) *Server {
	srv := &Server{
		config: cfg,
		logger: logger,
		db:     db,
		mux:    http.NewServeMux(),
	}
	srv.routes()
	return srv
}

func (s *Server) routes() {
	s.mux.Handle("GET /{$}", s.Index())

	assets := http.FileServer(http.Dir("./public"))
	s.mux.Handle("GET /public/", http.StripPrefix("/public/", assets))
}

func (s *Server) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
}
