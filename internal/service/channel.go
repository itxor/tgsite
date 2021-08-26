package service

import "github.com/itxor/tgsite/internal/repository"

type Channel interface {
	List()
	Add(int) error
	Delete()
	Update()
	IsExists(int) bool
}

type ChannelService struct {
	repo repository.Repository
}

func NewChannelService(repo repository.Repository) Channel {
	return &ChannelService{repo: repo}
}

func (s *ChannelService) Add(chatId int) error {
	return nil
}

func (s *ChannelService) List() {

}

func (s *ChannelService) Delete() {

}

func (s *ChannelService) Update() {

}

func (s *ChannelService) IsExists(chatId int) bool {
	return true
}