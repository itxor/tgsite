package repository

import (
	"context"
	"github.com/itxor/tgsite/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DatabaseChannels = "channels"
)

type Channel interface {
	IsExist(int) bool
	Add(int) error
}

type Post interface {
	Add(post *model.ChannelPost) error
}

type Repository struct {
	Post
	Channel
}

func NewRepository(db *mongo.Client, ctx context.Context) Repository {
	return Repository{
		Post:    NewPostMongo(ctx, db),
		Channel: NewChannelMongo(ctx, db),
	}
}
