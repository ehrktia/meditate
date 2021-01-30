package httpserver

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)
func Test_add_new_route(t *testing.T) {
		r:=createRouteList()
	t.Run("should create a new route list", func(t *testing.T) {
		assert.NotEmpty(t, r)
	})
	 t.Run("should be able to add new route to routelist", func(t *testing.T) {
		route:=routes{
			 path: "/login",
			method: http.MethodPost,
			handler: loginHandler,
		}
		r.addRoute(route)
		if len(r.routeList)<1 {
			t.Error("can not add route")

		}
	})

}
