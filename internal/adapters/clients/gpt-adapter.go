package aiclientadapter

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/vitorlivi/go-ai-cli/internal/data"
	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type GPTClient struct {
	client *openai.Client
	config types.AIClientConfig
}

func NewGPTClient(config types.AIClientConfig) *GPTClient {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY environment variable is required")
	}

	client := openai.NewClient(apiKey)

	if config.Model == "" {
		config.Model = "gpt-3.5-turbo"
	}

	return &GPTClient{
		client: client,
		config: config,
	}
}

func (g *GPTClient) Chat(prompt string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := openai.ChatCompletionRequest{
		Model: g.config.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens:   g.config.MaxTokens,
		Temperature: g.config.Temperature,
	}

	resp, err := g.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "Error: No response generated"
	}

	return resp.Choices[0].Message.Content
}

func (g *GPTClient) GetAvailableModels() []string {
	models := make([]string, len(data.GPTModels))
	for i, model := range data.GPTModels {
		models[i] = model.Id
	}
	return models
}

func (g *GPTClient) IsHealthy(ctx context.Context) error {
	req := openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
			},
		},
		MaxTokens: 1,
	}

	_, err := g.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return fmt.Errorf("GPT API health check failed: %v", err)
	}

	return nil
}
