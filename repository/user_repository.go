package repository

import (
	"context"
	"auth-service/config"
	"auth-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func init() {
	db := config.Connect()
	userCollection = db.Collection("users")
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return &user, err
}

func CreateUser(user *models.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}
