package repository

import (
	"context"
	"strconv"

	"github.com/itxor/tgsite/internal/domains/channel"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type channelMongo struct {
	ctx    context.Context
	client *mongo.Client
}

func NewChannelMongo(ctx context.Context, client *mongo.Client) (channel.ChannelRepositoryInterface, error) {
	return &channelMongo{
		ctx:    ctx,
		client: client,
	}, nil
}

func (r *channelMongo) IsExist(chatId int) bool {
	names, err := r.client.Database(channel.DatabaseChannels).ListCollectionNames(
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

func (r *channelMongo) Add(chatId int) error {
	return r.client.Database(channel.DatabaseChannels).CreateCollection(r.ctx, strconv.Itoa(chatId))
}
