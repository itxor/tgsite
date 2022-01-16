package repository

import (
	"errors"

	"github.com/itxor/tgsite/internal/domains/channel"
)

type channelMemory struct {
	db map[int]string
}

func NewChannelMemory(db map[int]string) channel.ChannelRepositoryInterface {
	return &channelMemory{
		db: db,
	}
}

func (r *channelMemory) Add(channelId int) error {
	_, ok := r.db[channelId]
	if ok {
		return errors.New("Ключ уже существует!")
	}

	r.db[channelId] = ""

	return nil
}

func (r *channelMemory) IsExist(channelId int) bool {
	if _, ok := r.db[channelId]; ok {
		return true
	}

	return false
}
