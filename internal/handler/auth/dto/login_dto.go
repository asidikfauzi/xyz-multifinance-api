package dto

type (
	LoginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Users RegisterResponse `json:"users"`
		Token string           `json:"token"`
	}
)
