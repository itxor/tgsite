package service

import (
	"errors"
	"fmt"
	"github.com/itxor/tgsite/internal/repository"
	"log"
)

type Channel interface {
	StartUpdatesLoop()
}

type Service struct {
	Channel
}

func NewService(repo repository.Repository) (*Service, error) {
	tgChannelService, err := NewTelegramChannelService(repo)
	if err != nil {
		msg := fmt.Sprintf("Error initialize TelegramChannelService")
		log.Printf("%s: %s", msg, err.Error())

		return nil, errors.New(msg)
	}

	return &Service{
		Channel: tgChannelService,
	}, nil
}

