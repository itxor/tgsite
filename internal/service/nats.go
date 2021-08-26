package service

import (
	"github.com/itxor/tgsite/internal/model"
	"github.com/itxor/tgsite/internal/repository"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
)

const (
	TopicNewPost = "new_post"
)

type Nats struct {
	services Service
	conn     *nats.EncodedConn
}

func NewNats() *Nats {
	db, ctx, err := repository.NewMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	return &Nats{
		services: *NewService(repository.NewRepository(db, ctx)),
	}
}

func (s *Nats) ConnectToMessageBus() (func(), error) {
	conn, deferFunc, err := connect()
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	ec, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	s.conn = ec

	df := func() {
		deferFunc()
		ec.Close()
	}

	return df, nil
}

func (s *Nats) SubscribeNewChannelPosts(wg *sync.WaitGroup) {
	ch := make(chan *model.ChannelPost)
	_, err := s.conn.BindRecvChan(TopicNewPost, ch)
	if err != nil {
		return
	}

	for post := range ch {
		// todo: логировать в отдельный файл журнала
		logrus.Debug(post)
		if !s.services.Channel.IsExists(post.ChatId) {
			if err := s.services.Channel.Add(post.ChatId); err != nil {
				logrus.Errorf("Ошибка при попытке создать канал: %s", err.Error())
			}
		}

		if err := s.services.Post.Add(post); err != nil {
			logrus.Errorf("Ошибка при попытке добавить новый пост: %s", err.Error())
		}
	}

	wg.Done()
}

func (s *Nats) PublishNewPost(post *model.ChannelPost) error {
	if err := s.conn.Publish(TopicNewPost, post); err != nil {
		return err
	}

	return nil
}

func connect() (*nats.Conn, func(), error) {
	nc, err := nats.Connect(
		nats.DefaultURL,
		nats.Name("Telegram channel parser message bus"),
		nats.ErrorHandler(func(conn *nats.Conn, subscription *nats.Subscription, err error) {
			if subscription != nil {
				logrus.Error(
					"Async error in %q/%q: %v",
					subscription.Subject,
					subscription.Queue,
					err,
				)
			} else {
				logrus.Error("Async error outside subscription: %v", err)
			}
		}),
	)
	if err != nil {
		logrus.Error("Ошибка при попытке установить соединение с nats: %s", err.Error())

		return nil, nil, err
	}

	return nc, nc.Close, err
}
