package nats

const (
	AddNewPostQueue = "add_new_post_queue"
)

type NatsServiceInterface interface {
	Dispatch(string, interface{}) error
}
