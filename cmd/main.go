package main

import (
	"context"
	"github.com/itxor/tgsite/internal"
	"github.com/itxor/tgsite/internal/handler"
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

	repos := repository.NewRepository(db, ctx)
	services, err := service.NewService(repos)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	handlers := handler.NewHandler(services)
	srv := new(internal.Server)

	go func () {
		if err := services.Telegram.StartUpdatesLoop(); err != nil {
			logrus.Fatalf(err.Error())
		}
	}()

	go func() {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("%s", err.Error())
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
