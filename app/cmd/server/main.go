package main

import (
	"log"
	// add the internal/server package here
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
