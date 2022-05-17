package fetcher

import (
	"context"

	"github.com/amchicas/go-auth-srv/internal/domain"
	"github.com/amchicas/go-auth-srv/pkg/log"
)

type Service interface {
	FetchUserByEmail(ctx context.Context, email string) (*domain.User, error)
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

func (s *service) FetchUserByEmail(ctx context.Context, email string) (*domain.User, error) {

	user, err := s.repo.GetByEmail(ctx, email)

	if err != nil {
		return &domain.User{}, err
	}

	return user, nil

}
