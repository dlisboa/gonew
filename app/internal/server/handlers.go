package server

import (
	"database/sql"
	"embed"
	"html/template"
	"log/slog"
	"net/http"
)

type Handlers struct {
	config config
	logger *slog.Logger
	db     *sql.DB
}

func (h *Handlers) NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", h.Index())

	assets := http.FileServer(http.Dir("./public"))
	mux.Handle("GET /public/", http.StripPrefix("/public/", assets))

	return mux
}

//go:embed templates
var templatesFS embed.FS

func (h *Handlers) Index() http.Handler {
	tmpl := template.Must(template.ParseFS(templatesFS, "templates/index.html.tmpl"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
