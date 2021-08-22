package repository

import (
	"context"
	"github.com/itxor/tgsite/internal/model"
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

// CreatePost сохраняет пост в базу
func (s *PostMongo) Add(post model.ChannelPost) (*mongo.InsertOneResult, error) {
	chatId := post.ChatId
	if chatId < 0 {
		chatId = chatId * -1
	}

	collection := s.db.
		Database(Database).
		Collection(strconv.Itoa(chatId))

	insertResult, err := collection.InsertOne(s.ctx, post)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}


