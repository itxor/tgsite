package service

import (
	"errors"
	"fmt"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/itxor/tgsite/internal/config"
	"github.com/itxor/tgsite/internal/model"
	"github.com/itxor/tgsite/internal/repository"
	"log"
	"os"
)

type TelegramChannelService struct {
	bot            *tgbot.BotAPI
	config         *config.TelegramConfig
	repo           repository.Repository
	channelService Channel
	postService    Post
}

// NewTelegramChannelService создаёт новый инстанс TelegramChannelService
func NewTelegramChannelService(repository repository.Repository) (*TelegramChannelService, error) {
	cfg, err := config.NewTelegramConfig()
	if err != nil {
		log.Printf("Ошибка при инициализации конфига: %v", err)
	}

	bot, err := tgbot.NewBotAPI(cfg.TgBotToken)
	if err != nil {
		log.Printf("Ошибка при создании bot api: %v", err)

		return nil, errors.New("Ошибка при создании bot api")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TelegramChannelService{
		bot:    bot,
		config: cfg,
		repo:   repository,
	}, nil
}

// StartUpdatesLoop запускает цикл прослушивания и обработки сообщений из telegram-каналов
func (s *TelegramChannelService) StartUpdatesLoop() error {
	updatesChannel, err := s.getUpdates()
	if err != nil {
		os.Exit(1)
	}

	for update := range updatesChannel {
		post, err := s.handleMessage(update)
		if err != nil {
			log.Printf("Ошибка при попытке обработать сообщение из telegram: %s", err.Error())
		}

		if !s.channelService.IsExists(post.ChatId) {
			if err := s.channelService.Add(post.ChatId); err != nil {
				return err
			}
		}

		if err := s.postService.Add(post); err != nil {
			return err
		}
	}

	return nil
}

// GetUpdates получает канал обновлений бота
func (s *TelegramChannelService) getUpdates() (tgbot.UpdatesChannel, error) {
	u := tgbot.NewUpdate(0)
	u.Timeout = 60

	updates, err := s.bot.GetUpdatesChan(u)
	if err != nil {
		msg := fmt.Sprintf("Ошибка при попытке создать канал обновления: %s", err)
		log.Printf(msg)
		return nil, errors.New(msg)
	}

	return updates, nil
}

// HandleMessage обрабатывает сообщение, полученное из канала обновлений
func (s *TelegramChannelService) handleMessage(message tgbot.Update) (*model.ChannelPost, error) {
	if message.ChannelPost == nil {
		return nil, nil
	}

	post, err := s.parseMessage(*message.ChannelPost)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// parseMessage парсит сообщение и возвращает стуктуру типа model.ChannelPost
func (s *TelegramChannelService) parseMessage(message tgbot.Message) (*model.ChannelPost, error) {
	post := new(model.ChannelPost)
	var err error

	post.MessageId = message.MessageID
	post.Date = message.Date
	post.ChatId = int(message.Chat.ID)
	post.ChatName = message.Chat.Title

	if "" != message.Text {
		post.Content.Text = message.Text
	}

	if message.Entities != nil {
		post.Content.Options = s.getTextFormatting(message)
	}

	if message.Sticker != nil {
		post.Content.StickerURL, err = s.getStickerURL(message)
		if err != nil {
			return nil, err
		}
	}

	if message.Voice != nil {
		post.Content.VoiceURL, err = s.getVoiceURL(message)
		if err != nil {
			return nil, err
		}
	}

	if message.Photo != nil {
		post.Content.Photo, err = s.getPhoto(message)
		if err != nil {
			return nil, err
		}
	}

	return post, nil
}

// getFileLink получает ссылку на файл
func (s *TelegramChannelService) getFileLink(fileID string) (string, error) {
	file, err := s.bot.GetFile(tgbot.FileConfig{
		FileID: fileID,
	})
	if err != nil {
		return "", errors.New("")
	}

	return file.Link(s.config.TgBotToken), nil
}

// getStickerURL возвращает ссылку на стикер
func (s *TelegramChannelService) getStickerURL(message tgbot.Message) (string, error) {
	link, err := s.getFileLink(message.Sticker.FileID)
	if err != nil {
		msg := fmt.Sprintf("Не удалось получить ссылку на стикер: %v", err)
		log.Printf(msg)

		return "", errors.New(msg)
	}

	return link, nil
}

// getVoiceURL возвращает ссылку на звуковое сообщение
func (s *TelegramChannelService) getVoiceURL(message tgbot.Message) (string, error) {
	link, err := s.getFileLink(message.Voice.FileID)
	if err != nil {
		msg := fmt.Sprintf("Не удалось получить ссылку на звуковое соообщение: %v", err)
		log.Printf(msg)

		return "", errors.New(msg)
	}

	return link, nil
}

// getTextFormatting возвращает массив с набором форматирования текста
func (s *TelegramChannelService) getTextFormatting(message tgbot.Message) []model.Formatting {
	options := make([]model.Formatting, len(*(message.Entities)))
	for _, entity := range *(message.Entities) {
		option := model.Formatting{
			FormattingType: entity.Type,
			Offset:         entity.Offset,
			Length:         entity.Length,
		}
		options = append(options, option)
	}

	return options
}

// getPhoto возвращает набор изображения в нескольких отресайзенных копиях
func (s *TelegramChannelService) getPhoto(message tgbot.Message) ([]model.Photo, error) {
	photos := make([]model.Photo, len(*(message.Photo)))
	for _, entity := range *(message.Photo) {
		link, err := s.getFileLink(entity.FileID)
		if err != nil {
			msg := fmt.Sprintf("Не удалось получить ссылку на изображение: %s", err.Error())
			log.Printf(msg)

			return nil, errors.New(msg)
		}

		photo := model.Photo{
			URL:      link,
			Width:    entity.Width,
			Height:   entity.Height,
			FileSize: entity.FileSize,
		}

		photos = append(photos, photo)
	}

	return photos, nil
}
