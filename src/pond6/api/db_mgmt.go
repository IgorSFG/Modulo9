package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"os"
	"path/filepath"
)

var database = func () string {
	currentDir, err := os.Getwd()
	parentDir := filepath.Dir(currentDir)
	if err != nil {
		fmt.Println(err)
	}

	dbPath := filepath.Join(parentDir, "database/data.db")

	return dbPath
}

func PostSensor(sensor string, value string) string {
	db, _ := sql.Open("sqlite3", database())
	defer db.Close()

	stmt := `INSERT INTO data(sensor, value) VALUES (?, ?)`
	statement, err := db.Prepare(stmt)
	if err != nil {
		fmt.Println(err.Error())
		return "Error preparing the statement"
	}

	_, err = statement.Exec(sensor, value)
	if err != nil {
		fmt.Println(err.Error())
		return "Error executing the statement"
	}

	response :=fmt.Sprintf("%s with value %s added successfully", sensor, value)

	return response
}

type SQLSensor struct {
	ID     string `json:"id"`
	Sensor string `json:"sensor"`
	Value  string `json:"value"`
}

func GetSensors() []SQLSensor {
	db, _ := sql.Open("sqlite3", database())
	defer db.Close()

	row, err := db.Query("SELECT * FROM data ORDER BY sensor")
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	var sensors []SQLSensor
	
	for row.Next() {
		var id string
		var sensor string
		var value string
		row.Scan(&id, &sensor, &value)
		fmt.Printf("ID: %s, Sensor: %s, Value: %s\n", id, sensor, value)
		sensors = append(sensors, SQLSensor{id, sensor, value})
	}

	return sensors
}

func CreateTable() {
	db, _ := sql.Open("sqlite3", database())
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS data
	(id INTEGER PRIMARY KEY AUTOINCREMENT, sensor TEXT, value TEXT);
	`

	command, err := db.Prepare(sqlStmt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	command.Exec()
}