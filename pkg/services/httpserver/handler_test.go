package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/meditate/pkg/model"
	"github.com/stretchr/testify/assert"
)

func Test_routing(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Run("respond to POST to login route", func(t *testing.T) {
		serv, err := NewHTTPServer()
		if err != nil {
			t.Fatal(err)
		}
		if err := serv.RegisterRoutes(); err != nil {
			t.Fatal(err)
		}
		errCh := make(chan error, 1)
		go func() {
			err := serv.Run(ctx)
			if err != nil {
				errCh <- err
			}
		}()
		client := http.DefaultClient
		url := fmt.Sprintf("http://%s/%s", defaultPort, "login")
		fromValues := strings.NewReader("username=testuser&password=password")
		req, err := http.NewRequest(http.MethodPost, url, fromValues)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		assert.Nil(t, err)
		resp, err := client.Do(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		select {
		case e := <-errCh:
			t.Fatal(e)
		case <-time.After(2 * time.Second):
			t.Log("time out")
			cancel()
		}
	})

}
func Test_registration(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Run("should get user reg values", func(t *testing.T) {
		serv, err := NewHTTPServer()
		if err != nil {
			t.Fatal(err)
		}
		if err := serv.RegisterRoutes(); err != nil {
			t.Fatal(err)
		}
		errCh := make(chan error, 1)
		go func() {
			err := serv.Run(ctx)
			if err != nil {
				errCh <- err
			}
		}()
		client := http.DefaultClient
		url := fmt.Sprintf("http://%s/%s", defaultPort, "register")
		fromValues := strings.NewReader("email=test@test.com&pwd=password")
		req, err := http.NewRequest(http.MethodPost, url, fromValues)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		assert.Nil(t, err)
		resp, err := client.Do(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		select {
		case e := <-errCh:
			t.Fatal(e)
		case <-time.After(2 * time.Second):
			t.Log("time out")
			cancel()
		}
	})

}
func Test_login_verify(t *testing.T) {
	testUser:=model.User{
		Email: t.Name(),
		Password: t.Name(),
	}
	t.Run("should validate the incoming parameters from login", func(t *testing.T) {
		got,err:=testUser.IsValid()
		assert.Nil(t, err)
		assert.True(t, got)
	})
}
