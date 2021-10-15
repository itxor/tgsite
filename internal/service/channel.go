package service

import "github.com/itxor/tgsite/internal/repository"

type Channel interface {
	Add(int) error
	IsExist(int) bool
}

type ChannelService struct {
	repo repository.Repository
}

func NewChannelService(repo repository.Repository) *ChannelService {
	return &ChannelService{repo: repo}
}

func (s *ChannelService) Add(chatId int) error {
	return s.repo.ChannelRepository.Add(chatId)
}

func (s *ChannelService) IsExist(chatId int) bool {
	return s.repo.ChannelRepository.IsExist(chatId)
}