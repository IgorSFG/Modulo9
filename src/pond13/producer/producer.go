package main

import (
	"os"
	"fmt"
	//"time"
	"regexp"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func loadEnv() {
	projectDirName := "pond6"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
}

func main() {
	loadEnv()

	broker := os.Getenv("BROKER_ADDR")
	port := 8883

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("GoPublisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	
	// Configurações do produtor
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"client.id":         "go-producer",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	// Enviar mensagem
	topic := "test_topic"
	for {
		message := Sensor("solar-radiation")
		producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		}, nil)
		
		// Aguardar a entrega de todas as mensagens
		producer.Flush(15 * 1000)
	}
}