package main

import (
	"fmt"
	"math/rand"
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

	sectors := [2]string{"refri", "freezer"}
	index := 0

	for {
		sector := sectors[index%2]
		text := Sector(sector)
		token := client.Publish("sector/topic", 0, false, text)
		token.Wait()
		fmt.Println(text)
		time.Sleep(2 * time.Second)
		index++
	}
}