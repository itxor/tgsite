package channel

const (
	DatabaseChannels = "channels"
)

type ChannelUseCaseInterface interface {
	Add(int) error
	IsExist(int) bool
}

type ChannelRepositoryInterface interface {
	Add(int) error
	IsExist(int) bool
}
