package dto

type (
	RegisterRequest struct {
		Email          string `json:"email"`
		Username       string `json:"username"`
		Password       string `json:"password"`
		PasswordConfir string `json:"password_confirm"`
	}

	RegisterResponse struct {
		Id int64 `json:"id"`
	}
)
