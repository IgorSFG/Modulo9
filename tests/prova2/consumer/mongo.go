package main

import (
	"context"
	"fmt"
	"os"

	config "prova2/config"

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
	IdSensor string 	`json:"idSensor"`
	Timestamp  string 	`json:"timestamp"`
	TipoPoluente string `json:"tipoPoluente"`
	Nivel float32 		`json:"nivel"`
}

func InsertSensors(data string) {
	client := ConnectMongoDB()

	defer client.Disconnect(context.TODO())

	var sensor Sensor

	// Access a MongoDB collection through a database
	err := bson.UnmarshalExtJSON([]byte(data), true, &sensor)
	if err != nil {
		panic(err)
	}

	collection := client.Database("Sensors").Collection(sensor.IdSensor)

	// Insert a single document
	result, err := collection.InsertOne(context.TODO(), sensor)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %v\n", result.InsertedID)
}