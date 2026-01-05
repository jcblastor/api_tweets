package post

import (
	"context"

	"github.com/jcblastor/api_tweets/internal/model"
)

func (r *postRepository) StorePost(ctx context.Context, model *model.Post_Model) (int64, error) {
	query := `
		INSERT INTO posts (user_id, title, content, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(ctx, query, model.UserId, model.Title, model.Content, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
