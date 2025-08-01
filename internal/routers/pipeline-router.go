package routers

import (
	"errors"

	"github.com/vitorlivi/go-ai-cli/internal/data"
	pipelinecontext "github.com/vitorlivi/go-ai-cli/internal/pipelines/contexts"
	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type PipelineRouter struct {
	talkToAIPipeline            types.IPipeline
	auxiliarInformationPipeline types.IPipeline
}

func NewPipelineRouter(
	talkToAIPipeline types.IPipeline,
	auxiliarInformationPipeline types.IPipeline,
) *PipelineRouter {
	return &PipelineRouter{
		talkToAIPipeline:            talkToAIPipeline,
		auxiliarInformationPipeline: auxiliarInformationPipeline,
	}
}

func (c *PipelineRouter) Route(context any, input ...any) (any, error) {
	switch ctx := context.(type) {
	case interface {
		types.HasParsedPrompt
		types.HasPipeline
		types.HasInteractionType
	}:
		interactionType := ctx.GetInteractionType()

		switch interactionType {
		case data.InteractionTypes.Main:
			ctx.SetPipeline(c.talkToAIPipeline)
			context := pipelinecontext.NewTalkToAIPipelineContext()
			context.SetParsedPrompt(ctx.GetParsedPrompt())
			ctx.SetPipelineContext(context)
		case data.InteractionTypes.Auxiliary:
			ctx.SetPipeline(c.auxiliarInformationPipeline)
			context := pipelinecontext.NewGetAuxiliarInformationPipelineContext()
			context.SetParsedPrompt(ctx.GetParsedPrompt())
			ctx.SetPipelineContext(context)
		default:
			return nil, errors.New("invalid interaction type")
		}

		return ctx.GetPipeline(), nil
	default:
		return nil, errors.New("context does not implement required types")
	}
}
