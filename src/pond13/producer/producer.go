package main

import (
	"os"
	"fmt"

	config "iotsimkafka/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config.LoadEnv()

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