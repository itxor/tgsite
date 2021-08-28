package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/itxor/tgsite/internal/service"
)

type Handler struct {
	services *service.ApiServices
}

func NewHandler(services *service.ApiServices) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	catalog := router.Group("/catalog")
	{
		catalog.GET("/", h.list)
	}

	return router
}
