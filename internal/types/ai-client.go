package types

import (
	"context"
)

type IAIClient interface {
	Chat(prompt string) string
	GetAvailableModels() []string
	IsHealthy(ctx context.Context) error
}

type AIClientConfig struct {
	Prompt      string
	Model       string
	Temperature float32
	Timeout     int
	MaxTokens   int
}

type HasAIClient interface {
	SetAIClient(IAIClient)
	GetAIClient() IAIClient
}

type HasAIClientConfig interface {
	SetAIClientConfig(AIClientConfig)
	GetAIClientConfig() AIClientConfig
}

type HasAIResponse interface {
	SetAIResponse(string)
	GetAIResponse() string
}
