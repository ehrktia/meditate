package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meditate/pkg/model"
)

func parseFormData(name, password string, user *model.User) error {
	user.Email = name
	user.Password = password
	if err := user.IsValid(); err != nil {
		return err
	}
	return nil
}

func loginHandler(c *gin.Context) {
	u := new(model.User)
	if err := parseFormData(c.PostForm("username"), c.PostForm("password"), u); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"username": u.Email,
			"error":    err,
		})
	}
	c.JSON(http.StatusOK, nil)
}
func registerHandler(c *gin.Context) {
	u := new(model.User)
	if err := parseFormData(c.PostForm("email"), c.PostForm("pwd"), u); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"username": u.Email,
			"error":    err,
		})
	}
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
