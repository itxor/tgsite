package main

import (
	"github.com/itxor/tgsite/internal/domains/channel"
	channel_repo "github.com/itxor/tgsite/internal/domains/channel/repository"
	post_repo "github.com/itxor/tgsite/internal/domains/post/repository"
	post_usecase "github.com/itxor/tgsite/internal/domains/post/usecase"
	"github.com/itxor/tgsite/internal/service/nats"
	"github.com/itxor/tgsite/pkg/mongo"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, client, err := mongo.NewMongoDB()
	if err != nil {
		logrus.Fatal(err)
	}

	postRepo, err := post_repo.NewPostMongo(ctx, client)
	if err != nil {
		logrus.Fatal(err)
	}

	channelRepo, err := channel_repo.NewChannelMongo(ctx, client)
	if err != nil {
		logrus.Fatal(err)
	}

	postUseCase := post_usecase.NewUseCaseForSubscribeNewPosts(postRepo)
	channelUseCase := channel.NewUseCase(channelRepo)

	natsService := nats.NewNatsPostSubscribeService(postUseCase, channelUseCase)

	go func() {
		if err := natsService.SubscribeToNewPostQueue(); err != nil {
			logrus.Fatal(err)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
}
