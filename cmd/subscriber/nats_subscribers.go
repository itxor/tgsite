package main

import (
	"github.com/itxor/tgsite/internal/service"
	"github.com/sirupsen/logrus"
	"sync"
)

func main() {
	nats := service.NewNats()
	df, err := nats.ConnectToMessageBus()
	if err != nil {
		logrus.Fatal(err)
	}
	defer df()

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go nats.SubscribeNewChannelPosts(wg)

	wg.Wait()
}
