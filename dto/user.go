package dto

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
	Age      uint64 `json:"age" validate:"required,numeric,min=8"`
}
