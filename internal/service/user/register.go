package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/jcblastor/api_tweets/internal/dto"
	"github.com/jcblastor/api_tweets/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error) {
	// check user already exists
	userExist, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, req.Username)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	if userExist != nil {
		return 0, http.StatusBadRequest, errors.New("user already exist")
	}

	// hash password
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	// create user
	userModel := &model.UserModel{
		Email:     req.Email,
		UserName:  req.Username,
		Password:  string(passHash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userId, err := s.userRepo.CreateUser(ctx, userModel)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return userId, http.StatusCreated, nil
}
