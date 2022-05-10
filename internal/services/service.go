package services

import (
	"context"
	"net/http"
	"time"

	"github.com/amchicas/go-auth-srv/internal/domain"
	"github.com/amchicas/go-auth-srv/pkg/log"
	"github.com/amchicas/go-auth-srv/pkg/pb"
	"github.com/amchicas/go-auth-srv/pkg/utils"
)

type Service struct {
	repo   domain.Repository
	logger *log.Logger
	jwt    utils.JwtWrapper
}

func New(repo domain.Repository, logger *log.Logger, jwt utils.JwtWrapper) *Service {

	return &Service{
		repo:   repo,
		logger: logger,
		jwt:    jwt,
	}

}

func (s *Service) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	var user domain.Auth
	_, err := s.repo.GetByEmail(ctx, req.Email)
	if err == nil {
		return &pb.RegisterResp{
			Status: http.StatusConflict,
			Error:  "Email already exist",
		}, nil
	}
	user.Username = req.Username
	user.Email = req.Email
	user.Role = req.Role
	user.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return &pb.RegisterResp{
			Status: http.StatusInternalServerError,
			Error:  "Couldn't hash password",
		}, nil

	}
	user.Created = time.Now().Unix()
	user.Modified = time.Now().Unix()
	err = s.repo.CreateUser(ctx, &user)
	if err != nil {
		s.logger.Error(err.Error())
		return &pb.RegisterResp{
			Status: http.StatusConflict,
			Error:  "Error created",
		}, nil

	}
	return &pb.RegisterResp{
		Status: http.StatusCreated,
		Error:  "Created",
	}, nil

}

func (s *Service) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return &pb.LoginResp{
			Status: http.StatusNotFound,
			Error:  "Bad credentials email",
		}, err
	}
	match := utils.CheckPasswordHash(req.Password, user.Password)
	if !match {

		return &pb.LoginResp{

			Status: http.StatusNotFound,
			Error:  "Bad credentialspassword ",
		}, err

	}
	token, _ := s.jwt.Sign(user)

	return &pb.LoginResp{

		Status: http.StatusNotFound,
		Token:  token,
	}, nil

}
func (s *Service) Validate(ctx context.Context, req *pb.ValidateReq) (*pb.ValidateResp, error) {
	claim, err := s.jwt.Validate(req.Token)
	if err != nil {

		return &pb.ValidateResp{

			Status: http.StatusNotFound,
			Error:  "Bad token",
		}, nil

	}

	return &pb.ValidateResp{

		Status: http.StatusNotFound,
		UserId: claim.Id,
	}, nil
}
