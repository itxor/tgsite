package main

import (
	"github.com/itxor/tgsite/internal/repository"
	"github.com/itxor/tgsite/internal/service"
	"os"
)

func main() {
	db, ctx, err := repository.NewMongoDB()
	if err != nil {
		os.Exit(1)
	}

	s, err := service.NewService(
		repository.NewRepository(db, ctx),
	)
	if err != nil {
		os.Exit(1)
	}

	s.Channel.StartUpdatesLoop()
}