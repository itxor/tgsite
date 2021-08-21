package main

import (
	"context"
	"github.com/itxor/tgsite/internal/repository"
	"github.com/itxor/tgsite/internal/service"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db, ctx, err := repository.NewMongoDB()
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	s, err := service.NewService(
		repository.NewRepository(db, ctx),
	)
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	go func () {
		if err := s.Channel.StartUpdatesLoop(); err != nil {
			logrus.Fatalf(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	if err := db.Disconnect(context.Background()); err != nil {
		logrus.Fatalf(err.Error())
	}
}
