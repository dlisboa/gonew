package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dlisboa/gonew/app/internal/templates"
	"github.com/dlisboa/gonew/app/static"
)

func (s *server) routes() {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", s.handleIndex())
	mux.Handle("GET /authors/", s.handleAuthorsIndex())
	mux.Handle("GET /authors/{id}", s.handleAuthorsShow())

	assets := http.FileServer(http.FS(static.FS))
	mux.Handle("GET /static/", http.StripPrefix("/static/", assets))

	s.handler = mux
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
			s.logger.Error("cannot list authors", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			errtpl.ExecuteTemplate(w, "layout/error", nil)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authors)
	}
}

func (s *server) handleAuthorsShow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid author id"))
			return
		}

		author, err := s.db.GetAuthor(r.Context(), int64(id))

		if err != nil {
			s.logger.Error("cannot get author", "err", err)
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(author)
	}
}
