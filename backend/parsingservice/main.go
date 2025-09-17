package main

import (
	"fmt"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"

	parser "parsingservice/parser"

	"parsingservice/rabbitmq"
)

var (
	msgs <-chan amqp.Delivery
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
	}
	msgs = rabbitClient.Consume()
	if msgs == nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}
	defer rabbitClient.Cleanup()

	waitForMessage()
}

func waitForMessage() {
	for d := range msgs {
		log.Printf("Received message: %s", d.Body)
		if demoID, err := parser.Parse(string(d.Body)); err != nil {
			log.Printf("Failed to parse demo: %v", err)
		} else {
			log.Printf("Demo parsed successfully (%v)", demoID)
		}
	}
}
