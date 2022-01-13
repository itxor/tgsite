package usecase

import (
	"github.com/itxor/tgsite/internal/domains/post"
	"github.com/itxor/tgsite/internal/service/nats"
	"github.com/itxor/tgsite/pkg/telegram"
)

type useCaseForTelegramUpdateLoop struct {
	repository  post.PostRepositoryInterface
	tgClient    telegram.TelegramClientInterface
	natsService nats.NatsServiceForTelegramUpdateLoopInterface
}

func NewUseCaseForTelegramUpdateLoop(
	repository post.PostRepositoryInterface,
	tgClient telegram.TelegramClientInterface,
	natsService nats.NatsServiceForTelegramUpdateLoopInterface,
) post.PostUseCaseForUpdateTelegramLoopInterface {
	return &useCaseForTelegramUpdateLoop{
		repository:  repository,
		tgClient:    tgClient,
		natsService: natsService,
	}
}

func (s *useCaseForTelegramUpdateLoop) BuildNewPostFromMessage(dto telegram.MessageDTO) (*post.Post, error) {
	newPost := new(post.Post)
	var err error

	newPost.MessageId = dto.MessageID
	newPost.Date = dto.Date
	newPost.ChatId = dto.ChatID
	newPost.ChatName = dto.ChatTitle

	if "" != dto.Text {
		newPost.Content.Text = dto.Text
	}

	if dto.Entities != nil {
		newPost.Content.Options = make([]post.Formatting, 0, len(dto.Entities))
		for _, formatOption := range dto.Entities {
			newPost.Content.Options = append(newPost.Content.Options, post.Formatting{
				FormattingType: formatOption.Type,
				Offset:         formatOption.Offset,
				Length:         formatOption.Length,
			})
		}
	}

	if dto.StickerFileID != nil {
		newPost.Content.StickerURL, err = s.tgClient.GetStickerURL(dto)
		if err != nil {
			return nil, err
		}
	}

	if dto.VoiceFileID != nil {
		newPost.Content.VoiceURL, err = s.tgClient.GetVoiceURL(dto)
		if err != nil {
			return nil, err
		}
	}

	if dto.Photo != nil {
		newPost.Content.Photo = make([]post.Photo, 0, len(dto.Photo))
		for _, photoDTO := range dto.Photo {
			newPost.Content.Photo = append(newPost.Content.Photo, post.Photo{
				URL:      photoDTO.URL,
				Width:    photoDTO.Width,
				Height:   photoDTO.Height,
				FileSize: photoDTO.FileSize,
			})
		}
	}

	return newPost, nil
}

func (s *useCaseForTelegramUpdateLoop) DispatchAddPost(post post.Post) error {
	return s.natsService.Dispatch(nats.NewPostsSubject, post)
}
