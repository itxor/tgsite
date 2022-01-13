package telegram

type UpdateLoopServiceInterface interface {
	StartUpdateLoop() error
}

type ProcessingNewPostServiceInterface interface {
}
