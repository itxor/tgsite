package main

import (
	"github.com/itxor/tgsite/internal/domains/post/repository"
	"github.com/itxor/tgsite/internal/service/telegram"
	"github.com/itxor/tgsite/pkg/mongo"
	tg_client "github.com/itxor/tgsite/pkg/telegram"
	"log"
)

func main() {
	tgClient, err := tg_client.NewClient("")
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

	tgUpdateService := telegram.NewUpdateLoopService(tgClient, postRepo)
	go func() {
		if err := tgUpdateService.StartUpdateLoop(); err != nil {
			log.Fatal(err)
		}
	}()
}
