package server

import (
    "fmt"
    "go.uber.org/zap"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Server struct {
    router    *gin.Engine
    port      string
    logger    *zap.Logger
    httpSrv   *http.Server
}

func NewServer(router *gin.Engine, port string, logger *zap.Logger) *Server {
    return &Server{
        router: router,
        port:   port,
        logger: logger,
    }
}

func (s *Server) Start() error {
    addr := fmt.Sprintf(":%s", s.port)
    s.httpSrv = &http.Server{
        Addr:    addr,
        Handler: s.router,
    }

    s.logger.Info("starting server", zap.String("address", addr))
    return s.httpSrv.ListenAndServe()
}