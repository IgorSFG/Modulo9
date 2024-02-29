package main

import (
	"math/rand"
	"time"
	"fmt"
)

func SolarRadiationSensor() string {
	// Generate a random value between 0 and 1280
	measurement := rand.Intn(1281)

	// Calculate the firing rate of the sensor
	firing_rate := 60.0 * 1e9

	// Convert the firing rate for educational purposes
	firing_rate = firing_rate/100.0

	// Sleep for a short duration to simulate real-time measurements
	time.Sleep(time.Duration(firing_rate))

	// Create a text message with the measurement
	text := fmt.Sprintf(`{"sensor": "solar-radiation", "value": "%d"}`, measurement)

	// Return the text message
	return text
}

func Sensor(sensor string) string {
	switch sensor {
	case "solar-radiation":
		return SolarRadiationSensor()
	default:
		return "Invalid sensor type."
	}
}
