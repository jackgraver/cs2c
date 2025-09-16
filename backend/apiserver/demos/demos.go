package demos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Create(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		// TODO: Implement database integration
		ctx.JSON(200, gin.H{
			"demos": []string{},
		})
	})

	router.GET("/demo/:demo_id/round/:round_num", func(c *gin.Context) {
		demoID := c.Param("demo_id")
		roundNum := c.Param("round_num")

		filePath := filepath.Join("./parsed_demos", demoID, fmt.Sprintf("r%s.json", roundNum))

		data, err := os.ReadFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file", "details": err.Error()})
			return
		}

		var roundJSON interface{}
		if err := json.Unmarshal(data, &roundJSON); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON", "details": err.Error()})
			return
		}

		response := gin.H{
			"round_data": roundJSON,
		}

		c.JSON(http.StatusOK, response)
	})

	router.GET("/demo/:demo_id", func(c *gin.Context) {
		// demoID := c.Param("demo_id")
		// TODO: Implement utils.ReadRounds integration
		response := gin.H{
			"demo_rounds": []string{},
		}
		c.JSON(http.StatusOK, response)
	})	
}