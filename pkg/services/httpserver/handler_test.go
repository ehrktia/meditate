package httpserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/meditate/pkg/model"
	"github.com/stretchr/testify/assert"
)

func Test_handler_response(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s, err := NewHTTPServer()
	assert.Nil(t, err)
	err = s.RegisterRoutes()
	assert.Nil(t, err)
	errCh := make(chan error, 1)
	go func() {
		err := s.Run(ctx)
		if err != nil {
			errCh <- err
		}
	}()
	t.Run("should respond to GET home path", func(t *testing.T) {
		url := fmt.Sprintf("http://%s:%s/", "0.0.0.0",defaultPort)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		assert.Nil(t, err)
		cli := http.DefaultClient
		resp, err := cli.Do(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		select {
		case e := <-errCh:
			t.Fatal(e)
		case <-time.After(2 * time.Second):
			t.Log("completed")
		}
	})
	t.Run("should respond to POST login route", func(t *testing.T) {
		url := fmt.Sprintf("http://%s:%s/%s", "0.0.0.0",defaultPort,"login")
		bbytes, err := json.Marshal(&model.User{
			Email:    "test@test.com",
			Password: "test!@£A",
		})
		assert.Nil(t, err)
		body := bytes.NewReader(bbytes)
		req, err := http.NewRequest(http.MethodPost, url, body)
		assert.Nil(t, err)
		cli := http.DefaultClient
		resp, err := cli.Do(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		select {
		case e := <-errCh:
			t.Fatal(e)
		case <-time.After(2 * time.Second):
			t.Log("completed")
		}
	})

}
