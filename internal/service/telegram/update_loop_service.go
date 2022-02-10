package telegram

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/itxor/tgsite/internal/domains/post"
	"github.com/itxor/tgsite/pkg/telegram"
)

type UpdateLoopService struct {
	tg          telegram.TelegramClientInterface
	postUseCase post.PostUseCaseForUpdateTelegramLoopInterface
}

func NewUpdateLoopService(
	client telegram.TelegramClientInterface,
	postUseCase post.PostUseCaseForUpdateTelegramLoopInterface,
) UpdateLoopServiceInterface {
	return &UpdateLoopService{
		tg:          client,
		postUseCase: postUseCase,
	}
}

// StartUpdateLoop запускает цикл, получающий сообщения из бота и отправляющий их в брокер для дальнейшей обработки
func (s *UpdateLoopService) StartUpdateLoop(ctx context.Context) error {
	for {
		select {
		case message := <-s.tg.GetUpdateChan():
			newPost, err := s.postUseCase.BuildNewPostFromMessage(message)
			if err != nil {
				logrus.Error(err)
				break
			}

			if err := s.postUseCase.DispatchAddPost(*newPost); err != nil {
				logrus.Error(err)
				break
			}

			break
		case <-ctx.Done():
			return nil
		}
	}

	return nil
}
