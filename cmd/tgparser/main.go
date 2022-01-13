package main

import (
	"github.com/itxor/tgsite/internal/domains/post/repository"
	"github.com/itxor/tgsite/internal/domains/post/usecase"
	"github.com/itxor/tgsite/internal/service/nats"
	"github.com/itxor/tgsite/internal/service/telegram"
	"github.com/itxor/tgsite/pkg/mongo"
	tg_client "github.com/itxor/tgsite/pkg/telegram"
	"log"
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
		usecase.NewUseCaseForTelegramUpdateLoop(postRepo, tgClient, nats.NewNatsTgUpdateService()),
	).StartUpdateLoop(); err != nil {
		log.Fatal(err)
	}
}
