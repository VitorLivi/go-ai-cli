package commandhandlers

import (
	"fmt"

	"github.com/vitorlivi/go-ai-cli/internal/data"
)

type ListProvidersHandler struct{}

func NewListProvidersHandler() *ListProvidersHandler {
	return &ListProvidersHandler{}
}

func (p *ListProvidersHandler) Handle() (any, error) {
	var formatted string = "Providers: \n"

	for _, provider := range data.Providers {
		formatted += fmt.Sprintf("- %s\n", provider.Id)
	}

	formatted += fmt.Sprintf("\nTotal Providers: %d\n", len(data.Providers))

	fmt.Println(formatted)

	return &formatted, nil
}
