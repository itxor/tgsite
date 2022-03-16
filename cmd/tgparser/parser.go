package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/itxor/tgsite/internal/domains/post/repository"
	"github.com/itxor/tgsite/internal/domains/post/usecase"
	"github.com/itxor/tgsite/internal/service/nats"
	"github.com/itxor/tgsite/internal/service/telegram"
	"github.com/itxor/tgsite/pkg/mongo"
	nats_client "github.com/itxor/tgsite/pkg/nats"
	tg_client "github.com/itxor/tgsite/pkg/telegram"
	"github.com/sirupsen/logrus"
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

	tgClient, err := tg_client.NewClient("configs/telegram.toml")
	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.NewMongoDB(ctx)
	if err != nil {
		log.Fatal(err)
	}

	postRepo, err := repository.NewPostMongo(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	useCase := usecase.NewUseCaseForTelegramUpdateLoop(
		postRepo,
		tgClient,
		nats.NewNatsTgUpdateService(nats_client.NewClient()),
	)
	tgUpdateLoopService := telegram.NewUpdateLoopService(tgClient, useCase)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if err := tgUpdateLoopService.StartUpdateLoop(ctx); err != nil {
			logrus.Error(err)
			wg.Done()
		}
	}(wg)

	wg.Wait()
}
