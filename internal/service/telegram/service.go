package telegram

import (
	"github.com/itxor/tgsite/internal/domains/post"
	"github.com/itxor/tgsite/pkg/telegram"
	"log"
)

type UpdateLoopService struct {
	tg          telegram.TelegramClientInterface
	postUseCase post.PostUserCaseInterface
}

func NewUpdateLoopService(
	client telegram.TelegramClientInterface,
	postRepository post.PostRepositoryInterface,
) UpdateLoopServiceInterface {
	return &UpdateLoopService{
		tg:          client,
		postUseCase: post.NewUseCase(postRepository),
	}
}

// StartUpdateLoop запускает цикл, получающий сообщения из бота и отправляющий их в брокер для дальнейшей обработки
func (s *UpdateLoopService) StartUpdateLoop() error {
	for message := range s.tg.GetUpdateChan() {
		// вероятна стратегия, в случае, если в бот будут приходить не только сообщения, но и команды
		newPost, err := s.postUseCase.BuildNewPostFromMessage(message)
		if err != nil {
			log.Fatal(err)
			continue
		}

		if err := s.postUseCase.DispatchAddPost(*newPost); err != nil {
			log.Fatal(err)
			continue
		}
	}

	return nil
}
