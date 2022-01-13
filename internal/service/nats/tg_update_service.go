package nats

import (
	nats_client "github.com/itxor/tgsite/pkg/nats"
	"log"
)

type tgUpdateService struct {
	client nats_client.NatsClientInterface
}

func NewNatsTgUpdateService() NatsServiceForTelegramUpdateLoopInterface {
	return &tgUpdateService{
		client: nats_client.NewClient(),
	}
}

func (s *tgUpdateService) Dispatch(subject string, object interface{}) error {
	defFunc, err := s.client.Connect()
	if err != nil {
		log.Fatal(err)

		return err
	}

	defer defFunc()

	if err := s.client.GetConnect().Publish(subject, object); err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}
