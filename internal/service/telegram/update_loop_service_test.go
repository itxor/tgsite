package telegram_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/itxor/tgsite/internal/domains/post"
	tgService "github.com/itxor/tgsite/internal/service/telegram"
	"github.com/itxor/tgsite/pkg/telegram"
	"github.com/stretchr/testify/assert"
)

func TestStartUpdateLoop(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tgMessageDTO := telegram.MessageDTO{
		MessageID: 1,
		Date:      2,
		ChatID:    3,
		ChatTitle: "chat title test",
		Text:      "new post text",
	}
	mockPost := post.Post{
		MessageId: tgMessageDTO.MessageID,
		Date:      tgMessageDTO.Date,
		ChatId:    tgMessageDTO.ChatID,
		ChatName:  tgMessageDTO.ChatTitle,
		Content: post.Content{
			Text: tgMessageDTO.Text,
		},
	}
	msgChan := make(chan telegram.MessageDTO, 1)

	tgClient := telegram.NewMockTelegramClientInterface(ctrl)
	tgClient.
		EXPECT().
		GetUpdateChan().
		DoAndReturn(func() chan telegram.MessageDTO {
			msgChan <- tgMessageDTO
			go func() {
				time.Sleep(1)
				close(msgChan)
			}()

			return msgChan
		})

	postUseCase := post.NewMockPostUseCaseForUpdateTelegramLoopInterface(ctrl)
	postUseCase.
		EXPECT().
		BuildNewPostFromMessage(gomock.Eq(tgMessageDTO)).
		DoAndReturn(func(dto telegram.MessageDTO) (*post.Post, error) {
			return &mockPost, nil
		})
	postUseCase.
		EXPECT().
		DispatchAddPost(gomock.Eq(mockPost)).
		Return(nil)

	updateLoopService := tgService.NewUpdateLoopService(tgClient, postUseCase)
	err := updateLoopService.StartUpdateLoop()

	assert.Nil(t, err)
}
