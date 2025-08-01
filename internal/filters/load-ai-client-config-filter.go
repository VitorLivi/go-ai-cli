package filters

import (
	"errors"

	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type LoadAIClientConfigFilter struct {
}

func NewLoadAIClientConfigFilter() *LoadAIClientConfigFilter {
	return &LoadAIClientConfigFilter{}
}

func (p *LoadAIClientConfigFilter) Process(context any, input ...any) (any, error) {
	switch ctx := context.(type) {
	case interface {
		types.HasParsedPrompt
		types.HasAIClientConfig
		types.HasInteractionType
	}:
		parsed := ctx.GetParsedPrompt()

		ctx.SetAIClientConfig(types.AIClientConfig{
			Model:       parsed.Model,
			Prompt:      parsed.Prompt,
			Temperature: parsed.Temperature,
			Timeout:     parsed.Timeout,
			MaxTokens:   parsed.MaxTokens,
		})

		return ctx.GetAIClientConfig(), nil
	default:
		return nil, errors.New("load ai client config filter context does not implement required types")
	}
}
