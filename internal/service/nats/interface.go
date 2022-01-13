package nats

const (
	NewPostsSubject = "new_posts"

	NewPostsQueue = "new_posts_queue"
)

type NatsServiceForTelegramUpdateLoopInterface interface {
	Dispatch(string, interface{}) error
}

type NatsServiceForNewPostsSubscribersInterface interface {
	SubscribeToNewPostQueue() error
}
