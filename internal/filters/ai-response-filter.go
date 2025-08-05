package filters

import (
	"errors"
	"fmt"

	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type AIResponseFilter struct{}

func NewAIResponseFilter() *AIResponseFilter {
	return &AIResponseFilter{}
}

func (c *AIResponseFilter) Process(context any, input ...any) (any, error) {
	switch ctx := context.(type) {
	case interface {
		types.HasAIResponse
	}:
		fmt.Printf("\n AI: %v", ctx.GetAIResponse())
		response := ctx.GetAIResponse()

		return &response, nil
	default:
		return nil, errors.New("ai response filter context does not implement required types")
	}
}
