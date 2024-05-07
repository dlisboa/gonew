package server

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/unrolled/render"
)

type Handlers struct {
	config config
	logger *slog.Logger
	db     *sql.DB
	render *render.Render
}

func (h *Handlers) NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", h.Index())

	assets := http.FileServer(http.Dir("./public"))
	mux.Handle("GET /public/", http.StripPrefix("/public/", assets))

	return mux
}

func (h *Handlers) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.render.HTML(w, http.StatusOK, "index.html", nil)
	})
}
