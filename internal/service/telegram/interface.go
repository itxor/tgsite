package telegram

import "context"

type UpdateLoopServiceInterface interface {
	StartUpdateLoop(context.Context) error
}

type ProcessingNewPostServiceInterface interface {
}
