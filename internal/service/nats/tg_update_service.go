package nats

import (
	"log"

	nats_client "github.com/itxor/tgsite/pkg/nats"
)

type tgUpdateService struct {
	client nats_client.NatsClientInterface
}

func NewNatsTgUpdateService(client nats_client.NatsClientInterface) NatsServiceForTelegramUpdateLoopInterface {
	return &tgUpdateService{
		client: client,
	}
}

func (s *tgUpdateService) Dispatch(subject string, object interface{}) error {
	defFunc, err := s.client.Connect()
	if err != nil {
		log.Fatal(err)

		return err
	}

	defer defFunc()

	if err := s.client.Publish(subject, object); err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}
