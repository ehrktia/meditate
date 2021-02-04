package main

import (
	"context"

	"github.com/meditate/pkg/services/httpserver"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server, err := httpserver.NewHTTPServer()
	if err != nil {
		panic(err)
	}
	if err := server.RegisterRoutes(); err != nil {
		panic(err)
	}
	if err := server.Run(ctx); err != nil {
		panic(err)
	}
}
