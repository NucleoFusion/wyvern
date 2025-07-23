package models

import "go.mongodb.org/mongo-driver/mongo"

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}
