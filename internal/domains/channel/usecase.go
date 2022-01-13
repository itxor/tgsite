package channel

type service struct {
	repository ChannelRepositoryInterface
}

func NewUseCase(storage ChannelRepositoryInterface) ChannelUseCaseInterface {
	return &service{
		repository: storage,
	}
}

func (s *service) Add(chatId int) error {
	return s.repository.Add(chatId)
}

func (s *service) IsExist(chatId int) bool {
	return s.repository.IsExist(chatId)
}
