package service

import (
	"github.com/itxor/tgsite/internal/model"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"log"
)

const (
	TopicNewPost = "new_post"
)

func PublishNewPost(post *model.ChannelPost) error {
	conn, deferFunc, err := Connect()
	if err != nil {
		return err
	}
	defer deferFunc()

	ec, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	if err := ec.Publish(TopicNewPost, post); err != nil {
		return err
	}

	return nil
}

func SubscribeToNewPost() error {
	_, f, err := Connect()
	if err != nil {
		return err
	}
	defer f()
	return nil
}

func Connect() (*nats.Conn, func(), error) {
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