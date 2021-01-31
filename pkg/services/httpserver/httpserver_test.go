package httpserver

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_create_new_server(t *testing.T) {
	customPort := "0.0.0.0:9999"
	s, err := NewHTTPServer()
	t.Run("should create new server", func(t *testing.T) {
		assert.Nil(t, err)
		assert.NotNil(t, s)
	})
	t.Run("should be able to set customPort", func(t *testing.T) {
		if err := os.Setenv(httpPort, customPort); err != nil {
			t.Fatal(err)
		}
		srv, err := NewHTTPServer()
		assert.Nil(t, err)
		assert.Equal(t, srv.server.Addr, customPort)
	})
	t.Cleanup(func() {
		if err := os.Unsetenv(httpPort); err != nil {
			t.Fatal(err)
		}
	})
}
func Test_run(t *testing.T) {
	server, err := NewHTTPServer()
	assert.Nil(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	t.Run("should start server", func(t *testing.T) {
		errCh := make(chan error, 1)
		go func() {
			if err := server.Run(ctx); err != nil {
				errCh <- err
			}
		}()
		select {
		case e := <-errCh:
			t.Fatal(e)
		case <-time.After(1 * time.Second):
			cancel()
			t.Log("completed")
		}
	})
}
func Test_routes(t *testing.T) {
	s, err := NewHTTPServer()
	assert.Nil(t, err)
	t.Run("should add routes to server", func(t *testing.T) {
		assert.Nil(t, s.RegisterRoutes())
	})
}
