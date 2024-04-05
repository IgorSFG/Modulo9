package main

import (
	"math/rand"
	"fmt"
	"time"
)

func SensorSimulator(id string, sensor string, measurement float32) string {
	// Get the current time
	timestamp := time.Now().Format(time.RFC3339)

	level := fmt.Sprintf("%.2f", measurement)

	// Create a text message with the measurement
	text := fmt.Sprintf(`{
		"idSensor": "%s",
		"tipo": "%s",
		"timestamp": "%s",
		"tipoPoluente": "PM2.5",
		"nivel": "%s"
		}`, id, sensor, timestamp, level)
	
	// Return the text message
	return text
}

func Sensors(sensor string) string {
	switch sensor {
	case "air-quality":
		id := "sensor_001"
		measurement := rand.Float32() * 50
		return SensorSimulator(id, sensor, measurement)
	default:
		return "Invalid sensor type."
	}
}
