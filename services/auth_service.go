package services

import (
	"context"
	"webhook-project/config"
	"webhook-project/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	_, err := config.DB.Collection("users").InsertOne(context.TODO(), user)
	return err
}

func FindUser(username string) (models.User, error) {
	var user models.User
	err := config.DB.Collection("users").
		FindOne(context.TODO(), map[string]string{"username": username}).
		Decode(&user)

	return user, err
}