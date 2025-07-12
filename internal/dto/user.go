package dto

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Passwrod string `json:"password"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Passwrod string `json:"password"`
}
