package commandhandlers

import (
	"github.com/vitorlivi/go-ai-cli/internal/data"
)

type ListModelsHandler struct{}

func NewListModelsHandler() *ListModelsHandler {
	return &ListModelsHandler{}
}

func (p *ListModelsHandler) Handle() (any, error) {
	models := []data.Model{}

	for _, provider := range data.Providers {
		for _, model := range provider.Models {
			models = append(models, model)
		}
	}

	return models, nil
}
