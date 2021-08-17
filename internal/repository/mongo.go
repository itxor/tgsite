package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Config struct {
	URL string
}

func NewMongoDB(cfg Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.URL))
	if err != nil {
		msg := fmt.Sprintf("Ошибка при попытке подключения к mongoDB: %s", err.Error())
		log.Printf(msg)

		return nil, errors.New(msg)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		msg := fmt.Sprintf("Ошибка при попытке загрузить контекст: %s", err.Error())
		log.Printf(msg)

		return nil, errors.New(msg)
	}
	//defer client.Disconnect(ctx)

	return client, nil
}
