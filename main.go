package main

import (
	"context"

	"github.com/meditate/pkg/logging"
	"github.com/meditate/pkg/services/httpserver"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	log, err := logging.NewLogger()
	if err != nil {
		cancel()
		panic(err)
	}
	server, err := httpserver.NewHTTPServer()
	if err != nil {
		cancel()
		log.Errorf("error creating new server: %v", err)
	}
	if err := server.RegisterRoutes(); err != nil {
		cancel()
		log.Errorf("error registering server routes: %v", err)
	}
	log.Info("starting app server..")
	if err := server.Run(ctx); err != nil {
		log.Errorf("error starting server", err)
	}
}
