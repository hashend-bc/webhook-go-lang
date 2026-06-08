package models

import "time"

type AccessLog struct {
	IPAddress string    `bson:"ip_address"`
	Endpoint  string    `bson:"endpoint"`
	Status    string    `bson:"status"`
	Error     string    `bson:"error"`
	Time      time.Time `bson:"time"`
}