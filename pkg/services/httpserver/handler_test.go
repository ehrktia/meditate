package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)
func addTestRoute() {
	r:=createRouteList()
	r.addRoutes(routes{
		path: "/login",
		method: http.MethodPost,
		handler: loginHandler,
	})
}

func Test_create_handler(t *testing.T) {
	addTestRoute()
	client:=http.Client{}
	t.Run("should respond to login", func(t *testing.T) {
		srv:=httptest.NewServer(http.HandlerFunc(loginHandler))
		defer srv.Close()
	url:=srv.URL+"/login"
	req,err:=http.NewRequest(http.MethodPost, url, nil)
	assert.Nil(t,err)
	res,err:=client.Do(req)
	assert.Nil(t, err)
	defer func(){assert.Nil(t,res.Body.Close())}()
})
}
