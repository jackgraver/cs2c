package main

import (
	"fmt"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"

	"github.com/jackgraver/cs2c/rabbitmq"
)

func main() {
	if os.Getenv("ENV") != "prod" {
		// local dev: load .env
		err := dotenv.Load("../.env")
		if err != nil {
			log.Fatalf("Failed to load .env file: %v", err)
		}
	}

	//rabbit connection
	rabbitClient, err := rabbitmq.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	} else {
		fmt.Println("Connected to RabbitMQ")
		rabbitmq.DeclareQueue(rabbitClient.Channel(), "demo_jobs")
	}
	defer rabbitClient.Cleanup()

	//start api
	InitApi(rabbitClient)
}
