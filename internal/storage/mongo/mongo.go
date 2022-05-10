package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func NewMongo(client *mongo.Client) domain.Repository {
	return &repository{db: db}
}
func (r *repository) add(auth domain.Auth) error {

}
func (r *repository) GetByEmail(email string) *domain.Auth {
	var auth domain.Auth
	c := r.Collection("Auth")
	err := c.FindOne(context.TODO(), bson.D{{"Email", email}}).Decode(&auth)
	if err != nil {
		return &domain.Auth{}, err
	}
	return &auth, nil
}
