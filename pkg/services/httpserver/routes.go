package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type routes struct {
	path    string
	method  string
	handler gin.HandlerFunc
}
type routeList struct {
	routeList []*routes
}

func createRouteList() *routeList {
	r := []*routes{}
	return &routeList{routeList: r}
}

func (r *routeList) addRoutes() {
	r.routeList = append(r.routeList,
		&routes{
			path:    "/",
			method:  http.MethodGet,
			handler: homeHandler,
		},
		&routes{
			path:    "/login",
			method:  http.MethodPost,
			handler: loginHandler,
		},
	)
}
