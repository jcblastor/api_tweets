package user

import (
	"context"
	"errors"
)

func (r *userRepository) DeleteRefreshTokenByUserId(ctx context.Context, userId int64) error {
	query := `
		DELETE FROM refresh_tokens
		WHERE user_id = ?
	`
	result, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("nothing to delete")
	}

	return nil
}
