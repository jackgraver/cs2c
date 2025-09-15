package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
    init_rabbit_conn();
    wait_for_message();
    <-forever
}

func init_rabbit_conn() {
    rabbitURL := os.Getenv("RABBITMQ_URL")
    if !rabbitmq {
        log.Fatalf("Missing RabbitMQ URL")
    }
    conn, err := amqp.Dial(rabbitURL)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open channel: %v", err)
    }
    defer ch.Close()

    queue, err := ch.QueueDeclare(
        "demo_jobs",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to declare queue: %v", err)
    }

    msgs, err := ch.Consume(
        queue.Name,
        "",
        true,  // auto-ack
        false, // exclusive
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to register consumer: %v", err)
    }

    log.Println("Parser waiting for messages...")
    forever := make(chan bool)
}

func wait_for_message() {
    go func() {
        for d := range msgs {
            log.Printf("Received message: %s", d.Body)
            parse(d.Body)
        }
    }()
}