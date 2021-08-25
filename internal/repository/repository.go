package repository

import (
	"context"
	"github.com/itxor/tgsite/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Database = "channels"
)

type ChannelRepository interface {
	Add1()
}

type PostRepository interface {
	Add(post *model.ChannelPost) error
}

type Repository struct {
	PostRepository
}

func NewRepository(db *mongo.Client, ctx context.Context) Repository {
	return Repository{
		PostRepository: NewPostMongo(ctx, db),
	}
}
