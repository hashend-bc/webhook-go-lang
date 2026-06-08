package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WhitelistIP struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	IPAddress string             `bson:"ip_address"`
	Enabled   bool               `bson:"enabled"`
}