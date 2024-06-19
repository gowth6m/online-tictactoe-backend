package db

import (
	"context"
	"log"
	"online-tictactoe/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE_NAME = "tictactoe"
const COLLECTION_GAME = "Game"

var Client *mongo.Client
var isConnected bool = false

// ConnectToMongoDB connects to MongoDB using the URI provided in the configuration.
// 1. Create a context with a 10-second timeout.
// 2. Connect to MongoDB using the URI from the configuration.
// 3. Ping the MongoDB server to check if the connection is successful.
// 4. If the connection is successful, set the isConnected flag to true.
func ConnectToMongoDB() {
	if Client != nil && isConnected {
		log.Println("Already connected to MongoDB.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig().Database.MongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	log.Println("Connected to MongoDB!")
	Client = client
	isConnected = true
}

// DisconnectFromMongoDB disconnects from MongoDB.
// 1. Check if the client is not nil and isConnected is true.
// 2. Disconnect from MongoDB.
// 3. If the disconnection is successful, set the isConnected flag to false.
func DisconnectFromMongoDB() {
	if Client != nil && isConnected {
		err := Client.Disconnect(context.Background())
		if err != nil {
			log.Fatal("Failed to disconnect from MongoDB: ", err)
		}
		log.Println("Disconnected from MongoDB.")
		isConnected = false
	}
}
