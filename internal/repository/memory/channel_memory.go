package memory

import (
	"github.com/itxor/tgsite/internal/model"
	"strconv"
)

type ChannelMemory struct {
	memory *MemoryDB
}

func NewChannelMemory(memory *MemoryDB) *ChannelMemory {
	return &ChannelMemory{
		memory: memory,
	}
}

func (m *ChannelMemory) IsExist(chatId int) bool {
	for key, _ := range m.memory.channels {
		if key == strconv.Itoa(chatId) {
			return true
		}
	}

	return false
}

func (m *ChannelMemory) Add(chatId int) error {
	m.memory.channels[strconv.Itoa(chatId)] = &model.ChannelPost{}

	return nil
}