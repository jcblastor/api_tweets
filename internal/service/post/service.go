package post

import (
	"context"

	"github.com/jcblastor/api_tweets/internal/config"
	"github.com/jcblastor/api_tweets/internal/dto"
	"github.com/jcblastor/api_tweets/internal/repository/post"
)

type PostService interface {
	CreatePost(ctx context.Context, req *dto.CreatePostRequest, userId int64) (int64, int, error)
}

type postService struct {
	cfg      *config.Config
	postRepo post.PostRepository
}

func NewPostService(cfg *config.Config, postRepo post.PostRepository) PostService {
	return &postService{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
