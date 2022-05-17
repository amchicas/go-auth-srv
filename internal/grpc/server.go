package grpc

import (
	"net"

	"github.com/amchicas/go-auth-srv/internal/adder"
	"github.com/amchicas/go-auth-srv/internal/fetcher"
	"github.com/amchicas/go-auth-srv/pkg/log"
	"github.com/amchicas/go-auth-srv/pkg/pb"
	"github.com/amchicas/go-auth-srv/pkg/utils"
	"google.golang.org/grpc"
)

type Server interface {
	Serve() error
}
type grpcServer struct {
	port   string
	aS     adder.Service
	fS     fetcher.Service
	logger *log.Logger
	jwt    utils.JwtWrapper
}

func NewServer(
	port string,
	aS adder.Service,
	fS fetcher.Service,
	logger *log.Logger,
	jwt utils.JwtWrapper,

) Server {

	return &grpcServer{port: port, aS: aS, fS: fS, logger: logger, jwt: jwt}

}

func (s *grpcServer) Serve() error {

	grpcServer := grpc.NewServer()
	srv := NewHandler(s.aS, s.fS, s.logger, s.jwt)
	pb.RegisterAuthServiceServer(grpcServer, srv)

	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		s.logger.Error("Failed at server" + err.Error())
	}

	return grpcServer.Serve(lis)

}
