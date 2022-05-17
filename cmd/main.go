package main

import (
	"context"

	"github.com/amchicas/go-auth-srv/config"
	"github.com/amchicas/go-auth-srv/internal/adder"
	"github.com/amchicas/go-auth-srv/internal/domain"
	"github.com/amchicas/go-auth-srv/internal/fetcher"
	"github.com/amchicas/go-auth-srv/internal/grpc"
	"github.com/amchicas/go-auth-srv/internal/storage/mysql"
	"github.com/amchicas/go-auth-srv/pkg/log"
	"github.com/amchicas/go-auth-srv/pkg/utils"
	"golang.org/x/sync/errgroup"
)

func main() {

	logger := log.New("Auth", "dev")
	c, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed at config" + err.Error())
	}
	repo := newMysql(c.MysqlUrl, logger)

	adderService := adder.New(repo, logger)
	fetcherService := fetcher.New(repo, logger)

	jwt := utils.JwtWrapper{
		SecretKey:       c.Secret,
		Issuer:          "enretaservo.com",
		ExpirationHours: 24 * 2,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		srv := grpc.NewServer(c.Port, adderService, fetcherService, logger, jwt)
		return srv.Serve()
	})
	logger.Fatal(g.Wait().Error())
}

func newMysql(url string, logger *log.Logger) domain.Repository {
	database, err := mysql.Connect(url)
	if err != nil {

		logger.Error("Failed mysql" + err.Error())
	}

	return mysql.NewRepository(database)

}
