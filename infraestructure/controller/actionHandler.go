package controller

import (
    "net/http"
    "weebhook/application"
    "weebhook/domain/entities"
    "github.com/gin-gonic/gin"
)

type ActionsHandler struct {
    useCase *application.ActionUseCase
}

func NewActionsHandler(useCase *application.ActionUseCase) *ActionsHandler {
    return &ActionsHandler{useCase: useCase}
}

func (h *ActionsHandler) HandleActions(c *gin.Context) {
    var actionPayload entities.ActionEventPayload
    if err := c.ShouldBindJSON(&actionPayload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := h.useCase.ProcessAction(c.Request.Context(), &actionPayload)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "GitHub Actions event received"})
}