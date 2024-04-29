//

package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Initialize the MongoDB client
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URL")

	// Set up client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create a new client
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB!")
}

// GetClient returns the MongoDB client instance
func GetClient() *mongo.Client {
	return client
}

// OpenCollection opens a collection with the given name
func OpenCollection(collectionName string) *mongo.Collection {
	return client.Database("cluster0").Collection(collectionName)
}

//package database
//
//import (
//	"context"
//	"fmt"
//	"github.com/joho/godotenv"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//	"os"
//	"time"
//)
//
//func DBinstance() *mongo.Client {
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	MongoDb := os.Getenv("MONGODB_URL")
//
//	clientOptions := options.Client().ApplyURI(MongoDb)
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	client, err := mongo.Connect(ctx, clientOptions)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Successfully connected to MongoDB!")
//
//	return client
//}
//
//var client *mongo.Client = DBinstance()
