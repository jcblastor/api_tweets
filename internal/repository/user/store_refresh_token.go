package user

import (
	"context"

	"github.com/jcblastor/api_tweets/internal/model"
)

func (r *userRepository) StoreRefreshToken(ctx context.Context, model *model.RefresTokenModel) error {
	query := `
		INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := r.db.ExecContext(ctx, query, model.UserId, model.RefreshToken, model.ExpiredAt, model.CreatedAt, model.UpdatedAt)

	return err
}
