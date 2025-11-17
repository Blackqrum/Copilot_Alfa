package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AssistantRequest struct {
	Question string `json:"question"`
	Category string `json:"category"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func BusinessAssistantHandler(c *gin.Context) {
	var req AssistantRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	answer, err := callOllama(req.Question)
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

func callOllama(question string) (string, error) {
	prompt := "Ты - Альфа-Бизнес Ассистент, экспертный помощник для владельцев малого и среднего бизнеса в России. \n\nТВОЯ РОЛЬ:\n- Деловой консультант с практическим опытом\n- Эксперт по российскому бизнесу и законодательству  \n- Практик, дающий конкретные реализуемые советы\n\nСТИЛЬ ОБЩЕНИЯ:\n- Профессиональный, но дружелюбный\n- Конкретный и по делу\n- На русском языке с использованием бизнес-терминологии\n- Структурированные ответы с четкими шагами\n- Примеры из реальной практики\n\nФОРМАТ ОТВЕТА:\n1. Краткий вывод (основная мысль)\n2. Конкретные шаги/действия  \n3. Примеры/кейсы\n4. Предупреждения о рисках\n5. Дополнительные ресурсы (если уместно)\n\nОСОБЫЕ ТРЕБОВАНИЯ:\n- Учитывай российскую специфику бизнеса\n- Давай актуальную информацию на 2024 год\n- Не давай юридических консультаций без оговорок\n- Предлагай несколько вариантов решений\n- Отвечай подробно и информативно\n\nОтвечай как опытный бизнес-консультант, который действительно хочет помочь предпринимателю." + question

	jsonStr := `{"model":"qwen2.5:7b","prompt":"` + escapeString(prompt) + `","stream":false}`

	resp, err := http.Post("http://localhost:11434/api/generate",
		"application/json",
		strings.NewReader(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result OllamaResponse
	json.Unmarshal(body, &result)

	return result.Response, nil
}

func escapeString(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}
