package server

import (
	"fmt"
	"net/http"

	"github.com/dlisboa/gonew/app/static"
)

func (s *server) routes() {
	s.mux.Handle("GET /{$}", s.handleIndex())

	assets := http.FileServer(http.FS(static.FS))
	s.mux.Handle("GET /static/", http.StripPrefix("/static/", assets))
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	}
}
