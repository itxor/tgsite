package main

import (
	"fmt"
	"github.com/itxor/tgsite/internal/repository"
	"github.com/sirupsen/logrus"
	"sync"

	"github.com/itxor/tgsite/internal/service"
)

type dto struct {
	wg       *sync.WaitGroup
	nats     *service.Nats
	services *service.Service
}

var (
	subscribeFuncs = []func(*dto){
		subscribeNewChannelPosts,
	}
)

func main() {
	nats := service.NewNats()
	df, err := nats.ConnectToMessageBus()
	if err != nil {
		logrus.Fatal(err)
	}
	defer df()

	db, ctx, err := repository.NewMongoDB()
	if err != nil {
		logrus.Fatal(err)
	}

	dto := &dto{
		wg:       new(sync.WaitGroup),
		nats:     nats,
		services: service.NewService(repository.NewRepository(db, ctx)),
	}

	for _, callFunc := range subscribeFuncs {
		dto.wg.Add(1)
		go callFunc(dto)
	}

	dto.wg.Wait()
}

// subscribeNewChannelPosts читаем сообщения из канала с новыми постами, сохраняет посты в базу
func subscribeNewChannelPosts(dto *dto) {
	fmt.Print("a")
	for post := range dto.nats.GetChannelForNewChannelPosts() {
		// todo: логировать в отдельный файл журнала
		logrus.Debug(post)
		if !dto.services.Channel.IsExist(post.ChatId) {
			if err := dto.services.Channel.Add(post.ChatId); err != nil {
				logrus.Errorf("Ошибка при попытке создать канал: %s", err.Error())
			}
		}

		if err := dto.services.Post.Add(post); err != nil {
			logrus.Errorf("Ошибка при попытке добавить новый пост: %s", err.Error())
		}
	}

	dto.wg.Done()
}
