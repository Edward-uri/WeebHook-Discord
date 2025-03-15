package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusHandler struct{}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}

func (h *StatusHandler) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
