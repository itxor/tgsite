package main

import (
	"github.com/itxor/tgsite/internal/model"
	"github.com/itxor/tgsite/internal/service"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

func main() {
	conn, deferFunc, err := service.Connect()
	if err != nil {
		return
	}
	defer deferFunc()

	ec, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go subscribeNewChannelPosts(ec, wg)

	wg.Wait()
}

func subscribeNewChannelPosts(ec *nats.EncodedConn, wg *sync.WaitGroup) {
	ch := make(chan *model.ChannelPost)
	_, err := ec.BindRecvChan(service.TopicNewPost, ch)
	if err != nil {
		return
	}

	for post := range ch {
		log.Printf("recieve: %#v", post)
	}

	wg.Done()
}
