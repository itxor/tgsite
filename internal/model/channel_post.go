package model

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

// PostContent определяет контент поста
type PostContent struct {
	Text       string       `bson:"text"`
	Options    []Formatting `bson:"options"`
	StickerURL string       `bson:"sticker_url"`
	VoiceURL   string       `bson:"voice_url"`
	Photo     []Photo     `bson:"photo"`
}

// ChannelPost определяет пост, отправленный в канал
type ChannelPost struct {
	Id        int         `bson:"id"`
	MessageId int         `bson:"message_id"`
	Date      int         `bson:"date"`
	ChatId    int         `bson:"chat_id"`
	Content   PostContent `bson:"post_contents"`
}
