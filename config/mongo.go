package config

import (
	"context"
	"go-mongo/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Sugar.Error(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Sugar.Error(err)
	}

	logger.Logger.Info("Connected to MongoDB")
	return client.Database("authentication"), nil
}
