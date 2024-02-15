package main

import (
	"math/rand"
	"time"
	"fmt"
)

func SolarRadiationSensor() string {
	// Generate a random value between 0 and 1280
	measurement := rand.Intn(1281)

	firing_rate := 1.0 / 512.0 * 1e9

	// Sleep for a short duration to simulate real-time measurements
	time.Sleep(time.Duration(firing_rate))

	// Create a text message with the measurement
	text := fmt.Sprintf("Solar Radiation Measurement: %d W/m2", measurement)

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
