package main

import (
	"context"
	"github.com/itxor/tgsite/internal/domains/channel"
	channel_repo "github.com/itxor/tgsite/internal/domains/channel/repository"
	post_repo "github.com/itxor/tgsite/internal/domains/post/repository"
	post_usecase "github.com/itxor/tgsite/internal/domains/post/usecase"
	"github.com/itxor/tgsite/internal/service/nats"
	"github.com/itxor/tgsite/pkg/mongo"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigCh := make(chan os.Signal)
		signal.Notify(sigCh, os.Interrupt)
		for {
			sig := <-sigCh
			switch sig {
			case os.Interrupt:
				cancel()

				return
			}
		}
	}()

	client, err := mongo.NewMongoDB(ctx)
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

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if err := natsService.SubscribeToNewPostQueue(ctx); err != nil {
			logrus.Fatal(err)
		}
	}(wg)

	wg.Wait()
}
