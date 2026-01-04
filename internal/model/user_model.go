package model

import "time"

type (
	UserModel struct {
		Id        int64
		Email     string
		UserName  string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	RefresTokenModel struct {
		Id           int64
		UserId       int64
		RefreshToken string
		ExpiredAt    time.Time
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)
