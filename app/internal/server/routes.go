package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dlisboa/gonew/app/static"
)

func (s *server) routes() {
	s.mux.Handle("GET /{$}", s.handleIndex())
	s.mux.Handle("GET /authors", s.handleAuthorsIndex())

	assets := http.FileServer(http.FS(static.FS))
	s.mux.Handle("GET /static/", http.StripPrefix("/static/", assets))
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	}
}

func (s *server) handleAuthorsIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := s.db.ListAuthors(r.Context())
		if err != nil {
			s.logger.Error("server.handleAuthorsIndex", "err", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authors)
	}
}
