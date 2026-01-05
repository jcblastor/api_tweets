package dto

type (
	CreatePostRequest struct {
		Title   string `json:"title" validate:"required"`
		Content string `json:"content" validate:"required"`
	}

	CreatePostResponse struct {
		Id int64 `json:"id"`
	}
)
