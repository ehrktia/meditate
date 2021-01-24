package main

import (
	"context"

	"github.com/meditate/pkg/services/httpserver"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := httpserver.NewHTTPServer()
	e, ctxerrgroup := errgroup.WithContext(ctx)
	e.Go(func() error {
		return httpserver.Run(ctxerrgroup, server)
	})
	if e.Wait() != nil {
		panic(e.Wait().Error())
	}
}
