package repository

import "go.mongodb.org/mongo-driver/mongo"

type PostMongo struct {
	db *mongo.Client
}

func NewPostMongo(db *mongo.Client) *PostMongo {
	return &PostMongo{
		db: db,
	}
}

func (s *PostMongo) CreatePost() (int, error) {
	return 1, nil
}


