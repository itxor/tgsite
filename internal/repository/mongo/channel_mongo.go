package mongo

import (
	"context"
	"github.com/itxor/tgsite/internal/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type ChannelMongo struct {
	ctx context.Context
	db *mongo.Client
}

func NewChannelMongo(ctx context.Context, db *mongo.Client) repository.Channel {
	return &ChannelMongo{
		ctx: ctx,
		db: db,
	}
}

func (r *ChannelMongo) IsExist(chatId int) bool {
	names, err := r.db.Database(repository.DatabaseChannels).ListCollectionNames(
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
	return r.db.Database(repository.DatabaseChannels).CreateCollection(r.ctx, strconv.Itoa(chatId))
}
