package nats

import (
	"github.com/itxor/tgsite/pkg/nats"
	"log"
)

type service struct {
	client nats.NatsClientInterface
}

func NewService() NatsServiceInterface {
	return &service{
		client: nats.NewClient(),
	}
}

func (s *service) Dispatch(queue string, object interface{}) error {
	defFunc, err := s.client.Connect()
	if err != nil {
		log.Fatal(err)

		return err
	}

	defer defFunc()

	if err := s.client.GetConnect().Publish(queue, object); err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}
