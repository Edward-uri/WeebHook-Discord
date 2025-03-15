package main

import (
	"weebhook/infraestructure"
	"weebhook/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	webhookHandler, reviewHandler, actionsHandler := infraestructure.Init()

	router := gin.Default()

	routes.Routes(router, webhookHandler, reviewHandler, actionsHandler)

	router.Run(":8080")
}
