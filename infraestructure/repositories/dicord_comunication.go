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

func (r *PayloadRepository) ProcessPullRequestPayload(ctx context.Context, payload entities.PullRequestEventPayload) error {
	return nil
}

func (r *PayloadRepository) FormatDiscordMessage(payload entities.PullRequestEventPayload) interface{} {
	// Existing code omitted for brevity
	return nil
}

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

func (r *PayloadRepository) ProcessReviewPayload(ctx context.Context, payload entities.ReviewEventPayload) error {
	return nil
}

func (r *PayloadRepository) FormatReviewMessage(payload entities.ReviewEventPayload) interface{} {
	return map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       fmt.Sprintf("Review #%d: %s", payload.Review.ID, payload.Review.State),
				"description": payload.Review.Body,
				"url":         payload.PullRequest.URL,
				"color":       3447003,
				"author": map[string]interface{}{
					"name":     payload.Review.User.Login,
					"icon_url": payload.Review.User.URL,
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
						"value":  payload.Review.State,
						"inline": true,
					},
				},
			},
		},
	}
}
