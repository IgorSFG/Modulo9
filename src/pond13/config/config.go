package config

import (
	"os"
	"fmt"
	"regexp"

	godotenv "github.com/joho/godotenv"
)

func LoadEnv() {
	projectDirName := "pond13"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/config/.env`)

	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
}