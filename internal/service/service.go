package service

import (
	"github.com/itxor/tgsite/internal/repository"
	"github.com/sirupsen/logrus"
)

type ApiServices struct {
	Channel
	Post
}

type TelegramParserService struct {
	Telegram
}

func NewTelegramParserService(repo repository.Repository) *TelegramParserService {
	telegramService, err := NewTelegramChannelService(repo)
	if err != nil {
		logrus.Fatal(err)
	}

	return &TelegramParserService{telegramService}
}

func NewAPIServices(repo repository.Repository) *ApiServices {
	return &ApiServices{
		Post:    NewPostService(repo),
		Channel: NewChannelService(repo),
	}
}
