package main

import (
	"fmt"
	"github.com/itxor/tgsite/internal/model"
	"github.com/itxor/tgsite/internal/service"
)

func main() {
	//db, ctx, err := repository.NewMongoDB()
	//if err != nil {
	//	logrus.Fatalf(err.Error())
	//}
	//
	//repos := repository.NewRepository(db, ctx)
	//services, err := service.NewService(repos)
	//if err != nil {
	//	logrus.Fatalf(err.Error())
	//}
	//handlers := handler.NewHandler(services)
	//srv := new(internal.Server)
	//
	//go func () {
	//	if err := services.Telegram.StartUpdatesLoop(); err != nil {
	//		logrus.Fatalf(err.Error())
	//	}
	//}()
	//
	//go func() {
	//	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
	//		logrus.Fatalf("%s", err.Error())
	//	}
	//}()
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	//<- quit
	//
	//if err := db.Disconnect(context.Background()); err != nil {
	//	logrus.Fatalf(err.Error())
	//}
	//
	//if err := srv.Shutdown(context.Background()); err != nil {
	//	logrus.Fatalf(err.Error())
	//}
	var test string
	var i int = 0
	for {
		fmt.Scanf("%s", &test)
		if err := service.PublishNewPost(&model.ChannelPost{
			MessageId: i,
			Date:      0,
			ChatId:    0,
			Content:   model.PostContent{},
			ChatName:  "",
		}); err != nil {
			return
		}

		i = i + 1
	}
}
