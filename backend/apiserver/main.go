package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitURL := "amqp://guest:guest@localhost:5672/"
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
        "demo_jobs", // name
        true,        // durable
        false,       // delete when unused
        false,       // exclusive
        false,       // no-wait
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to declare queue: %v", err)
    }

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		body := "parse_demo" // message body, could include demo ID
        err = ch.Publish(
            "",         // exchange
            queue.Name, // routing key = queue name
            false,      // mandatory
            false,      // immediate
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(body),
            },
        )
        if err != nil {
            log.Printf("Failed to publish: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "message sent"})
	})

	router.Run()	
}