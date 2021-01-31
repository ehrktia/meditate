package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meditate/pkg/model"
)

func loginHandler(c *gin.Context) {
	u := &model.User{}
	u.Email = c.PostForm("username")
	u.Password = c.PostForm("password")
	 valid,validErr:=u.IsValid()
	 if !valid {
		 c.JSON(http.StatusNotAcceptable, gin.H{
			 "user":u.Email,
			 "error": validErr.Error(),
		 })
	 }
	c.JSON(http.StatusOK,nil)
}
func registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (h *httpServer) register(routes *routeList) error {
	for _, route := range routes.routeList {
		switch route.method {
		case http.MethodGet:
			h.engine.GET(route.path, route.handler)
		case http.MethodPost:
			h.engine.POST(route.path, route.handler)
		case http.MethodDelete:
			h.engine.DELETE(route.path, route.handler)
		case http.MethodPatch:
			h.engine.PATCH(route.path, route.handler)
		default:
			return fmt.Errorf("not a valid http action")
		}
	}
	return nil
}

