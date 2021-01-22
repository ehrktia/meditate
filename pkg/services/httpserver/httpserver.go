package httpserver

import (
	"context"
	"net/http"
	"os"
)
const (
    httpPort = "HTTPPORT"
    defaultPort = "0.0.0.0:8033"
)
func NewHTTPServer() *http.Server {
    var port string
    if port=os.Getenv(httpPort);port=="" {
        port="0.0.0.0:8033"
    }
    return &http.Server{
        Addr:port,
        Handler: http.DefaultServeMux,
    }
}
func Run(ctx context.Context,server *http.Server) error {
    if err:=server.ListenAndServe();err!=nil  {
        return err
    }
    <-ctx.Done()
        if err:=server.Close();err!=nil {
            return err
        }
        return nil
}

