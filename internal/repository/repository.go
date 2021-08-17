package repository

type Post interface {
	CreatePost () (int, error)
}

type Repository struct {
	Post
}

func NewRepository() *Repository {
	return &Repository{
	}
}
