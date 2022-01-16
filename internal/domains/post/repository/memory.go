package repository

import (
	"bytes"
	"encoding/json"

	"github.com/itxor/tgsite/internal/domains/post"
)

type storage struct {
	db map[int][]byte
}

func NewPostStorage(db map[int][]byte) post.PostRepositoryInterface {
	return &storage{
		db: db,
	}
}

func (r *storage) Add(post post.Post) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(post)
	if err != nil {
		return err
	}

	r.db[post.ChatId] = buf.Bytes()

	return nil
}
