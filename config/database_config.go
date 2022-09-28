package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goter.com.vn/server/environment"
)

func SetupDatabaseConnection() *mongo.Client {
	dataSourceName := fmt.Sprintf("mongodb+srv://%s:%s@%s", environment.DB_USER, environment.DB_PASSWORD, environment.DB_HOST)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().ApplyURI(dataSourceName).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseDatabaseConnection(db *mongo.Client) {

	err := db.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
