package commandhandlers

import (
	"github.com/vitorlivi/go-ai-cli/internal/data"
)

type ListProvidersHandler struct{}

func NewListProvidersHandler() *ListProvidersHandler {
	return &ListProvidersHandler{}
}

func (p *ListProvidersHandler) Handle() (any, error) {
	return data.Providers, nil
}
