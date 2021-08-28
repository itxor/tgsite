package mongo

import (
	"context"
	"github.com/itxor/tgsite/internal/model"
	"github.com/itxor/tgsite/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type PostMongo struct {
	db *mongo.Client
	ctx context.Context
}

// NewPostMongo создаёт новый экземпляр PostMongo
func NewPostMongo(ctx context.Context, db *mongo.Client) *PostMongo {
	return &PostMongo{
		ctx: ctx,
		db: db,
	}
}

func (s *PostMongo) Add(post *model.ChannelPost) error {
	chatId := post.ChatId
	if chatId < 0 {
		chatId = chatId * -1
	}

	collection := s.db.
		Database(repository.DatabaseChannels).
		Collection(strconv.Itoa(chatId))

	_, err := collection.InsertOne(s.ctx, post)
	if err != nil {
		return err
	}

	return nil
}


