package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{})
}
