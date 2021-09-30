package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/itxor/tgsite/internal"
	"github.com/itxor/tgsite/internal/handler"
	"github.com/itxor/tgsite/internal/repository"
	"github.com/itxor/tgsite/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	db, ctx, err := repository.NewMongoDB()
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

	go func() {
		if err := service.NewTelegramParserService(repo).StartUpdatesLoop(); err != nil {
			logrus.Fatalf(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := db.Disconnect(context.Background()); err != nil {
		logrus.Fatalf(err.Error())
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf(err.Error())
	}
}
