package telegram

import (
	"errors"
	"github.com/BurntSushi/toml"
	"log"
)

type TelegramConfig struct {
	TgBotToken    string `toml:"tg_bot_token"`
	TgBotUsername string `toml:"tg_bot_username"`
}

func NewTelegramConfig(path string) (*TelegramConfig, error) {
	var tgCong TelegramConfig
	_, err := toml.DecodeFile(path, &tgCong)
	if err != nil {
		log.Printf("Error for reading local config: %v", err)

		return nil, errors.New("Error for reading local config")
	}

	return &tgCong, nil
}
