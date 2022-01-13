package post

import "github.com/itxor/tgsite/pkg/telegram"

// Post определяет пост, отправленный в канал
type Post struct {
	MessageId int     `bson:"message_id"`
	Date      int     `bson:"date"`
	ChatId    int     `bson:"chat_id"`
	Content   Content `bson:"post_contents"`
	ChatName  string  `bson:"chat_name"`
}

type PostUseCaseForUpdateTelegramLoopInterface interface {
	BuildNewPostFromMessage(dto telegram.MessageDTO) (*Post, error)
	DispatchAddPost(post Post) error
}

type PostUseCaseForSubscribeNewPostsInterface interface {
	Add(post Post) error
}

type PostRepositoryInterface interface {
	Add(post Post) error
}

// Formatting определяет единицу форматированния переданного текста
type Formatting struct {
	FormattingType string `bson:"formatting_type"`
	Offset         int    `bson:"offset"`
	Length         int    `bson:"length"`
}

// Photo определяет массив с разными размерами изображения
type Photo struct {
	URL      string `bson:"url"`
	Width    int    `bson:"width"`
	Height   int    `bson:"height"`
	FileSize int    `bson:"file_size"`
}

// Content определяет контент поста
type Content struct {
	Text       string       `bson:"text"`
	Options    []Formatting `bson:"options"`
	StickerURL string       `bson:"sticker_url"`
	VoiceURL   string       `bson:"voice_url"`
	Photo      []Photo      `bson:"photo"`
}
