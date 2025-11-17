package handlers

import (
	"alfa-backend/modules/llm"

	"github.com/gin-gonic/gin"
)

type AssistantRequest struct {
	Question string `json:"question"`
	Category string `json:"category"`
}

func BusinessAssistantHandler(c *gin.Context) {
	var req AssistantRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	assistant := llm.NewBusinessAssistant()
	answer, err := assistant.AskQuestion(req.Question, req.Category)

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   "Ошибка AI: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"answer":  answer,
	})
}
