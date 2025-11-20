package utils

import (
	"an-backend/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type ResponseBody struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func Chat(userContent string, modelName string, systemPrompt string, assistantAndUser []models.Assistant) string {
	messages := []Message{
		{Role: "system", Content: systemPrompt},
	}

	for _, content := range assistantAndUser {
		if content.UserContent != "" {
			messages = append(messages, Message{
				Role:    "user",
				Content: content.UserContent,
			})
		}
		if content.AssistantContent != "" {
			messages = append(messages, Message{
				Role:    "assistant",
				Content: content.AssistantContent,
			})
		}
	}

	messages = append(messages, Message{
		Role:    "user",
		Content: userContent,
	})

	requestBody := RequestBody{
		Model:       modelName,
		Messages:    messages,
		Temperature: 0.5,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "请求体序列化错误: " + err.Error()
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "创建请求错误: " + err.Error()
	}

	apiKey := os.Getenv("DASHSCOPE_API_KEY")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "发送请求错误: " + err.Error()
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "读取响应错误: " + err.Error()
	}

	var responseBody ResponseBody
	if err := json.Unmarshal(bodyText, &responseBody); err != nil {
		return "解析响应错误: " + err.Error() + "，原始响应: " + string(bodyText)
	}

	if responseBody.Error.Message != "" {
		return "API错误: " + responseBody.Error.Message
	}

	if len(responseBody.Choices) > 0 && responseBody.Choices[0].Message.Content != "" {
		return responseBody.Choices[0].Message.Content
	}

	return "未获取到有效响应"
}
