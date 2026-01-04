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
)

func (s userService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userId int64) (string, string, int, error) {
	// check user exist
	userExist, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExist == nil {
		return "", "", http.StatusBadRequest, errors.New("user not found")
	}

	// get refresh token by user id
	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, userId, time.Now())
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExists == nil {
		return "", "", http.StatusUnauthorized, errors.New("refresh token was required")
	}

	// check refresh token is match with request body
	if req.RefreshToken != refreshTokenExists.RefreshToken {
		return "", "", http.StatusUnauthorized, errors.New("refresh token not found")
	}

	// generate new token
	token, err := jwt.CreateToken(userId, userExist.UserName, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// delete old refresh token and save new refresh token
	err = s.userRepo.DeleteRefreshTokenByUserId(ctx, userId)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	s.userRepo.StoreRefreshToken(ctx, &model.RefresTokenModel{
		UserId:       userId,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
	})

	return token, refreshToken, http.StatusOK, nil
}
