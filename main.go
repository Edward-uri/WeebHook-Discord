package main

import (
    "weebhook/infraestructure"
    "weebhook/infraestructure/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    webhookHandler := infraestructure.Init()

    router := gin.Default()

    routes.Routes(router, webhookHandler)

    router.Run(":8080")
}