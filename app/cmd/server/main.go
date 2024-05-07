package main

import (
	"context"
	"log"
	"os"

	"github.com/dlisboa/gonew/app/internal/server"
)

func main() {
	ctx := context.Background()
	if err := server.Run(ctx, os.Args, os.Stdout, os.Getenv); err != nil {
		log.Fatalf("main: %s. Exiting.", err)
	}
}
