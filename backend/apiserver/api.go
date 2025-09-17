package main

import (
	"time"

	"apiserver/rabbitmq"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"apiserver/demos"
	"apiserver/parsing"
)

func InitApi(client *rabbitmq.Client) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.SetTrustedProxies(nil)

	api := router.Group("/")
	{
		parsing.Init(api, client)
		demos.Create(api)
	}

	router.Run(":8080")
}