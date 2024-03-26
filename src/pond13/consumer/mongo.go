package main

import (
	"context"
	"fmt"
	"os"

	config "iotsimkafka/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	config.LoadEnv()

	// Get the MongoDB connection string
	mongodb := os.Getenv("MONGODB")

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodb).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client
}

type Sensor struct {
	Sensor string `json:"sensor"`
	Value  string `json:"value"`
}

func InsertSensors(client *mongo.Client, data string) {
	var sensor Sensor

	// Access a MongoDB collection through a database
	err := bson.Unmarshal([]byte(data), &sensor)
	if err != nil {
		panic(err)
	}

	collection := client.Database("Sensors").Collection(sensor.Sensor)

	// Insert a single document
	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %v\n", result.InsertedID)
}