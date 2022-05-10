package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	MysqlUrl string
	Secret   string
}

func LoadConfig() (c Config, err error) {
	err = godotenv.Load("./config/envs/dev.env")
	if err != nil {
		log.Fatalf("Some error occued .env Err: %s", err)
	}
	c = Config{Port: os.Getenv("PORT"), MysqlUrl: os.Getenv("MYSQL_URL"), Secret: os.Getenv("SECRET")}
	return
}
