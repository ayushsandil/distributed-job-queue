package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func ClassifyJob(payload string) (string, error) {

	apiKey := os.Getenv("OPENAI_API_KEY")

	reqBody := OpenAIRequest{
		Model: "gpt-4o-mini",
		Messages: []Message{
			{
				Role: "user",
				Content: fmt.Sprintf(
					"Classify this task into one of: email, reporting, etl, analytics, notification. Return only category name. Task: %s",
					payload,
				),
			},
		},
	}

	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest(
		"POST",
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+apiKey,
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result OpenAIResponse

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return "", err
	}

	return result.Choices[0].Message.Content, nil
}