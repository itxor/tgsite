package service

import (
	"errors"
	"fmt"
	"github.com/itxor/tgsite/internal/model"
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
	Add(*model.ChannelPost) error
	List()
	Delete()
	Update()
}

type Service struct {
	Channel
	Post
	Telegram
}

func NewService(repo repository.Repository) (*Service, error) {
	tgChannelService, err := NewTelegramChannelService(repo)
	if err != nil {
		msg := fmt.Sprintf("Error initialize TelegramChannelService")

		return nil, errors.New(msg)
	}

	return &Service{
		Telegram: tgChannelService,
		Post:     NewPostService(repo),
	}, nil
}
