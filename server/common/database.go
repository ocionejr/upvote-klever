package common

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase(config *Config) *mongo.Database {
	connString := fmt.Sprintf("mongodb://%s:%s@%s:%s/", config.DBUser, config.DBPassword, config.DBHost, config.DBPort)
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(config.DBName)

	return database
}