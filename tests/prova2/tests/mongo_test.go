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
	Tipo string 		`json:"tipo"`
	Timestamp  string 	`json:"timestamp"`
	TipoPoluente string `json:"tipoPoluente"`
	Nivel string 		`json:"nivel"`
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

	collection := client.Database("Sensors").Collection(sensor.Tipo)

	// Insert a single document
	result, err := collection.InsertOne(context.TODO(), sensor)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %v\n", result.InsertedID)
}

func GetSensors() []Sensor {
	client := ConnectMongoDB()

	defer client.Disconnect(context.TODO())

	// Access a MongoDB collection through a database
	collection := client.Database("Sensors").Collection("air-quality")

	// Find all documents
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var sensors []Sensor
	for cursor.Next(context.TODO()) {
		var sensor Sensor
		err := cursor.Decode(&sensor)
		if err != nil {
			panic(err)
		}
		sensors = append(sensors, sensor)
	}

	return sensors
}