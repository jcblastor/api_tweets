package user

import (
	"context"

	"github.com/jcblastor/api_tweets/internal/model"
)

func (r *userRepository) CreateUser(ctx context.Context, m *model.UserModel) (int64, error) {
	query := `
		INSERT INTO users (email, username, password, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(ctx, query, m.Email, m.UserName, m.Password, m.CreatedAt, m.UpdatedAt)
	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}
