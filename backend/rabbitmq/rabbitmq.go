package rabbitmq

import (
	"context"
	"fmt"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Client represents a RabbitMQ client
type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// Connect establishes a connection to RabbitMQ
func Connect() (*Client, error) {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		return nil, fmt.Errorf("missing RabbitMQ URL env var")
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	return &Client{
		conn:    conn,
		channel: channel,
	}, nil
}

func (c *Client) Cleanup() {
	err := c.Close()
	if err != nil {
		fmt.Printf("Error closing RabbitMQ connection: %v", err)
	}
}

// Publish publishes a message to the specified queue
func (c *Client) Publish(queue string, body []byte) error {
	// Declare the queue to ensure it exists
	_, err := c.channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publish the message
	err = c.channel.PublishWithContext(
		ctx,
		"",     // exchange
		queue,  // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

// Close closes the RabbitMQ connection and channel
func (c *Client) Close() error {
	var err error
	
	if c.channel != nil {
		if channelErr := c.channel.Close(); channelErr != nil {
			err = fmt.Errorf("failed to close channel: %w", channelErr)
		}
	}
	
	if c.conn != nil {
		if connErr := c.conn.Close(); connErr != nil {
			if err != nil {
				err = fmt.Errorf("%v; failed to close connection: %w", err, connErr)
			} else {
				err = fmt.Errorf("failed to close connection: %w", connErr)
			}
		}
	}
	
	return err
}
