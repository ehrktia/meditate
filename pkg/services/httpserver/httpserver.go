package httpserver

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	httpPort    = "HTTPPORT"
	defaultPort = "0.0.0.0:8033"
)
type httpServer struct {
	engine *gin.Engine
	server *http.Server
}

func NewHTTPServer() *httpServer {
	var port string
	r:=gin.New()
	if port = os.Getenv(httpPort); port == "" {
		port=defaultPort
	}
	return &httpServer{
		engine: r,
		server: &http.Server {
				Addr: port,
				Handler:r,
			},
		}
}

func (h *httpServer)Run(ctx context.Context) error {
	go func()error {
	<- ctx.Done()
	if err:=h.server.Shutdown(ctx);err!=nil {
		return err
	}
	return nil
	}()
	if err := h.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
func (h *httpServer) RegisterRoutes() error {
	rList:=createRouteList()
	rList.addRoutes()
	return h.register(rList)
}

