package post

import (
	"context"
	"database/sql"

	"github.com/jcblastor/api_tweets/internal/model"
)

type PostRepository interface {
	StorePost(ctx context.Context, model *model.Post_Model) (int64, error)
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}
