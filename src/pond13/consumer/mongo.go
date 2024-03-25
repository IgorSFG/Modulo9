package main

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadEnv() {
	projectDirName := "pond6"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
	}
}

func ConnectMongoDB() *mongo.Client {
	loadEnv()

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

	collectionName := sensor.Sensor
	collection := client.Database("Sensors").Collection(collectionName)

	// Insert a single document
	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %v\n", result.InsertedID)
}