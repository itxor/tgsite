package repository

import "go.mongodb.org/mongo-driver/mongo"

type Post interface {
	CreatePost () (int, error)
}

type Repository struct {
	Post
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Post: NewPostMongo(db),
	}
}
