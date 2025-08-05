package filters

import (
	"errors"

	aiclientadapter "github.com/vitorlivi/go-ai-cli/internal/adapters/clients"
	"github.com/vitorlivi/go-ai-cli/internal/data"
	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type AiClientFilter struct {
}

func NewAiClientFilter() *AiClientFilter {
	return &AiClientFilter{}
}

func (p *AiClientFilter) Process(context any, input ...any) (any, error) {
	switch ctx := context.(type) {
	case interface {
		types.HasAIClient
		types.HasAIClientConfig
	}:
		config := ctx.GetAIClientConfig()

		for _, provider := range data.Providers {
			for _, providerModel := range provider.Models {
				if providerModel.Id == config.Model {
					switch provider.Id {
					case "gpt":
						ctx.SetAIClient(aiclientadapter.NewGPTClient(config))
						return ctx.GetAIClient(), nil
					case "gemini":
						ctx.SetAIClient(aiclientadapter.NewGeminiClient(config))
						return ctx.GetAIClient(), nil
					}
				}
			}
		}
	default:
		return nil, errors.New("ai client filter context does not implement required types")
	}

	return nil, errors.New("model not found")
}
