package aiclientadapter

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/vitorlivi/go-ai-cli/internal/data"
	"github.com/vitorlivi/go-ai-cli/internal/types"
	"google.golang.org/api/option"
)

type GeminiClientAdapter struct {
	client *genai.Client
	config types.AIClientConfig
}

func NewGeminiClient(config types.AIClientConfig) *GeminiClientAdapter {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("GEMINI_API_KEY environment variable is required")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		panic(fmt.Sprintf("Failed to create Gemini client: %v", err))
	}

	if config.Model == "" {
		config.Model = "gemini-pro"
	}

	return &GeminiClientAdapter{
		client: client,
		config: config,
	}
}

func (g *GeminiClientAdapter) Chat(prompt string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := g.client.GenerativeModel(g.config.Model)

	model.SetTemperature(g.config.Temperature)
	model.SetMaxOutputTokens(int32(g.config.MaxTokens))

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	if len(resp.Candidates) == 0 {
		return "Error: No response generated"
	}

	candidate := resp.Candidates[0]
	if candidate.Content == nil || len(candidate.Content.Parts) == 0 {
		return "Error: Empty response"
	}

	var result string
	for _, part := range candidate.Content.Parts {
		if textPart, ok := part.(genai.Text); ok {
			result += string(textPart)
		}
	}

	return result
}

func (g *GeminiClientAdapter) GetAvailableModels() []string {
	models := make([]string, len(data.GeminiModels))
	for i, model := range data.GeminiModels {
		models[i] = model.Id
	}
	return models
}

func (g *GeminiClientAdapter) IsHealthy(ctx context.Context) error {
	model := g.client.GenerativeModel("gemini-pro")

	_, err := model.GenerateContent(ctx, genai.Text("Hello"))
	if err != nil {
		return fmt.Errorf("Gemini API health check failed: %v", err)
	}

	return nil
}

func (g *GeminiClientAdapter) Close() error {
	return g.client.Close()
}
