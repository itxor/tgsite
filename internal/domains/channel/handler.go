package channel

import (
	"github.com/gin-gonic/gin"
	"github.com/itxor/tgsite/internal/handlers"
)

const (
	testHandlerForChannelAction = "/test"
)

type handler struct {
    useCase ChannelUseCaseInterface
}

func NewChannelHandler(
    useCase ChannelUseCaseInterface,
) handlers.HandlerInterface {
    return &handler{
        useCase: useCase,
    }
}

func (h *handler) channelsHandler(c *gin.Context) {
    h.useCase.FindAll()

	c.JSON(200, gin.H{
		"channels": "",
	})
}

func (h *handler) Register(e *gin.Engine) {
	e.GET(testHandlerForChannelAction, h.channelsHandler)
}
