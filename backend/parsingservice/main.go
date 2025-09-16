package main

import (
	"fmt"
	"log"

	dotenv "github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"

	parser "parsingservice/parser"

	"github.com/jackgraver/cs2c/rabbitmq"
)

var (
	msgs <-chan amqp.Delivery
)

func main() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
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
	fmt.Println("message", len(msgs))
	for d := range msgs {
		log.Printf("Received message: %s", d.Body)
		parser.Parse(string(d.Body))
	}
}
