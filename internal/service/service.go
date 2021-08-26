package service

import (
	"github.com/itxor/tgsite/internal/repository"
	"log"
)

type Service struct {
	Channel
	Post
	Telegram
}

func NewService(repo repository.Repository) *Service {
	tgChannelService, err := NewTelegramChannelService(repo)
	if err != nil {
		log.Fatal(err)

		return nil
	}

	return &Service{
		Telegram: tgChannelService,
		Post:     NewPostService(repo),
	}
}
