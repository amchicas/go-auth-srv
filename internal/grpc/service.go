package grpc

import (
	"context"
	"net/http"

	"github.com/amchicas/go-auth-srv/internal/adder"
	"github.com/amchicas/go-auth-srv/internal/fetcher"
	"github.com/amchicas/go-auth-srv/pkg/log"
	"github.com/amchicas/go-auth-srv/pkg/pb"
	"github.com/amchicas/go-auth-srv/pkg/utils"
)

type authHandler struct {
	aS     adder.Service
	fS     fetcher.Service
	logger *log.Logger
	jwt    utils.JwtWrapper
}

func NewHandler(adderService adder.Service, fetcherService fetcher.Service, logger *log.Logger, jwt utils.JwtWrapper) pb.AuthServiceServer {

	return &authHandler{
		aS:     adderService,
		fS:     fetcherService,
		logger: logger,
		jwt:    jwt,
	}

}

func (s *authHandler) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	_, err := s.fS.FetchUserByEmail(ctx, req.Email)
	if err == nil {
		return &pb.RegisterResp{
			Status: http.StatusConflict,
			Error:  "Email already exist",
		}, nil
	}
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return &pb.RegisterResp{
			Status: http.StatusInternalServerError,
			Error:  "Couldn't hash password",
		}, nil

	}
	_, err = s.aS.AddUser(ctx, req.Username, req.Email, password, req.Role)
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

func (s *authHandler) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	user, err := s.fS.FetchUserByEmail(ctx, req.Email)
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
func (s *authHandler) Validate(ctx context.Context, req *pb.ValidateReq) (*pb.ValidateResp, error) {
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
