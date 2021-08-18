package repository

import (
	"context"
	"github.com/itxor/tgsite/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostMongo struct {
	db *mongo.Client
	ctx context.Context
}

// NewPostMongo создаёт новый экземпляр PostMongo
func NewPostMongo(db *mongo.Client, ctx context.Context) *PostMongo {
	return &PostMongo{
		db: db,
		ctx: ctx,
	}
}

// CreatePost сохраняет пост в базу
func (s *PostMongo) CreatePost(post model.ChannelPost) (*mongo.InsertOneResult, error) {
	collection := s.db.Database(Database).Collection(CollectionPosts)
	insertResult, err := collection.InsertOne(s.ctx, post)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}


