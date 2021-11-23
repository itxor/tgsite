package post

import (
	"github.com/itxor/tgsite/internal/service/nats"
	"github.com/itxor/tgsite/pkg/telegram"
	"log"
)

type useCase struct {
	repository  PostRepositoryInterface
	tgClient    telegram.TelegramClientInterface
	natsService nats.NatsServiceInterface
}

func NewUseCase(
	repository PostRepositoryInterface,
) PostUserCaseInterface {
	tg, err := telegram.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	return &useCase{
		repository:  repository,
		tgClient:    tg,
		natsService: nats.NewService(),
	}
}

func (s *useCase) Add(post Post) error {
	return s.repository.Add(post)
}

func (s *useCase) BuildNewPostFromMessage(dto telegram.MessageDTO) (*Post, error) {
	newPost := new(Post)
	var err error

	newPost.MessageId = dto.MessageID
	newPost.Date = dto.Date
	newPost.ChatId = dto.ChatID
	newPost.ChatName = dto.ChatTitle

	if "" != dto.Text {
		newPost.Content.Text = dto.Text
	}

	if dto.Entities != nil {
		newPost.Content.Options = make([]Formatting, 0, len(dto.Entities))
		for _, formatOption := range dto.Entities {
			newPost.Content.Options = append(newPost.Content.Options, Formatting{
				FormattingType: formatOption.Type,
				Offset:         formatOption.Offset,
				Length:         formatOption.Length,
			})
		}
	}

	if dto.StickerFileID != "" {
		newPost.Content.StickerURL, err = s.tgClient.GetStickerURL(dto)
		if err != nil {
			return nil, err
		}
	}

	if dto.VoiceFileID != "" {
		newPost.Content.VoiceURL, err = s.tgClient.GetVoiceURL(dto)
		if err != nil {
			return nil, err
		}
	}

	if dto.Photo != nil {
		newPost.Content.Photo = make([]Photo, 0, len(dto.Photo))
		for _, photoDTO := range dto.Photo {
			newPost.Content.Photo = append(newPost.Content.Photo, Photo{
				URL:      photoDTO.URL,
				Width:    photoDTO.Width,
				Height:   photoDTO.Height,
				FileSize: photoDTO.FileSize,
			})
		}
	}

	return newPost, nil
}

func (s *useCase) DispatchAddPost(post Post) error {
	return s.natsService.Dispatch(nats.AddNewPostQueue, post)
}
