package service

import (
	"github.com/itxor/tgsite/internal/model"
	"github.com/itxor/tgsite/internal/repository"
)

type PostService struct {
	repo repository.Repository
}

func NewPostService(repo repository.Repository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Add(post *model.ChannelPost) error {
	return s.repo.PostRepository.Add(post)
}

func (s *PostService) List() {

}

func (s *PostService) Delete() {

}

func (s *PostService) Update() {

}