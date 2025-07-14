package main

import (
	"context"

	"github.com/eyagovbusiness/GSWB.Users/internal/di"
)

func main() {
	server := di.InitializeServer()

	if err := server.Start(context.Background()); err != nil {
		panic(err)
	}
}
