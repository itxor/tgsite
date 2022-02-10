package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
}

func NewServer() *Server {
	return &Server{
		gin: gin.Default(),
	}
}

func (s *Server) RegisterHandlers(handlers ...HandlerInterface) {
	for _, handler := range handlers {
		handler.Register(s.gin)
	}
}

func (s *Server) Start(ctx context.Context) error {
	return s.gin.Run()
}
