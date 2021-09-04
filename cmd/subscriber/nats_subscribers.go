package main

import (
	"github.com/itxor/tgsite/internal/repository"
	"github.com/itxor/tgsite/internal/service"
	"github.com/sirupsen/logrus"
	"sync"
)

type subscriberDTO struct {
	wg             *sync.WaitGroup
	channelService service.Channel
	postService    service.Post
	nats           *service.Nats
}

var (
	subscriberFuncs = []func(dto *subscriberDTO){
		subscribeToNewChannelPosts,
	}
)

func main() {
	db, ctx, err := repository.NewMongoDB()
	if err != nil {
		logrus.Fatal(err)
	}
	repo := repository.NewRepository(db, ctx)
	channelService := service.NewChannelService(repo)
	postService := service.NewPostService(repo)
	nats := service.NewNats()

	df, err := nats.ConnectToMessageBus()
	if err != nil {
		logrus.Fatal(err)
	}
	defer df()

	wg := new(sync.WaitGroup)

	dto := &subscriberDTO{
		wg:             wg,
		channelService: channelService,
		postService:    postService,
		nats:           nats,
	}

	for _, subFunc := range subscriberFuncs {
		wg.Add(1)
		go subFunc(dto)
	}

	wg.Wait()
}

func subscribeToNewChannelPosts(dto *subscriberDTO) {
	logrus.Error("new post!")
	for post := range dto.nats.GetChannelPostsChan() {
		// todo: логировать в отдельный файл журнала, вывести уровень debug
		logrus.Error(post)
		if !dto.channelService.IsExist(post.ChatId) {
			if err := dto.channelService.Add(post.ChatId); err != nil {
				logrus.Errorf("Ошибка при попытке создать канал: %s", err.Error())
			}
		}

		if err := dto.postService.Add(post); err != nil {
			logrus.Errorf("Ошибка при попытке добавить новый пост: %s", err.Error())
		}
	}

	dto.wg.Done()
}
