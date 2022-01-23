package main

import (
	"log"

	"github.com/itxor/tgsite/internal/domains/post/repository"
	"github.com/itxor/tgsite/internal/domains/post/usecase"
	"github.com/itxor/tgsite/internal/service/nats"
	"github.com/itxor/tgsite/internal/service/telegram"
	"github.com/itxor/tgsite/pkg/mongo"
	nats_client "github.com/itxor/tgsite/pkg/nats"
	tg_client "github.com/itxor/tgsite/pkg/telegram"
)

func main() {
	tgClient, err := tg_client.NewClient("configs/telegram.toml")
	if err != nil {
		log.Fatal(err)
	}

	ctx, client, err := mongo.NewMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	postRepo, err := repository.NewPostMongo(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	if err := telegram.NewUpdateLoopService(
		tgClient,
		usecase.NewUseCaseForTelegramUpdateLoop(postRepo, tgClient, nats.NewNatsTgUpdateService(nats_client.NewClient())),
	).StartUpdateLoop(); err != nil {
		log.Fatal(err)
	}
}
