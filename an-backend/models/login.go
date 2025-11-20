package models

type LoginRequestByEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SendCodeRequest struct {
	Email string `json:"email"`
}

type RegisterRequestByEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
