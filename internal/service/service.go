package service

import (
	"errors"
	"fmt"
	"github.com/itxor/tgsite/internal/repository"
)

type Telegram interface {
	StartUpdatesLoop() error
}

type Channel interface {
	List()
	Add(int) error
	Delete()
	Update()
	IsExists(int) bool
}

type Post interface {
	List()
	Add()
	Delete()
	Update()
}

type Service struct {
	Telegram
	Channel
	Post
}

func NewService(repo repository.Repository) (*Service, error) {
	channelService := NewChannelService(repo)
	tgChannelService, err := NewTelegramChannelService(repo, channelService)
	if err != nil {
		msg := fmt.Sprintf("Error initialize TelegramChannelService")

		return nil, errors.New(msg)
	}

	return &Service{
		Telegram: tgChannelService,
	}, nil
}

