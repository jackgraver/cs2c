package parser

import (
	"fmt"
	"parsingservice/parser/handlers"
)

func CreateHandlers(context *handlers.HandlerContext) {
	handlers.RegisterGrenadeHandlers(context)
	handlers.RegisterPlayerHandler(context)
	handlers.RegisterRoundHandler(context)
	handlers.RegisterBombHandler(context)
	handlers.RegisterKillHandlers(context)
}

func Parse(fileName string) {
	// This function will be called when a message is received from RabbitMQ
	// For now, it's a placeholder that logs the received filename
	fmt.Printf("Received file to parse: %s\n", fileName)
	// TODO: Implement actual file parsing logic here
}
