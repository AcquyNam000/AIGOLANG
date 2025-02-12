package handlers

import (
	"Bai3/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DialogHandler struct {
	Service *services.DialogService
}

func NewDialogHandler(service *services.DialogService) *DialogHandler {
	return &DialogHandler{Service: service}
}

func (h *DialogHandler) ProcessDialog(c *gin.Context) {
	var request struct {
		Prompt string `json:"prompt"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	dialog, words, err := h.Service.ProcessDialog(request.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ✅ Trả về JSON đúng format
	c.JSON(http.StatusOK, gin.H{
		"dialog": dialog,
		"words":  words,
	})
}

