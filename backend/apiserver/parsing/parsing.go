package parsing

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jackgraver/cs2c/rabbitmq"

	"github.com/gin-gonic/gin"
)

var (
	RabbitClient *rabbitmq.Client
)

func Init(router *gin.RouterGroup, client *rabbitmq.Client) {
	RabbitClient = client
	routes(router)
}

func routes(router *gin.RouterGroup) {
	router.POST("/parse", func(c *gin.Context) {
		demoID, err := uploadFile(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "message sent", "demo_id": demoID})
	})
}

func uploadFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid file upload"})
		return "", err
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open uploaded file"})
		return "", err
	}
	defer openedFile.Close()

	saveDir := os.Getenv("STORE_DEMO_PATH")
	if saveDir == "" {
		c.JSON(500, gin.H{"error": "No demo storage directory specified"})
		return "", fmt.Errorf("no demo storage directory specified")
	}

	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	savePath := filepath.Join(saveDir, file.Filename)
	
	out, err := os.Create(savePath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create file"})
		return "", err
	}
	defer out.Close()

	if _, err := io.Copy(out, openedFile); err != nil {
		c.JSON(500, gin.H{"error": "Failed to copy file"})
		return "", err
	}

	body := "parse_demo" // message body, could include demo ID
	err = RabbitClient.Publish("demo_jobs", []byte(body))
	if err != nil {
		log.Printf("Failed to publish: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return "", err
	}
	return file.Filename, nil
}