package httpserver

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/meditate/pkg/logging"
	"go.uber.org/zap"
)

const (
	httpPort    = "HTTPPORT"
	defaultPort = "0.0.0.0:8033"
)

type httpServer struct {
	logger *zap.SugaredLogger
	engine *gin.Engine
	server *http.Server
}

func NewHTTPServer() (*httpServer, error) {
	var port string
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	if port = os.Getenv(httpPort); port == "" {
		port = defaultPort
	}
	log, err := logging.NewLogger()
	if err != nil {
		return nil, err
	}
	return &httpServer{
		engine: r,
		logger: log,
		server: &http.Server{
			Addr:    port,
			Handler: r,
		},
	}, nil
}

func (h *httpServer) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		if err := h.server.Shutdown(ctx); err != nil && err != ctx.Err() {
			h.logger.Error(err)
		}
	}()
	if err := h.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
func (h *httpServer) RegisterRoutes() error {
	rList := createRouteList()
	rList.addRoutes()
	return h.register(rList)
}
