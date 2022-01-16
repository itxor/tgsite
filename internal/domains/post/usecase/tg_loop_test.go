package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/itxor/tgsite/internal/domains/post"
	"github.com/itxor/tgsite/internal/domains/post/usecase"
	"github.com/itxor/tgsite/internal/service/nats"
	"github.com/itxor/tgsite/pkg/telegram"
	"github.com/stretchr/testify/assert"
)

func TestBuildNewPostFromMessage(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := post.NewMockPostRepositoryInterface(ctrl)
	tgClient := telegram.NewMockTelegramClientInterface(ctrl)
	natsService := nats.NewMockNatsServiceForTelegramUpdateLoopInterface(ctrl)

	service := usecase.NewUseCaseForTelegramUpdateLoop(
		repo,
		tgClient,
		natsService,
	)

	tgMessageDTO := telegram.MessageDTO{
		MessageID: 1,
		Date:      2,
		ChatID:    3,
		ChatTitle: "chat title test",
		Text:      "new post text",
		Entities:  []telegram.FormattingDTO{},
		Photo:     []telegram.PhotoDTO{},
	}

	post, err := service.BuildNewPostFromMessage(tgMessageDTO)

	assert.Nil(t, err)
	assert.Equal(t, 1, post.MessageId)
	assert.Equal(t, 2, post.Date)
	assert.Equal(t, 3, post.ChatId)
	assert.Equal(t, "chat title test", post.ChatName)
}
