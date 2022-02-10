package main

import (
	"context"

	"github.com/itxor/tgsite/internal/domains/channel"
	"github.com/itxor/tgsite/internal/domains/channel/repository"
	"github.com/itxor/tgsite/internal/handlers"
	"github.com/itxor/tgsite/pkg/mongo"
	"github.com/sirupsen/logrus"
)

func main() {
	srv := handlers.NewServer()

	ctx := context.Background()
	client, err := mongo.NewMongoDB(ctx)
	if err != nil {
		logrus.Error(err)
	}

	channelStorage, err := repository.NewChannelMongo(ctx, client)
	if err != nil {
		logrus.Error(err)
	}

	channelUseCase := channel.NewUseCase(channelStorage)

	srv.RegisterHandlers(
		channel.NewChannelHandler(channelUseCase),
	)
}
