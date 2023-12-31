package database

import (
	"context"
	"fmt"
	"venus/internal/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const connectionString = "mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority"

func createConnectionString(dbConf config.DatabaseConfiguration) string {
	return fmt.Sprintf(connectionString, dbConf.Username, dbConf.Password, dbConf.Host)
}

func ConnectToMongoDB(databaseUser config.DatabaseConfiguration) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(createConnectionString(databaseUser)).SetServerAPIOptions(serverAPI)

	innerClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	client = innerClient
	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	if err != nil {
		panic(err)
	}
}

func DisconnectFromMongoDB() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("Pomyślnie rozłączono baze danych")
}
