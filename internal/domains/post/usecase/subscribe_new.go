package usecase

import (
	"github.com/itxor/tgsite/internal/domains/post"
)

type useCaseForSubscribeNewPosts struct {
	postRepo post.PostRepositoryInterface
}

func NewUseCaseForSubscribeNewPosts(
	postRepo post.PostRepositoryInterface,
) post.PostUseCaseForSubscribeNewPostsInterface {
	return &useCaseForSubscribeNewPosts{
		postRepo: postRepo,
	}
}

func (s *useCaseForSubscribeNewPosts) Add(post post.Post) error {
	return s.postRepo.Add(post)
}
