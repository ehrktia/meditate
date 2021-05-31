package httpserver

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/require"
)

func Test_handler_response(t *testing.T) {
	r := require.New(t)
	testServer, err := NewHTTPServer()
	r.Nil(err, "error initializing server for test")
	t.Run("should respond to GET home path", func(t *testing.T) {
		url := buildURL()
		expectedResult := `{"status":"ok"}`
		apitest.New(t.Name()).
			Debug().
			Handler(testServer.engine).
			Get(url).
			Expect(t).
			Body(expectedResult).
			Status(http.StatusOK).
			End()
	})
	t.Run("should respond to POST login route", func(t *testing.T) {
		r.Nil(err, "error marshalling data for req")
		respBody := `{"token":"testToken!@1"}`
		url := buildURL("login")
		apitest.New(t.Name()).
			Debug().
			Handler(testServer.engine).
			Post(url).
			FormData("idtoken", "testToken!@1").
			Expect(t).
			Body(respBody).
			Status(http.StatusOK).End()
	})
}
func buildURL(path ...string) string {
	if len(path) < 1 {
		return fmt.Sprintf("http://0.0.0.0:%s/", defaultPort)
	}
	return fmt.Sprintf("http://0.0.0.0:%s/%s", defaultPort, path[0])
}
