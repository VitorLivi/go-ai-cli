package filters

import (
	"errors"

	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type ExecutorFilter struct {
}

func NewExecutorFilter() *ExecutorFilter {
	return &ExecutorFilter{}
}

func (c *ExecutorFilter) Process(context any, input ...any) (any, error) {
	if ctx, ok := context.(types.HasCommandHandler); ok {
		return ctx.GetCommandHandler().Handle()
	}

	if ctx, ok := context.(types.HasPipeline); ok {
		pipeline := ctx.GetPipeline()

		if pipeline == nil {
			return nil, errors.New("pipeline is nil")
		}

		return ctx.GetPipeline().Start(ctx.GetPipelineContext())
	}

	return nil, errors.New("executor filter context does not have a command handler")
}
