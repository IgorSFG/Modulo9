package main

import (
	"os"
	"fmt"
	"time"

	config "iotsimkafka/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	
	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}
	
	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v", err)
	}

	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		text := Sensor("solar-radiation")
		token := client.Publish("dataMQTT", 0, false, text)
		token.Wait()
		fmt.Println("Published: ", text)
		time.Sleep(5 * time.Second)
	}
}