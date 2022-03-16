package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/itxor/tgsite/internal/domains/channel"
	"github.com/itxor/tgsite/internal/domains/channel/repository"
	"github.com/itxor/tgsite/internal/handlers"
	"github.com/itxor/tgsite/pkg/mongo"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, os.Interrupt)
		for {
			sig := <-ch
			switch sig {
			case os.Interrupt:
				cancel()

				return
			}
		}
	}()

	srv := handlers.NewServer()

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

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if err := srv.Start(ctx); err != nil {
			logrus.Error(err)

			return
		}
	}(wg)

	wg.Wait()
}
