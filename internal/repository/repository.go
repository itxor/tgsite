package repository

import (
	"context"
	"github.com/itxor/tgsite/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DatabaseChannels = "channels"
)

type ChannelRepository interface {
	IsExist(int) bool
	Add(int) error
}

type PostRepository interface {
	Add(post *model.ChannelPost) error
}

type Repository struct {
	PostRepository
	ChannelRepository
}

func NewRepository(db *mongo.Client, ctx context.Context) Repository {
	return Repository{
		PostRepository: NewPostMongo(ctx, db),
		ChannelRepository: NewChannelMongo(ctx, db),
	}
}
