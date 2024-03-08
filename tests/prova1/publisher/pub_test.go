package main

import (
	"fmt"
	"time"
	"testing"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func TestPublisher(t *testing.T) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var qos byte = 1
	if qos != 1 {
		t.Error("Invalid QoS")
	}

	sectors := [2]string{"refri", "freezer"}
	index := 0

	for index < 2{
		sector := sectors[index%2]
		text := Sector(sector)
		token := client.Publish("sector/topic", qos, false, text)
		token.Wait()
		fmt.Println(text)
		time.Sleep(2 * time.Second)
		index++
	}
}