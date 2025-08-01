package pipelines

import (
	filteradapter "github.com/vitorlivi/go-ai-cli/internal/adapters/filters"
	"github.com/vitorlivi/go-ai-cli/internal/common"
	"github.com/vitorlivi/go-ai-cli/internal/filters"
	"github.com/vitorlivi/go-ai-cli/internal/routers"
	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type MainPipeline struct{}

func NewMainPipeline() *MainPipeline {
	return &MainPipeline{}
}

func (f *MainPipeline) Start(context any) (*types.InteractionType, error) {
	var pipeline = common.NewPipeline[types.InteractionType](context)

	pipelineRouter := routers.NewPipelineRouter(
		NewTalkToAIPipeline(),
		NewGetAuxiliarInformationPipeline(),
	)

	pipelineRouterFilter := filteradapter.NewRouterAsFilterAdapter(pipelineRouter)

	pipeline.AddFilter(filters.NewParsePromptFilter())
	pipeline.AddFilter(filters.NewInteractionTypeFilter())
	pipeline.AddFilter(pipelineRouterFilter)
	pipeline.AddFilter(filters.NewExecutorFilter())

	return pipeline.Start()
}
