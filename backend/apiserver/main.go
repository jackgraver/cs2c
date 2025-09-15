package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"

    "apiserver"
)

func main() {
    init_rabbit_conn();
    apiserver.api();
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
}