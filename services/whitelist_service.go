package services

import (
	"context"
	"webhook-project/config"
	"webhook-project/models"
)

func AddIP(ip models.WhitelistIP) error {
	_, err := config.DB.Collection("whitelist").InsertOne(context.TODO(), ip)
	return err
}