package user

import (
	"context"
	"database/sql"

	"github.com/jcblastor/api_tweets/internal/model"
)

func (r *userRepository) GetUserById(ctx context.Context, userId int64) (*model.UserModel, error) {
	query := `
		SELECT id, username, email, created_at, updated_at
		FROM users
		WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, query, userId)
	var result model.UserModel

	err := row.Scan(&result.Id, &result.UserName, &result.Email, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}
