package main

import (
	"context"
	"github.com/itxor/tgsite/internal"
	"github.com/itxor/tgsite/internal/handler"
	"github.com/itxor/tgsite/internal/repository"
	"github.com/itxor/tgsite/internal/repository/mongo"
	"github.com/itxor/tgsite/internal/service"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db, ctx, err := mongo.NewMongoDB()
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	repo := repository.NewRepository(db, ctx)
	handlers := handler.NewHandler(
		service.NewAPIServices(repo),
	)
	srv := new(internal.Server)

	go func() {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("%s", err.Error())
		}
	}()

	parserServices := service.NewTelegramParserService(repo)
	go func () {
		if err := parserServices.StartUpdatesLoop(); err != nil {
			logrus.Fatalf(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	if err := db.Disconnect(context.Background()); err != nil {
		logrus.Fatalf(err.Error())
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf(err.Error())
	}
}
