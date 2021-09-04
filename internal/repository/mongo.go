package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/itxor/tgsite/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() (*mongo.Client, context.Context, error) {
	cfg, err := config.NewDatabaseConfig()
	if err != nil {
		msg := fmt.Sprintf("Ошибка при чтении конфига для подключения к mongodb: %s", err.Error())

		return nil, nil, errors.New(msg)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.URL))
	if err != nil {
		msg := fmt.Sprintf("Ошибка при попытке подключения к mongoDB: %s", err.Error())

		return nil, nil, errors.New(msg)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		msg := fmt.Sprintf("Ошибка при попытке загрузить контекст: %s", err.Error())

		return nil, nil, errors.New(msg)
	}

	return client, ctx, nil
}
