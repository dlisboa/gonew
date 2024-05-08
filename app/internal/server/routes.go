package server

import (
	"encoding/json"
	"net/http"

	"github.com/dlisboa/gonew/app/internal/templates"
	"github.com/dlisboa/gonew/app/static"
)

func (s *server) routes() {
	s.mux.Handle("GET /{$}", s.handleIndex())
	s.mux.Handle("GET /authors", s.handleAuthorsIndex())

	assets := http.FileServer(http.FS(static.FS))
	s.mux.Handle("GET /static/", http.StripPrefix("/static/", assets))
}

func (s *server) handleIndex() http.HandlerFunc {
	tpl := templates.MustParse("layout/base", "page/home")

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "layout/base", nil)
	}
}

func (s *server) handleAuthorsIndex() http.HandlerFunc {
	errtpl := templates.MustParse("layout/error", "error/500")

	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := s.db.ListAuthors(r.Context())

		if err != nil {
			s.logger.Error("server.handleAuthorsIndex", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			errtpl.ExecuteTemplate(w, "layout/error", nil)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authors)
	}
}
