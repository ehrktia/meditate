package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meditate/pkg/model"
)

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "welcome",
	})
}

func loginHandler(c *gin.Context) {
	u := new(model.User)
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
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
		default:
			return fmt.Errorf("not a valid method")
		}
	}
	return nil
}
func logout(c *gin.Context) {
	u := new(model.User)
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, nil)
}
