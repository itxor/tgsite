package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type ChannelMongo struct {
	ctx context.Context
	db  *mongo.Client
}

func NewChannelMongo(ctx context.Context, db *mongo.Client) ChannelRepository {
	return &ChannelMongo{
		ctx: ctx,
		db:  db,
	}
}

func (r *ChannelMongo) IsExist(chatId int) bool {
	names, err := r.db.Database(DatabaseChannels).ListCollectionNames(
		context.Background(),
		bson.D{},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	for _, name := range names {
		if name == strconv.Itoa(chatId) {
			return true
		}
	}

	return false
}

func (r *ChannelMongo) Add(chatId int) error {
	return r.db.Database(DatabaseChannels).CreateCollection(r.ctx, strconv.Itoa(chatId))
}
