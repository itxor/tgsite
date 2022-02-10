package nats

import (
	"context"
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

func (s *postSubscribeService) SubscribeToNewPostQueue(ctx context.Context) error {
	defFunc, err := s.client.Connect()
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer defFunc()

	ch := make(chan post.Post)
	_, err = s.client.GetConnect().BindRecvQueueChan(NewPostsSubject, NewPostsQueue, ch)
	if err != nil {
		logrus.Error(err)
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case msgPost := <-ch:
			logrus.Print(msgPost)

			if !s.channelUseCase.IsExist(msgPost.ChatId) {
				if err := s.channelUseCase.Add(msgPost.ChatId); err != nil {
					logrus.Error(err)
				}
			}

			if err := s.postUseCase.Add(msgPost); err != nil {
				logrus.Error(err)
			}

			break
		}
	}
}
