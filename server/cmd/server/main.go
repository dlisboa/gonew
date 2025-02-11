package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dlisboa/go-templates/server/internal/app"
)

func main() {
	ctx := context.Background()
	// Handle signals for graceful shutdown.
	nctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))

	srv, err := app.New(port, log)
	if err != nil {
		log.Error("cannot init application", "err", err)
		os.Exit(1)
	}

	go func() {
		log.Info("starting HTTP server", "addr", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error("server closed", "err", err)
			os.Exit(1)
		}
	}()

	<-nctx.Done()
	log.Info("received signal; shutting down")
	// Shutdown timeout should be no less than the server's WriteTimeout.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Info("shutdown")
}
