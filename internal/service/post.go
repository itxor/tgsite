package service

import (
	"github.com/itxor/tgsite/internal/model"
	"github.com/itxor/tgsite/internal/repository"
)

type Post interface {
	Add(*model.ChannelPost) error
	List()
	Delete()
	Update()
}

type PostService struct {
	repo repository.Repository
}

func NewPostService(repo repository.Repository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Add(post *model.ChannelPost) error {
	return s.repo.Post.Add(post)
}

func (s *PostService) List() {

}

func (s *PostService) Delete() {

}

func (s *PostService) Update() {

}