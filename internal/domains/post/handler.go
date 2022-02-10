package post

import (
	"github.com/gin-gonic/gin"
	"github.com/itxor/tgsite/internal/handlers"
)

const (
    testRoute = "/test"
)

type handler struct {
}

func CreatePostHandler() handlers.HandlerInterface {
    return &handler{}
}

func (h *handler) Test(c *gin.Context) {
}

func (h *handler) Register(e *gin.Engine) {
    
}
