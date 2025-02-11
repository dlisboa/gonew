package app

import (
	"log/slog"
	"net/http"
	"time"
)

type application struct {
	*http.Server
	log *slog.Logger
}

func New(port string, log *slog.Logger) (*application, error) {
	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	app := &application{Server: srv, log: log}
	app.Server.Handler = app.routes()
	return app, nil
}
