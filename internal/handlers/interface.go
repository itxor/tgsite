package handlers

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
    Register(e *gin.Engine)
}
