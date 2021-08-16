package main

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/itxor/tgsite/configs/tg"
	"log"
	"os"
)

type TelegramChannelService struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramChannelService() (*TelegramChannelService, error) {
	cfg, err := tg.New()
	if err != nil {
		log.Fatalf("Ошибка при инициализации конфига: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TgBotToken)
	if err != nil {
		log.Fatalf("Ошибка при создании bot api: %v", err)

		return nil, errors.New("Ошибка при создании bot api")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TelegramChannelService{
		bot: bot,
	}, nil
}

func (s *TelegramChannelService) getUpdates() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := s.bot.GetUpdatesChan(u)
	if err != nil {
		msg := fmt.Sprintf("Ошибка при попытке создать канал обновления: %s", err)
		log.Fatalf(msg)
		return nil, errors.New(msg)
	}

	return updates, nil
}

func (s *TelegramChannelService) handleMessage(update tgbotapi.Update) {
	log.Printf("[%s]", update.ChannelPost.Text)

	msg := tgbotapi.NewMessage(update.ChannelPost.Chat.ID, update.ChannelPost.Text)
	msg.ReplyToMessageID = update.ChannelPost.MessageID

	_, err := s.bot.Send(msg)
	if err != nil {
		return
	}
}

func (s *TelegramChannelService) getFileByFileID(fileID string) (*tgbotapi.File, error) {
	file, err := s.bot.GetFile(tgbotapi.FileConfig{
		FileID: fileID,
	})
	if err != nil {
		return nil, errors.New("")
	}

	return &file, nil
}

func main() {
	tgChannelService, err := NewTelegramChannelService()
	if err != nil {
		os.Exit(1)
	}

	updatesChannel, err := tgChannelService.getUpdates()
	if err != nil {
		os.Exit(1)
	}

	for update := range updatesChannel {
		// игнорируем любые путые сообщения
		if update.ChannelPost == nil {
			continue
		}

		tgChannelService.handleMessage(update)
	}
}

