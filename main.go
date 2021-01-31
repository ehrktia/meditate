package main

import (
	"context"

	"github.com/meditate/pkg/services/httpserver"
	"golang.org/x/sync/errgroup"
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
	e, ctxerrgroup := errgroup.WithContext(ctx)
	e.Go(func() error {
		return server.Run(ctxerrgroup)
	})
	if e.Wait() != nil {
		panic(e.Wait().Error())
	}
}
