package mongodb

import (
	"context"
	"fmt"
	"time"

	"wyvern-server/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*models.Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use internal Docker hostname (service name in docker-compose)
	uri := "mongodb://root:password@wyvern_mongo:27017"
	dbName := "wyvern"

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect error: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping error: %w", err)
	}

	db := client.Database(dbName)

	fmt.Println("[CONNECTED] MongoDB")

	return &models.Mongo{
		Client:   client,
		Database: db,
	}, nil
}
