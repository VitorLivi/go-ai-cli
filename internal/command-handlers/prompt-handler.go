package commandhandlers

import (
	"github.com/vitorlivi/go-ai-cli/internal/data"
)

type PromptHandler struct{}

func NewPromptHandler() *PromptHandler {
	return &PromptHandler{}
}

func (p *PromptHandler) Handle() ([]data.Provider, error) {
	return data.Providers, nil
}
