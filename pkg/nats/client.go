package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"log"
)

type client struct {
	conn *nats.EncodedConn
}

func NewClient() NatsClientInterface {
	return &client{}
}

func (s *client) Connect() (func(), error) {
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

func (s *client) GetConnect() *nats.EncodedConn {
	return s.conn
}

func connect() (*nats.Conn, func(), error) {
	nc, err := nats.Connect(
		nats.DefaultURL,
		nats.Name("Message bus"),
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
