package server

import (
	"net/http"

	"github.com/Benjosh95/auto-voc/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.ServerConfig, handler *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.Address,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
