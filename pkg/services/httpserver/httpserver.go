package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/meditate/pkg/logging"
	"go.uber.org/zap"
)

const (
	httpPort    = "PORT"
	defaultPort = "8083"
)

type httpServer struct {
	engine *gin.Engine
	logger *zap.SugaredLogger
	server *http.Server
}

func NewHTTPServer() (*httpServer, error) {
	var port string
	log, err := logging.NewLogger()
	if err != nil {
		return nil, err
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowMethods=[]string{"GET","POST","OPTIONS"}
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials=true
	config.AllowHeaders=[]string{"Content-Type"}
	r.Use(cors.New(config))
	if port = os.Getenv(httpPort); port == "" {
		port = defaultPort
	}
	log.Info("server initalised in address ", port)
	h:=&httpServer{
		engine: r,
		logger: log,
		server: &http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%s", port),
			Handler: r,
		},
	}
	if err:=h.RegisterRoutes();err!=nil{
		return nil,err
	}
	log.Info("routes registration completed")
	return h, nil
}
func (s *httpServer) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		if err := s.server.Shutdown(ctx); err != nil {
			s.logger.Errorf("error closing server: %v", err)
		}
	}()
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (h *httpServer) RegisterRoutes() error {
	rList := &routeList{routeList: []*routes{}}
	rList.addRoutes()
	return h.register(rList)
}
