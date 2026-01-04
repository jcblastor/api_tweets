package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/jcblastor/api_tweets/internal/model"
)

func (r *userRepository) GetRefreshToken(ctx context.Context, userId int64, now time.Time) (*model.RefresTokenModel, error) {
	query := `
		SELECT id, user_id, refresh_token, expired_at
		FROM refresh_tokens
		WHERE user_id = ?
		AND expired_at >= ?
	`
	row := r.db.QueryRowContext(ctx, query, userId, now)
	var result model.RefresTokenModel

	err := row.Scan(&result.Id, &result.UserId, &result.RefreshToken, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}
