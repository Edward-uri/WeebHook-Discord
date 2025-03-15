package routes

import (
	"weebhook/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

// rutas de la aplicacion
// Routes define las rutas de la aplicacion
func Routes(router *gin.Engine, webhookHandler *controller.WebhookHandler, reviewHandler *controller.ReviewHandler, statusHandler *controller.StatusHandler) {
	routes := router.Group("pull_request")
	{
		routes.POST("/webhook", webhookHandler.HandlePullRequest)
	}

	reviews := router.Group("review")
	{
		reviews.POST("/webhook", reviewHandler.HandleReview)
	}
	sever := router.Group("server")
	{
		sever.GET("/status", statusHandler.GetStatus)
	}

}
