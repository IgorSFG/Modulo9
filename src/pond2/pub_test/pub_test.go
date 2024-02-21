package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func testSensor(t *testing.T) string {
	startTime := time.Now()
	
	measurement := rand.Intn(1281)
	firing_rate := 60.0 * 1e9
	firing_rate = firing_rate/100.0
	
	time.Sleep(time.Duration(firing_rate))
	text:= fmt.Sprintf("Solar Radiation Measurement: %d W/m2", measurement)
	
	elapsed_time := float64(time.Since(startTime))
	expected_time := float64(firing_rate)

	if math.Abs(elapsed_time - expected_time) > expected_time {
		elapsed_time = elapsed_time / 1e9
		expected_time = expected_time / 1e9
		
		t.Errorf("Solar Radiation Sensor took %f seconds to execute, expected %f seconds", elapsed_time, expected_time)
	}

	return string(text)
}

func testPublisher(t *testing.T, text string) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_test_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	token := client.Publish("test/topic", 0, false, text)
	token.Wait()
	
	t.Log("Published: ", text)
}

func TestIotsim(t *testing.T) {
	text := testSensor(t)

	testPublisher(t, text)
}