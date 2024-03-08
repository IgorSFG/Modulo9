package main

import (
	"math/rand"
	"time"
	"fmt"
)

func Refri() string {
	const id = "lj00f00"
	
	// Generate a random measurement between -10 and 20
	measurement := rand.Intn(31) - 10

	// Check if the measurement is too high or too low
	alert := ""
	if measurement > 10 {
		alert = "[ALERTA: Temperatura ALTA]"
	} else if measurement < 2 { 
		alert = "[ALERTA: Temperatura BAIXA]"
	}

	// Create a text message with the measurement
	text := fmt.Sprintf(`
	{
		"id": "%s",
		"tipo": "geladeira",
		"temperatura": "%d",
		"timestamp": "%s"
		"alerta": "%s",
	  }
	`, id, measurement, time.Now().Format(time.RFC3339), alert)

	// Return the text message
	return text
}

// Freezer generates a random temperature measurement for a freezer
func Freezer() string {
	const id = "lj01f01"
	
	// Generate a random measurement between -50 and 0
	measurement := rand.Intn(51) - 50

	// Check if the measurement is too high or too low
	alert := ""
	if measurement > -15 {
		alert = "[ALERTA: Temperatura ALTA]"
	} else if measurement < -25 { 
		alert = "[ALERTA: Temperatura BAIXA]"
	}

	// Create a text message with the measurement
	text := fmt.Sprintf(`
	{
		"id": "%s",
		"tipo": "freezer",
		"temperatura": "%d",
		"timestamp": "%s"
		"alerta": "%s",
	  }
	`, id, measurement, time.Now().Format(time.RFC3339), alert)

	// Return the text message
	return text
}

func Sector(sector string) string {
	switch sector {
	case "freezer":
		return Freezer()
	case "refri":
		return Refri()
	default:
		return "Setor inválido!"
	}
}
