package channel_test

import (
	"testing"

	"github.com/itxor/tgsite/internal/domains/channel"
	"github.com/itxor/tgsite/internal/domains/channel/repository"
	"github.com/stretchr/testify/assert"
)

var db map[int]string

func TestAdd(t *testing.T) {
	t.Parallel()
	db := make(map[int]string)

	service := channel.NewUseCase(repository.NewChannelMemory(db))
	err := service.Add(1)
	assert.Nil(t, err)

	value, ok := db[1]
	assert.Equal(t, true, ok)

	assert.Equal(t, "", value)
}

func TestIsExist(t *testing.T) {
	t.Parallel()
	db := make(map[int]string)

	service := channel.NewUseCase(repository.NewChannelMemory(db))
	service.Add(1)
	assert.Equal(t, true, service.IsExist(1))

	assert.Equal(t, false, service.IsExist(2))
}
