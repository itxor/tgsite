package service

import (
	"errors"
	"fmt"
	"github.com/itxor/tgsite/internal/repository"
)

type Channel interface {
	StartUpdatesLoop() error
}

type Service struct {
	Channel
}

func NewService(repo repository.Repository) (*Service, error) {
	tgChannelService, err := NewTelegramChannelService(repo)
	if err != nil {
		msg := fmt.Sprintf("Error initialize TelegramChannelService")

		return nil, errors.New(msg)
	}

	return &Service{
		Channel: tgChannelService,
	}, nil
}

