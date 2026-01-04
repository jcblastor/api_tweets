package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/jcblastor/api_tweets/internal/dto"
	"github.com/jcblastor/api_tweets/internal/model"
	"github.com/jcblastor/api_tweets/pkg/jwt"
	"github.com/jcblastor/api_tweets/pkg/refreshtoken"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {
	// check user exists
	userExist, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExist == nil {
		return "", "", http.StatusNotFound, errors.New("wrong email or password")
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(req.Password))
	if err != nil {
		return "", "", http.StatusNotFound, errors.New("wrong email or password")
	}

	// generate access token
	token, err := jwt.CreateToken(userExist.Id, userExist.UserName, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// get refresh token if exist
	refreshTokenExist, err := s.userRepo.GetRefreshToken(ctx, userExist.Id, time.Now())
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExist != nil {
		return token, refreshTokenExist.RefreshToken, http.StatusOK, nil
	}

	// generate and store refresh token
	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepo.StoreRefreshToken(ctx, &model.RefresTokenModel{
		UserId:       userExist.Id,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
	})
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	// return
	return token, refreshToken, http.StatusOK, nil
}
