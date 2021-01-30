package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/meditate/pkg/model"
	"github.com/stretchr/testify/assert"
)

func Test_routing(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Run("respond to POST to login route", func(t *testing.T) {
		server, err := NewHTTPServer()
		assert.Nil(t, err)
		assert.Nil(t, server.RegisterRoutes())
		errCh := make(chan error, 1)
		go func() {
			if err := server.Run(ctx); err != nil {
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
		data, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		gotValues := new(model.User)
		err = json.Unmarshal(data, gotValues)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Contains(t, gotValues.Email, "testuser")
		assert.Contains(t, gotValues.Password, "password")
	})
	t.Cleanup(func() {
		cancel()
	})

}
