package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(ctx context.Context) (*mongo.Client, error) {
	cfg, err := NewDatabaseConfig()
	if err != nil {
		msg := fmt.Sprintf("Ошибка при чтении конфига для подключения к mongodb: %s", err.Error())

		return nil, errors.New(msg)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.URL))
	if err != nil {
		msg := fmt.Sprintf("Ошибка при попытке подключения к mongoDB: %s", err.Error())

		return nil, errors.New(msg)
	}

	err = client.Connect(ctx)
	if err != nil {
		msg := fmt.Sprintf("Ошибка при попытке загрузить контекст: %s", err.Error())

		return nil, errors.New(msg)
	}

	return client, nil
}
