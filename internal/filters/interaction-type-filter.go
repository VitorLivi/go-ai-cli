package filters

import (
	"errors"

	"github.com/vitorlivi/go-ai-cli/internal/data"
	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type InteractionTypeFilter struct {
}

func NewInteractionTypeFilter() *InteractionTypeFilter {
	return &InteractionTypeFilter{}
}

func (p *InteractionTypeFilter) Process(context any, input ...any) (any, error) {
	switch ctx := context.(type) {
	case interface {
		types.HasParsedPrompt
		types.HasInteractionType
	}:
		parsedPrompt := ctx.GetParsedPrompt()

		if parsedPrompt.Prompt == "" {
			ctx.SetInteractionType(data.InteractionTypes.Auxiliary)
		} else {
			ctx.SetInteractionType(data.InteractionTypes.Main)
		}

		return ctx.GetInteractionType(), nil
	default:
		return nil, errors.New("interaction type filter context does not implement required types")
	}
}
