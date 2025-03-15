package application

import (
    "context"
    "weebhook/domain"
    "weebhook/domain/entities"
)

type ActionUseCase struct {
    repo               domain.IPayloadRepository
    discordWeebHookURL string
}

func NewActionUseCase(repo domain.IPayloadRepository, weebHookURL string) *ActionUseCase {
    return &ActionUseCase{
        repo:               repo,
        discordWeebHookURL: weebHookURL,
    }
}

func (uc *ActionUseCase) ProcessAction(ctx context.Context, payload *entities.ActionEventPayload) error {
    discordMessage := uc.repo.FormatActionMessage(*payload)
    err := uc.repo.SendDiscordNotification(ctx, uc.discordWeebHookURL, discordMessage)
    if err != nil {
        return err
    }
    return nil
}