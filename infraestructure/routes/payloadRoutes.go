package routes

import (
    "weebhook/infraestructure/controller"
    "github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, webhookHandler *controller.WebhookHandler) {
    routes := router.Group("pull_request")
    {
        routes.POST("/webhook", webhookHandler.HandlePullRequest)
    }
}