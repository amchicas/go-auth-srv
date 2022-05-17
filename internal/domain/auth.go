package domain

import (
	"context"
	"time"
)

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int64  `json:"role"`
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
}

func NewUser(username, email, password string, role int64) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
		Created:  time.Now().Unix(),
		Modified: time.Now().Unix(),
	}

}

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
