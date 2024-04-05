package main

import (
	"fmt"
	"os"

	config "prova2/config"
	
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config.LoadEnv()

	// Configurações do produtor
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"client.id":         os.Getenv("CLUSTER_ID"),
		"auto.offset.reset": "earliest",
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     os.Getenv("API_KEY"),
		"sasl.password":     os.Getenv("API_SECRET"),
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	fmt.Println("Producer started")

	// Enviar mensagem
	topic := "qualidadeAr"

	for {
		message := Sensor("air-quality")
		
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)

		// Aguardar a entrega de todas as mensagens
		producer.Flush(5 * 1000)
		
		// Exibir mensagem enviada
		fmt.Println("")
		fmt.Printf("Message sent: %s\n", message)
	}
}