package commandhandlers

import (
	"fmt"

	"github.com/vitorlivi/go-ai-cli/internal/data"
)

type ListModelsHandler struct{}

func NewListModelsHandler() *ListModelsHandler {
	return &ListModelsHandler{}
}

func (p *ListModelsHandler) Handle() (any, error) {
	models := []data.Model{}

	var formatted string
	for _, provider := range data.Providers {
		formatted += fmt.Sprintf("\nProvider: %s\n", provider.Id)

		for _, model := range provider.Models {
			models = append(models, model)
			formatted += fmt.Sprintf("- %s\n", model.Id)
		}
	}

	formatted += fmt.Sprintf("\nTotal Models: %d\n", len(models))

	fmt.Println(formatted)
	return &formatted, nil
}
