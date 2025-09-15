package parsing

func create(router *gin.RouterGroup) {
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
}