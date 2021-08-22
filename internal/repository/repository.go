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

type Post interface {
	Add(post model.ChannelPost) (*mongo.InsertOneResult, error)
}

type Repository struct {
	Post
}

func NewRepository(db *mongo.Client, ctx context.Context) Repository {
	return Repository{
		Post: NewPostMongo(ctx, db),
	}
}
