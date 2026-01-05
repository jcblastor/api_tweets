package post

import (
	"context"
	"net/http"
	"time"

	"github.com/jcblastor/api_tweets/internal/dto"
	"github.com/jcblastor/api_tweets/internal/model"
)

func (s postService) CreatePost(ctx context.Context, req *dto.CreatePostRequest, userId int64) (int64, int, error) {
	// store post
	insertedId, err := s.postRepo.StorePost(ctx, &model.Post_Model{
		UserId:    userId,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	//return
	return insertedId, http.StatusCreated, nil
}
