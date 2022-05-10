package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConn(host, port, db string) (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(conntext.Background(), 10*time.Second)
	defer concel()

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		concel()
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(db)
	return database, cancel, nil
}
