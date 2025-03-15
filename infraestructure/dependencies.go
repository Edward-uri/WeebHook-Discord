package infraestructure

import (
	"os"
	"weebhook/application"
	"weebhook/infraestructure/controller"
	"weebhook/infraestructure/repositories"

	"github.com/joho/godotenv"
)

func Init() (*controller.WebhookHandler, *controller.ReviewHandler, *controller.ActionsHandler) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	payloadRepo := repositories.NewPayloadRepository()

	discordWebhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

	payloadUseCase := application.NewPayloadUseCase(payloadRepo, discordWebhookURL)
	reviewUseCase := application.NewReviewUseCase(payloadRepo, discordWebhookURL)

	webhookHandler := controller.NewWebhookHandler(*payloadUseCase)
	reviewHandler := controller.NewReviewHandler(*reviewUseCase)

	actionUseCase := application.NewActionUseCase(payloadRepo, discordWebhookURL)
	actionsHandler := controller.NewActionsHandler(actionUseCase)

	return webhookHandler, reviewHandler, actionsHandler
}
