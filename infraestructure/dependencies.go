package infraestructure

import (
    "weebhook/application"
    "weebhook/infraestructure/controller"
    "weebhook/infraestructure/repositories"
    "os"
    "github.com/joho/godotenv"
)

func Init() (*controller.WebhookHandler) {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    payloadRepo := repositories.NewPayloadRepository()

    discordWebhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

    payloadUseCase := application.NewPayloadUseCase(payloadRepo, discordWebhookURL)

    webhookHandler := controller.NewWebhookHandler(*payloadUseCase)

    return webhookHandler
}