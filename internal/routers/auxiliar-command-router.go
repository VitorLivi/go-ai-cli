package routers

import (
	"errors"

	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type AuxiliarCommandRouter struct {
	listModelsHandler    types.ICommandHandler[any]
	listProvidersHandler types.ICommandHandler[any]
	showHelpHandler      types.ICommandHandler[any]
	showVersionHandler   types.ICommandHandler[any]
}

func NewAuxiliarCommandRouter(
	listModelsHandler types.ICommandHandler[any],
	listProvidersHandler types.ICommandHandler[any],
	showHelpHandler types.ICommandHandler[any],
	showVersionHandler types.ICommandHandler[any],
) *AuxiliarCommandRouter {
	return &AuxiliarCommandRouter{
		listModelsHandler:    listModelsHandler,
		listProvidersHandler: listProvidersHandler,
		showHelpHandler:      showHelpHandler,
		showVersionHandler:   showVersionHandler,
	}
}

func (c *AuxiliarCommandRouter) Route(context any, input ...any) (any, error) {
	switch ctx := context.(type) {
	case interface {
		types.HasParsedPrompt
		types.HasAIClientConfig
		types.HasCommandHandler
	}:
		parsedPrompt := ctx.GetParsedPrompt()

		if parsedPrompt.ListModels {
			ctx.SetCommandHandler(c.listModelsHandler)
			return c.listModelsHandler, nil
		}

		if parsedPrompt.ListProviders {
			ctx.SetCommandHandler(c.listProvidersHandler)
			return c.listProvidersHandler, nil
		}

		if parsedPrompt.ShowHelp {
			ctx.SetCommandHandler(c.showHelpHandler)
			return c.showHelpHandler, nil
		}

		if parsedPrompt.ShowVersion {
			ctx.SetCommandHandler(c.showVersionHandler)
			return c.showVersionHandler, nil
		}
	default:
		return nil, errors.New("context does not implement required types")
	}

	return nil, nil
}
