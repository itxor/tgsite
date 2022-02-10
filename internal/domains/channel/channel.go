package channel

const (
	DatabaseChannels = "channels"
)

type Channel struct {
    
}

type ChannelUseCaseInterface interface {
	Add(int) error
	IsExist(int) bool
    FindAll() 
}

type ChannelRepositoryInterface interface {
	Add(int) error
	IsExist(int) bool
}
