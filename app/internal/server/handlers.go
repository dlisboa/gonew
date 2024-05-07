package server

import (
	"database/sql"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/dlisboa/gonew/app/internal/templates"
)

type Handlers struct {
	config config
	logger *slog.Logger
	db     *sql.DB
}

func (h *Handlers) NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", h.Index())

	assets := http.FileServer(http.Dir("./public"))
	mux.Handle("GET /public/", http.StripPrefix("/public/", assets))

	return mux
}

func (h *Handlers) Index() http.Handler {
	tmpl := template.Must(template.ParseFS(templates.FS, "pages/index.html"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
