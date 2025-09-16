package main

import (
	"fmt"
	"log"

	"github.com/jackgraver/cs2c/rabbitmq"

	dotenv "github.com/joho/godotenv"
)

func main() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	rabbitClient, err := rabbitmq.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	} else {
		fmt.Println("Connected to RabbitMQ")
	}
	defer rabbitClient.Cleanup()
}
