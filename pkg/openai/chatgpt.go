package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gabrielssssssssss/kestrel/internal/config"
)

type Response struct {
	Output []struct {
		Type    string `json:"type"`
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	} `json:"output"`
}

func PromptWebSearch() {}

func PromptTurbo(prompt string) (string, error) {
	var payload Response
	values := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"input": prompt,
	}

	body, err := json.Marshal(values)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/responses", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+config.GetConfig("OPENAI_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		return "", err
	}

	text := payload.Output[0].Content[0].Text
	text = strings.TrimSpace(text)
	text = strings.TrimPrefix(text, "```json")
	text = strings.TrimPrefix(text, "```")
	text = strings.TrimSuffix(text, "```")
	text = strings.TrimSpace(text)
	return text, nil
}
