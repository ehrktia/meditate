package httpserver

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/web-alytics/meditate/pkg/services/httpserver/mocks"
)

func Test_create_new_server(t *testing.T) {
	customPort := "9399"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTracer := mocks.NewMockTracer(mockCtrl)
	t.Run("should create new server", func(t *testing.T) {
		server, err := NewHTTPServer(context.Background(), gin.Default(), mockTracer)
		assert.Nil(t, err)
		assert.NotNil(t, server)
	})
	t.Run("should be able to set customPort", func(t *testing.T) {
		if err := os.Setenv(httpPort, customPort); err != nil {
			t.Fatal(err)
		}
		srv, err := NewHTTPServer(context.Background(), gin.Default(), mockTracer)
		assert.Nil(t, err)
		assert.Equal(t, srv.Server.Addr, ":"+customPort)
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
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTracer := mocks.NewMockTracer(mockCtrl)
	s, err := NewHTTPServer(context.Background(), gin.Default(), mockTracer)
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
