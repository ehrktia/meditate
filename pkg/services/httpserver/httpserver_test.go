package httpserver

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_create_new_server(t *testing.T) {
	customPort := "9399"
	t.Run("should create new server", func(t *testing.T) {
		server, err := NewHTTPServer()
		assert.Nil(t, err)
		assert.NotNil(t, server)
	})
	t.Run("should be able to set customPort", func(t *testing.T) {
		if err := os.Setenv(httpPort, customPort); err != nil {
			t.Fatal(err)
		}
		srv, err := NewHTTPServer()
		assert.Nil(t, err)
		t.Logf("%s", srv.server.Addr)
		assert.Equal(t, "0.0.0.0:"+customPort, srv.server.Addr, "error matching server port with address")
		assert.NotNil(t, srv, "error create server with custom port")
	})
	t.Cleanup(func() {
		if err := os.Unsetenv(httpPort); err != nil {
			t.Fatal(err)
		}
	})
}
func Test_run(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s, err := NewHTTPServer()
	assert.Nil(t, err)
	t.Run("should start server", func(t *testing.T) {
		errCh := make(chan error, 1)
		go func() {
			err := s.Run(ctx)
			if err != nil {
				errCh <- err
			}
		}()
		select {
		case e := <-errCh:
			t.Fatal(e)
		case <-time.After(2 * time.Second):
			t.Log("completed")
		}
	})
}
