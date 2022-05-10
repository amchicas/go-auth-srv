package main

import (
	"net"

	"github.com/amchicas/go-auth-srv/pkg/log"
	"github.com/amchicas/go-auth-srv/pkg/pb"
	"github.com/amchicas/go-auth-srv/pkg/utils"
	"google.golang.org/grpc"

	"github.com/amchicas/go-auth-srv/config"
	"github.com/amchicas/go-auth-srv/internal/domain"
	"github.com/amchicas/go-auth-srv/internal/services"
	"github.com/amchicas/go-auth-srv/internal/storage/mysql"
)

func main() {

	logger := log.New("Auth", "dev")
	c, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed at config" + err.Error())
	}
	repo := newMysql(c.MysqlUrl, logger)
	jwt := utils.JwtWrapper{
		SecretKey:       c.Secret,
		Issuer:          "enretaserno",
		ExpirationHours: 24 * 2,
	}
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {

		logger.Error("Failed at server" + err.Error())

	}

	srv := services.New(repo, logger, jwt)
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("Failed to server :" + err.Error())

	}
}

func newMysql(url string, logger *log.Logger) domain.Repository {
	database, err := mysql.Connect(url)
	if err != nil {

		logger.Error("Failed mysql" + err.Error())
	}

	return mysql.NewRepository(database)

}
