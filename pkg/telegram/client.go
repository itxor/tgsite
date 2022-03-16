package telegram

import (
	"errors"
	"fmt"
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type client struct {
	client *tgbot.BotAPI
	config *TelegramConfig

	ch *chan MessageDTO
}

func NewClient(path string) (TelegramClientInterface, error) {
	cfg, err := NewTelegramConfig(path)
	if err != nil {
		log.Printf("Ошибка при инициализации конфига: %v", err)
	}

	bot, err := tgbot.NewBotAPI(cfg.TgBotToken)
	if err != nil {
		log.Printf("Ошибка при создании bot api: %v", err)

		return nil, errors.New("Ошибка при создании bot api")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	if err != nil {
		// todo: добавить логи, преобразовать ошибку
		return nil, err
	}

	return &client{
		client: bot,
		config: cfg,
	}, nil
}

func (s *client) GetUpdateChan() chan MessageDTO {
	if s.ch != nil {
		return *s.ch
	}

	ch, err := s.client.GetUpdatesChan(tgbot.UpdateConfig{})
	if err != nil {
		log.Fatal(err)

		return nil
	}

	customCh := make(chan MessageDTO)
	go func(customCh chan MessageDTO, ch tgbot.UpdatesChannel) {
		for update := range ch {
			fmt.Printf("%#v", update)
			customCh <- s.transformUpdateToMessageDTO(update)
		}
	}(customCh, ch)
	s.ch = &customCh

	return *s.ch
}

func (s *client) GetStickerURL(message MessageDTO) (string, error) {
	link, err := s.getFileLink(*message.StickerFileID)
	if err != nil {
		msg := fmt.Sprintf("Не удалось получить ссылку на стикер: %v", err)
		log.Printf(msg)

		return "", errors.New(msg)
	}

	return link, nil
}

func (s *client) GetVoiceURL(message MessageDTO) (string, error) {
	link, err := s.getFileLink(*message.VoiceFileID)
	if err != nil {
		msg := fmt.Sprintf("Не удалось получить ссылку на звуковое соообщение: %v", err)
		log.Printf(msg)

		return "", errors.New(msg)
	}

	return link, nil
}

// getFileLink инкапсулирует telegram-токен, поэтому вынесен из сервиса сюда
func (s *client) getFileLink(fileID string) (string, error) {
	file, err := s.client.GetFile(tgbot.FileConfig{
		FileID: fileID,
	})
	if err != nil {
		return "", errors.New("")
	}

	return file.Link(s.config.TgBotToken), nil
}

func (s *client) transformUpdateToMessageDTO(upd tgbot.Update) MessageDTO {
	var stickerFileId, voiceFileId *string

	if upd.Message.Sticker != nil {
		stickerFileId = &upd.Message.Sticker.FileID
	} else {
		stickerFileId = nil
	}

	if upd.Message.Voice != nil {
		voiceFileId = &upd.Message.Voice.FileID
	} else {
		voiceFileId = nil
	}

	return MessageDTO{
		MessageID:     upd.Message.MessageID,
		Date:          upd.Message.Date,
		ChatID:        int(upd.Message.Chat.ID),
		ChatTitle:     upd.Message.Chat.Title,
		StickerFileID: stickerFileId,
		VoiceFileID:   voiceFileId,
		Text:          upd.Message.Text,
		Entities:      s.transformEntities(upd.Message.Entities),
		Photo:         s.transformPhotos(upd.Message.Photo),
	}
}

func (s *client) transformEntities(entities *[]tgbot.MessageEntity) []FormattingDTO {
	if entities == nil {
		return []FormattingDTO{}
	}

	formatting := make([]FormattingDTO, len(*entities))
	for _, value := range *entities {
		formatting = append(formatting, FormattingDTO{
			Type:   value.Type,
			Offset: value.Offset,
			Length: value.Length,
		})
	}

	return formatting
}

func (s *client) transformPhotos(photos *[]tgbot.PhotoSize) []PhotoDTO {
	if photos == nil {
		return []PhotoDTO{}
	}

	customPhotos := make([]PhotoDTO, len(*photos))

	for _, photo := range *photos {
		photoURL, err := s.getFileLink(photo.FileID)
		if err != nil {
			log.Fatal(err)
			continue
		}

		customPhotos = append(customPhotos, PhotoDTO{
			URL:      photoURL,
			Width:    photo.Width,
			Height:   photo.Height,
			FileSize: photo.FileSize,
		})
	}

	return customPhotos
}
