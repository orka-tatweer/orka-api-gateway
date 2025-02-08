package auth

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
