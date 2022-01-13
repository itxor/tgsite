package telegram

type MessageDTO struct {
	MessageID     int
	Date          int
	ChatID        int
	ChatTitle     string
	StickerFileID *string
	VoiceFileID   *string
	Text          string
	Entities      []FormattingDTO
	Photo         []PhotoDTO
}

type FormattingDTO struct {
	Type   string
	Offset int
	Length int
}

type PhotoDTO struct {
	URL      string
	Width    int
	Height   int
	FileSize int
}
