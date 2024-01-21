package dto

type RegisterUserDTO struct {
	Username  string `json:"username" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

type CreateUserDTO struct {
	Username       string
	Name           string
	HashedPassword string
	IsActive       bool
	IsAdmin        bool
}

type AuthUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int    `json:"expires"`
	ExpiresAt    int    `json:"expires_at"`
}

type UserRetrieveDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}
