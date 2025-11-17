package handlers

import (
	"encoding/json"
	"fmt"
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
	Model    string `json:"model"`
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

var userMessages []string
var conversationSummary string

func BusinessAssistantHandler(c *gin.Context) {
	var req AssistantRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	userMessages = append(userMessages, req.Question)
	if len(userMessages) > 5 {
		userMessages = userMessages[len(userMessages)-5:]
	}

	conversationSummary = updateSummary(conversationSummary, req.Question)

	answer, err := callOllama(req.Question, req.Category, conversationSummary)
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

func updateSummary(currentSummary, newQuestion string) string {
	newKeywords := extractKeywords(newQuestion)

	if currentSummary == "" {
		return "Ключевые темы: " + newKeywords
	}

	updatedSummary := currentSummary + "; " + newKeywords

	if len(updatedSummary) > 150 {
		parts := strings.Split(updatedSummary, ";")
		if len(parts) > 2 {
			updatedSummary = "Последние темы: " + parts[len(parts)-2] + "; " + parts[len(parts)-1]
		} else {
			updatedSummary = "Темы: " + parts[len(parts)-1]
		}
	}

	return updatedSummary
}

func extractKeywords(text string) string {
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, "!", "")
	words := strings.Fields(text)

	keywords := []string{}

	for _, word := range words {
		word = strings.ToLower(word)

		if len(word) < 4 || isCommonWord(word) {
			continue
		}

		keywords = append(keywords, word)

		if len(keywords) >= 4 {
			break
		}
	}

	if len(keywords) == 0 && len(words) >= 3 {
		return words[0] + " " + words[1] + " " + words[2]
	}

	return strings.Join(keywords, ", ")
}

func isCommonWord(word string) bool {
	commonWords := []string{
		"как", "что", "где", "когда", "почему", "зачем",
		"можно", "нужно", "хочу", "должен", "могу",
		"есть", "быть", "стать", "сделать", "хотел",
		"очень", "более", "менее", "хорошо", "плохо",
		"если", "тот", "этот", "такой", "какой",
		"вот", "так", "еще", "уже", "только",
		"после", "перед", "потом", "сейчас", "тогда",
		"мой", "твой", "наш", "ваш", "свой",
		"или", "и", "но", "а", "да", "нет",
	}

	for _, common := range commonWords {
		if word == common {
			return true
		}
	}
	return false
}

func callOllama(question, category, context string) (string, error) {
	prompt := buildPrompt(question, category, context)

	requestData := map[string]interface{}{
		"model":  "qwen2.5:7b",
		"prompt": prompt,
		"stream": false,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		strings.NewReader(string(jsonData)),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Ollama API error: %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	var result OllamaResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.Response, nil
}
func buildPrompt(question, category, context string) string {
	basePrompt := `Ты - Альфа-Бизнес Ассистент, экспертный помощник для владельцев малого и среднего бизнеса в России.

ТВОЯ РОЛЬ:
- Деловой консультант с практическим опытом
- Эксперт по российскому бизнесу и законодательству
- Практик, дающий конкретные реализуемые советы

СТИЛЬ ОБЩЕНИЯ:
- Профессиональный, но дружелюбный
- Конкретный и по делу
- На русском языке с использованием бизнес-терминологии
- Структурированные ответы с четкими шагами
- Примеры из реальной практики

ЖЕСТКИЕ ТРЕБОВАНИЯ К ЯЗЫКУ:
- ОТВЕЧАЙ ТОЛЬКО НА РУССКОМ ЯЗЫКЕ
- НЕ ИСПОЛЬЗУЙ АНГЛИЙСКИЙ, КИТАЙСКИЙ ИЛИ ДРУГИЕ ЯЗЫКИ
- ВСЕ ПРИМЕРЫ, ТЕРМИНЫ И ОТВЕТЫ ДОЛЖНЫ БЫТЬ НА РУССКОМ



ОСОБЫЕ ТРЕБОВАНИЯ:
- Учитывай российскую специфику бизнеса
- Давай актуальную информацию на 2024 год
- Не давай юридических консультаций без оговорок
- Предлагай несколько вариантов решений`

	switch category {
	case "legal":
		basePrompt += "\n\nСПЕЦИАЛИЗАЦИЯ: ЮРИДИЧЕСКИЕ ВОПРОСЫ - договоры, регистрация бизнеса, налоги, compliance, трудовое право."
	case "marketing":
		basePrompt += "\n\nСПЕЦИАЛИЗАЦИЯ: МАРКЕТИНГ И ПРОДВИЖЕНИЕ - SMM, таргетированная реклама, брендинг, маркетинговые стратегии."
	case "finance":
		basePrompt += "\n\nСПЕЦИАЛИЗАЦИЯ: ФИНАНСЫ - учет, отчетность, налоговое планирование, оптимизация расходов, финансовая аналитика."
	case "management":
		basePrompt += "\n\nСПЕЦИАЛИЗАЦИЯ: УПРАВЛЕНИЕ БИЗНЕСОМ - процессы, найм персонала, KPI, оптимизация операционной деятельности."
	default:
		basePrompt += "\n\nСПЕЦИАЛИЗАЦИЯ: ОБЩИЕ БИЗНЕС-ВОПРОСЫ - стратегия, развитие, проблемы управления, масштабирование."
	}

	if context != "" {
		basePrompt += "\n\nКОНТЕКСТ РАЗГОВОРА: " + context +
			"\n\nУЧТИ КОНТЕКСТ ВЫШЕ ПРИ ОТВЕТЕ! Отвечай с учетом указанных тем."
	}

	basePrompt += "\n\nТЕКУЩИЙ ВОПРОС: " + question

	return basePrompt
}
