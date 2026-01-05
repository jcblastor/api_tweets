package model

import "time"

type (
	Post_Model struct {
		Id        int64
		UserId    int64
		Title     string
		Content   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
	}
)
