package usecase_test

import (
	"testing"

	"github.com/itxor/tgsite/internal/domains/post"
	"github.com/itxor/tgsite/internal/domains/post/repository"
	"github.com/itxor/tgsite/internal/domains/post/usecase"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	db := make(map[int][]byte)

	service := usecase.NewUseCaseForSubscribeNewPosts(repository.NewPostStorage(db))
	post := post.Post{
		MessageId: 1,
		Date:      1111,
		ChatId:    1111,
		Content: post.Content{
			Text: "Привет",
		},
		ChatName: "122133312",
	}
	err := service.Add(post)
	assert.Nil(t, err)

	_, ok := db[post.ChatId]
	assert.Equal(t, true, ok)
}
