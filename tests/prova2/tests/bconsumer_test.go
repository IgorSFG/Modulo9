package main

import (
	"os"
	"testing"

	config "prova2/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestConsumer(t *testing.T) {
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
		t.Error(err)
	}
	defer consumer.Close()

	t.Log("Consumer started")

	// Assinar tópico
	topic := "qualidadeAr"
	consumer.SubscribeTopics([]string{topic}, nil)

	// Consumir mensagens
	for {
		msg, err := consumer.ReadMessage(-1)
		if msg != nil && err == nil {
			msgValue := string(msg.Value)
			t.Log("")
			t.Logf("Received message: %s\n", msgValue)
			InsertSensors(msgValue)
			t.Log(GetSensors())
			break
		} else if err != nil {
			t.Errorf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}