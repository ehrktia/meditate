package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)
func startServer(ctx context.Context,t *testing.T,errch chan<- error)  {
	server:=NewHTTPServer()
	t.Helper()
	if err:=server.RegisterRoutes();err!=nil {
		errch<-err
	}
	if err:=server.Run(ctx);err!=nil {
		errch<-err
	}
}

func Test_routing(t *testing.T) {
	ctx,cancel:=context.WithCancel(context.Background())
	errch:=make(chan error, 1)
	go startServer(ctx,t,errch)
	t.Run("respond to POST to login route", func(t *testing.T) {
		client:=http.DefaultClient
		url:=fmt.Sprintf("http://%s/%s", defaultPort,"login")
		req,err:=http.NewRequest(http.MethodPost, url, nil)
		assert.Nil(t, err)
		resp,err:=client.Do(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Cleanup(func() {
		cancel()
	})

}

