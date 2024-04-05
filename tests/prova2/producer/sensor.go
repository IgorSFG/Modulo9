package main

import (
	"math/rand"
	"fmt"
	"time"
)

func SensorSimulator(id string, measurement float32) string {
	// Get the current time
	timestamp := time.Now().Format(time.RFC3339)

	// Create a text message with the measurement
	text := fmt.Sprintf(`{
		"idSensor": "%s",
		"timestamp": "%s",
		"tipoPoluente": "PM2.5",
		"nivel": "%f"
		}`, id, timestamp, measurement)
	
	// Return the text message
	return text
}

func Sensor(sensor string) string {
	switch sensor {
	case "air-quality":
		id := "sensor_001"
		measurement := rand.Float32() * 100
		return SensorSimulator(id, measurement)
	default:
		return "Invalid sensor type."
	}
}
