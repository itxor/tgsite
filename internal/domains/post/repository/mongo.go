package repository

import (
	"context"
	"github.com/itxor/tgsite/internal/domains/channel"
	"github.com/itxor/tgsite/internal/domains/post"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type postMongo struct {
	client *mongo.Client
	ctx    context.Context
}

func NewPostMongo(ctx context.Context, client *mongo.Client) (post.PostRepositoryInterface, error) {
	return &postMongo{
		ctx:    ctx,
		client: client,
	}, nil
}

func (s *postMongo) Add(post post.Post) error {
	chatId := post.ChatId
	if chatId < 0 {
		chatId = chatId * -1
	}

	collection := s.client.
		Database(channel.DatabaseChannels).
		Collection(strconv.Itoa(chatId))

	_, err := collection.InsertOne(s.ctx, post)
	if err != nil {
		return err
	}

	return nil
}
