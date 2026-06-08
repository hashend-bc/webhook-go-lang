package services

import (
	"context"
	"time"

	"webhook-project/config"
	"webhook-project/models"
)

func SaveLog(ip string, endpoint string, status string, errMsg string) error {

	log := models.AccessLog{
		IPAddress: ip,
		Endpoint:  endpoint,
		Status:    status,
		Error:     errMsg,
		Time:      time.Now(),
	}

	_, err := config.DB.Collection("logs").InsertOne(context.TODO(), log)
	return err
}