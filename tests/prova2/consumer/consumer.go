package main

import (
	"fmt"
	"os"

	config "prova2/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config.LoadEnv()

	// Configurações do consumidor
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"group.id":          os.Getenv("CLUSTER_ID"),
		"auto.offset.reset": "earliest",
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     os.Getenv("API_KEY"),
		"sasl.password":     os.Getenv("API_SECRET"),
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	fmt.Println("Consumer started")

	// Assinar tópico
	topic := "qualidadeAr"
	consumer.SubscribeTopics([]string{topic}, nil)

	// Consumir mensagens
	for {
		msg, err := consumer.ReadMessage(-1)
		if msg != nil && err == nil {
			msgValue := string(msg.Value)
			fmt.Printf("Received message: %s\n", msgValue)
			InsertSensors(msgValue)
		} else if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}