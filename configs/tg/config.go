package tg

import (
	"errors"
	"github.com/BurntSushi/toml"
	"log"
)

type TelegramConfig struct {
	TgBotToken     string
	TgBotUserName  string
}

func New() (*TelegramConfig, error) {
	var tgCong TelegramConfig
	_, err := toml.DecodeFile("configs/tg/config.toml", &tgCong)
	if err != nil {
		log.Printf("Error for reading local config: %v", err)

		return nil, errors.New("Error for reading local config")
	}

	return &tgCong, nil
}
