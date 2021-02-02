package httpserver

import (
	"context"
	"fmt"
	"io/ioutil"
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
		fromValues := strings.NewReader("username=testuser@test.com&password=123password")
		req, err := http.NewRequest(http.MethodPost, url, fromValues)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		assert.Nil(t, err)
		resp, err := client.Do(req)
		assert.Nil(t, err)
		if resp.StatusCode!=http.StatusOK {
			bbytes,err:=ioutil.ReadAll(resp.Body)
			assert.Nil(t, err)
			t.Logf("resp body: %+v", string(bbytes))
		}
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
		fromValues := strings.NewReader("email=test@test.com&pwd=123password")
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
func Test_parse_data(t *testing.T) {
	tests := []struct {
		name string
		uname string
		pwd string
		user *model.User
		wantErr bool
		err error
	}{
		{
			name: "valid parse",
			uname: t.Name(),
			pwd:  "123"+t.Name(),
			user: &model.User{},
			wantErr: false,
		},
		{
			name: "invalid parse",
			uname: t.Name(),
			pwd: "",
			user: &model.User{},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err:=parseFormData(test.uname, test.pwd,test.user)
			if test.wantErr && err==nil {
				t.Fatal("expected to fail")
			}
		})
	}
}
