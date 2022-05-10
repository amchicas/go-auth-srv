package domain

import "context"

type Auth struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int64  `json:"role"`
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
}

type Repository interface {
	CreateUser(ctx context.Context, auth *Auth) error
	GetByEmail(ctx context.Context, email string) (*Auth, error)
}
