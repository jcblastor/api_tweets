package user

import (
	"context"

	"github.com/jcblastor/api_tweets/internal/config"
	"github.com/jcblastor/api_tweets/internal/dto"
	"github.com/jcblastor/api_tweets/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error)
	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error)
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userId int64) (string, string, int, error)
}

type userService struct {
	cfg      *config.Config
	userRepo user.UserRepository
}

func NewService(cfg *config.Config, userRepo user.UserRepository) UserService {
	return &userService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}
