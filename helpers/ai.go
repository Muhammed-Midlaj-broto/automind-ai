package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var openRouterKey string

// InitAI loads your API key from .env file
func InitAI() {
	_ = godotenv.Load()
	openRouterKey = os.Getenv("OPENROUTER_API_KEY")

	if openRouterKey == "" {
		fmt.Println("⚠️ OPENROUTER_API_KEY not found in .env")
	} else {
		fmt.Println("✅ OpenRouter API initialized")
	}
}

// SendAIRequest sends prompt to OpenRouter API
func SendAIRequest(prompt string) (string, error) {
	url := "https://openrouter.ai/api/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "meta-llama/llama-3.1-8b-instruct", // you can also try "openai/gpt-4o-mini"
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	jsonData, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openRouterKey)
	req.Header.Set("HTTP-Referer", "https://railway.app")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("no response from model")
	}

	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	return message["content"].(string), nil
}
