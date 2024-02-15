package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		text := Sensor("solar-radiation")
		token := client.Publish("sensor/topic", 0, false, text)
		token.Wait()
		fmt.Println("Published: ", text)
		time.Sleep(2 * time.Second)
	}
}