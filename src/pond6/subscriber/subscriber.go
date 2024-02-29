package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"regexp"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

func loadEnv() {
	projectDirName := "pond6"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
	}
}

func PostData(message string) {
	const url string = "http://localhost:8000/sensors"

	req, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(message)))
	if err != nil {
		fmt.Println("Error making the request:", err)
	}

	if req.StatusCode == 200 {
		fmt.Println("Data sent successfully")
	} else {
		fmt.Printf("Error sending data: %s\n", req.Status)
	}
	fmt.Println("")
}

func main() {
	loadEnv()

	broker := os.Getenv("BROKER_ADDR")
	port := 8883

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("GoSubscriber")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))

	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received: %s\n", msg.Payload())
		PostData(string(msg.Payload()))
	}
	
	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}
	
	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v", err)
	}

	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("sensor/topic", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscribed running...")
	fmt.Println("")
	select {}
}