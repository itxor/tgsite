package telegram

type TelegramClientInterface interface {
	GetStickerURL(message MessageDTO) (string, error)
	GetVoiceURL(message MessageDTO) (string, error)
	GetUpdateChan() chan MessageDTO
}
