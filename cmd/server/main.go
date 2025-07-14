// @title GSWB Users API
// @version 1.0
// @description API for managing users.
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"context"
	"log"

	"github.com/eyagovbusiness/GSWB.Users/internal/di"
)

func main() {
	server, err := di.InitializeServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	if err := server.Start(context.Background()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
