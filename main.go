package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/web-alytics/meditate/pkg/logging"
	"github.com/web-alytics/meditate/pkg/services/httpserver"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	prodLogger, err := zap.NewProduction()
	if err != nil {
		cancel()
		panic(err)
	}
	log, err := logging.NewLogger(prodLogger.Sugar())
	if err != nil {
		cancel()
		panic(err)
	}
	engine := gin.New()
	server, err := httpserver.NewHTTPServer(log, engine)
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
