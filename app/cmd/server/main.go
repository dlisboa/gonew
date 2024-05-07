package main

import (
	"context"
	"log"
	"os"

	"github.com/dlisboa/gonew/app/internal/server"
)

func main() {
	ctx := context.Background()
	if err := server.Run(ctx, os.Args, os.Getenv, os.Stdout); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
