package nats

import (
	"sync"

	"github.com/itxor/tgsite/internal/domains/channel"
	"github.com/itxor/tgsite/internal/domains/post"
	nats_client "github.com/itxor/tgsite/pkg/nats"
	"github.com/sirupsen/logrus"
)

type postSubscribeService struct {
	client         nats_client.NatsClientInterface
	postUseCase    post.PostUseCaseForSubscribeNewPostsInterface
	channelUseCase channel.ChannelUseCaseInterface
}

func NewNatsPostSubscribeService(
	postUseCase post.PostUseCaseForSubscribeNewPostsInterface,
	channelUseCase channel.ChannelUseCaseInterface,
) NatsServiceForNewPostsSubscribersInterface {
	return &postSubscribeService{
		client:         nats_client.NewClient(),
		postUseCase:    postUseCase,
		channelUseCase: channelUseCase,
	}
}

func (s *postSubscribeService) SubscribeToNewPostQueue() error {
	defFunc, err := s.client.Connect()
	if err != nil {
		logrus.Error(err)

		return err
	}

	defer defFunc()

	wg := sync.WaitGroup{}
	for {
		wg.Add(1)
		if _, err := s.client.GetConnect().QueueSubscribe(
			NewPostsSubject,
			NewPostsQueue,
			func(post post.Post) {
				if !s.channelUseCase.IsExist(post.ChatId) {
					if err := s.channelUseCase.Add(post.ChatId); err != nil {
						logrus.Error(err)
					}
				}

				if err := s.postUseCase.Add(post); err != nil {
					logrus.Error(err)
				}

				wg.Done()
			}); err != nil {
			logrus.Error(err)

			continue
		}

		wg.Wait()
	}

	return nil
}
