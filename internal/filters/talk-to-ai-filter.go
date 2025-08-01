package filters

import (
	"errors"

	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type TalkToAIFilter struct{}

func NewTalkToAIFilter() *TalkToAIFilter {
	return &TalkToAIFilter{}
}

func (c *TalkToAIFilter) Process(context any, input ...any) (any, error) {
	switch ctx := context.(type) {
	case interface {
		types.HasAIClient
		types.HasAIClientConfig
	}:
		res := ctx.GetAIClient().Chat(ctx.GetAIClientConfig().Prompt)
		return res, nil
	default:
		return nil, errors.New("talk to ai filter context does not implement required types")
	}
}
