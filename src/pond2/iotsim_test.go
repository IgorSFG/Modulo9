package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func testSensor(t *testing.T) uint16 {
	startTime := time.Now()
	
	measurement := rand.Intn(1281)
	firing_rate := 1.0 / 512.0 * 1e9
	
	time.Sleep(time.Duration(firing_rate))
	t.Logf("Solar Radiation Measurement: %d W/m2", measurement)
	
	elapsed_time := float64(time.Since(startTime))
	expected_time := float64(firing_rate)

	if math.Abs(elapsed_time - expected_time) > expected_time {
		elapsed_time = elapsed_time / 1e9
		expected_time = expected_time / 1e9
		
		t.Errorf("Solar Radiation Sensor took %f seconds to execute, expected %f seconds", elapsed_time, expected_time)
	}

	return uint16(measurement)
}

func testPublisher(t *testing.T, measurement uint16) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_test_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	token := client.Publish("test/topic", 0, false, measurement)
	token.Wait()
	
	t.Logf("Published: %d", measurement)

	time.Sleep(2 * time.Second)
}

func testSubscriber(t *testing.T, measurement uint16) {
	var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		t.Logf("Received: %s\n", msg.Payload())

		if fmt.Sprint(-2000) != fmt.Sprint(measurement) {
			t.Errorf("Received: %s, expected: %d", msg.Payload(), measurement)
		}
	}
	
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_test_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	if token := client.Subscribe("test/topic", 1, nil); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	time.Sleep(2 * time.Second)

	t.Error("Test failed")
}

func TestIotsim(t *testing.T) {
	measurement := testSensor(t)

	testPublisher(t, measurement)

	testSubscriber(t, measurement)
}