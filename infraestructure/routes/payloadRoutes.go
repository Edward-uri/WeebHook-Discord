package routes

import (
	"weebhook/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, webhookHandler *controller.WebhookHandler, reviewHandler *controller.ReviewHandler, actionsHandler *controller.ActionsHandler) {
	routes := router.Group("pull_request")
	{
		routes.POST("/webhook", webhookHandler.HandlePullRequest)
	}

	reviews := router.Group("review")
	{
		reviews.POST("/webhook", reviewHandler.HandleReview)
	}
	actions := router.Group("actions")
    {
        actions.POST("/webhook", actionsHandler.HandleActions)
    }
}
