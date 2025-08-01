package filters

import (
	"errors"

	commandhandlers "github.com/vitorlivi/go-ai-cli/internal/command-handlers"
	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type EmptyCommandValidationFilter struct{}

func NewEmptyCommandValidationFilter() *EmptyCommandValidationFilter {
	return &EmptyCommandValidationFilter{}
}

func (c *EmptyCommandValidationFilter) Process(context any, input ...any) (any, error) {
	if ctx, ok := context.(types.HasParsedPrompt); ok {
		parsedPrompt := ctx.GetParsedPrompt()
		if parsedPrompt.Prompt == "" &&
			!parsedPrompt.ListModels &&
			!parsedPrompt.ListProviders &&
			!parsedPrompt.ShowHelp &&
			!parsedPrompt.ShowVersion {
			commandhandlers.NewShowHelpHandler().Handle()
			return nil, errors.New("empty command")
		}
	}

	return input, nil
}
