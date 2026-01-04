package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/jcblastor/api_tweets/internal/model"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error)
	CreateUser(ctx context.Context, model *model.UserModel) (int64, error)
	GetRefreshToken(ctx context.Context, userId int64, now time.Time) (*model.RefresTokenModel, error)
	StoreRefreshToken(ctx context.Context, model *model.RefresTokenModel) error
	GetUserById(ctx context.Context, userId int64) (*model.UserModel, error)
	DeleteRefreshTokenByUserId(ctx context.Context, userId int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
