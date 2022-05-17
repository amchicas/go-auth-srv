package adder

import (
	"context"

	"github.com/amchicas/go-auth-srv/internal/domain"
	"github.com/amchicas/go-auth-srv/pkg/log"
)

type Service interface {
	AddUser(ctx context.Context, username, email, password string, role int64) (*domain.User, error)
}

type service struct {
	repo   domain.Repository
	logger *log.Logger
}

func New(repo domain.Repository, logger *log.Logger) Service {

	return &service{
		repo:   repo,
		logger: logger,
	}

}

func (s *service) AddUser(ctx context.Context, username, email, password string, role int64) (*domain.User, error) {

	user := domain.NewUser(username, email, password, role)

	err := s.repo.CreateUser(ctx, user)

	if err != nil {
		return &domain.User{}, err
	}

	return user, nil

}
