package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"weebhook/domain/entities"
)

type PayloadRepository struct{}

func NewPayloadRepository() *PayloadRepository {
	return &PayloadRepository{}
}

// ProcessPullRequestPayload procesa el payload de un evento de Pull Request
func (r *PayloadRepository) ProcessPullRequestPayload(ctx context.Context, payload entities.PullRequestEventPayload) error {
	// Esta implementación podría incluir lógica adicional como registrar el evento,
	// pero por ahora delegamos directamente al método SendDiscordNotification
	return nil
}

// FormatDiscordMessage formatea un mensaje para Discord según la especificación de Embeds
func (r *PayloadRepository) FormatDiscordMessage(payload entities.PullRequestEventPayload) interface{} {
	// Mapeo entre acciones de PR y colores para los embeds
	colorMap := map[string]int{
		"opened":      5025616,  // Verde
		"closed":      15158332, // Rojo
		"reopened":    3447003,  // Azul
		"synchronize": 16776960, // Amarillo
	}

	// Determinar color basado en la acción
	color, exists := colorMap[payload.Action]
	if !exists {
		color = 9807270 // Gris por defecto
	}

	// Crear el mensaje con formato de embeds para Discord
	return map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       fmt.Sprintf("Pull Request #%d: %s", payload.PullRequest.ID, payload.PullRequest.Title),
				"description": payload.PullRequest.Title,
				"url":         payload.PullRequest.URL,
				"color":       color,
				"author": map[string]interface{}{
					"name":     payload.PullRequest.User.Login,
					"icon_url": payload.PullRequest.User.URL,
				},
				"fields": []map[string]interface{}{
					{
						"name":   "Repository",
						"value":  payload.Repository.FullName,
						"inline": true,
					},
					{
						"name":   "Action",
						"value":  payload.Action,
						"inline": true,
					},
					{
						"name":   "State",
						"value":  payload.PullRequest.Base.Ref,
						"inline": true,
					},
				},
			},
		},
	}
}

// SendDiscordNotification envía un mensaje a Discord usando su API de webhooks
func (r *PayloadRepository) SendDiscordNotification(ctx context.Context, webhookURL string, message interface{}) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshalling discord message: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending discord notification: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("discord API responded with status: %d", resp.StatusCode)
	}

	return nil
}
