package memory

import (
	"github.com/itxor/tgsite/internal/model"
)

type MemoryDB struct {
	channels map[string]*model.ChannelPost
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		channels: make(map[string]*model.ChannelPost),
	}
}