package main

import (
	"weebhook/infraestructure"
	"weebhook/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	webhookHandler, reviewHandler, statusHandler := infraestructure.Init()

	router := gin.Default()

	routes.Routes(router, webhookHandler, reviewHandler, statusHandler)
	println("Server running on port 8080")
	router.Run(":8080")
}
