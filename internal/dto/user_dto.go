package dto

type (
	RegisterRequest struct {
		Email          string `json:"email" validate:"required,email"`
		Username       string `json:"username" validate:"required,min=3"`
		Password       string `json:"password" validate:"required"`
		PasswordConfir string `json:"password_confirm" validate:"required,eqfield=Password"`
	}

	RegisterResponse struct {
		Id int64 `json:"id"`
	}
)
