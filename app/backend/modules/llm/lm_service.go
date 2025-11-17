package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Model    string `json:"model"`
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

type BusinessAssistant struct {
	ModelName string
	BaseURL   string
	Client    *http.Client
}

func NewBusinessAssistant() *BusinessAssistant {
	return &BusinessAssistant{
		ModelName: "qwen2.5:7b",
		BaseURL:   "http://localhost:11434",
		Client: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

func (ba *BusinessAssistant) AskQuestion(question, category string) (string, error) {
	systemPrompt := ba.buildSystemPrompt(category)
	fullPrompt := fmt.Sprintf("%s\n\nВопрос предпринимателя: %s", systemPrompt, question)

	requestData := OllamaRequest{
		Model:  ba.ModelName,
		Prompt: fullPrompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	resp, err := ba.Client.Post(
		fmt.Sprintf("%s/api/generate", ba.BaseURL),
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", fmt.Errorf("error making request to Ollama: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama API error: %s - %s", resp.Status, string(body))
	}

	var response OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	return response.Response, nil
}

func (ba *BusinessAssistant) buildSystemPrompt(category string) string {
	basePrompt := `Ты - AI-помощник для владельцев малого бизнеса в России. 
Отвечай подробно, практично и на русском языке. Давай конкретные шаги и примеры.`

	switch category {
	case "legal":
		return basePrompt + "\nСпециализация: ЮРИДИЧЕСКИЕ ВОПРОСЫ - договоры, регистрация, налоги, compliance."
	case "marketing":
		return basePrompt + "\nСпециализация: МАРКЕТИНГ - продвижение, SMM, реклама, брендинг."
	case "finance":
		return basePrompt + "\nСпециализация: ФИНАНСЫ - учет, отчетность, планирование, оптимизация налогов."
	case "management":
		return basePrompt + "\nСпециализация: УПРАВЛЕНИЕ - процессы, найм, KPI, оптимизация."
	default:
		return basePrompt + "\nСпециализация: общие бизнес-вопросы."
	}
}
