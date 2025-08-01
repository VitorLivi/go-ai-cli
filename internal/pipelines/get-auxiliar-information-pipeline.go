package pipelines

import (
	filteradapter "github.com/vitorlivi/go-ai-cli/internal/adapters/filters"
	commandhandlers "github.com/vitorlivi/go-ai-cli/internal/command-handlers"
	"github.com/vitorlivi/go-ai-cli/internal/common"
	"github.com/vitorlivi/go-ai-cli/internal/filters"
	"github.com/vitorlivi/go-ai-cli/internal/routers"
)

type GetAuxiliarInformationPipeline struct{}

func NewGetAuxiliarInformationPipeline() *GetAuxiliarInformationPipeline {
	return &GetAuxiliarInformationPipeline{}
}

func (f *GetAuxiliarInformationPipeline) Start(context any) (any, error) {
	var pipeline = common.NewPipeline[any](context)

	auxiliarCommandRouter := routers.NewAuxiliarCommandRouter(
		commandhandlers.NewListModelsHandler(),
		commandhandlers.NewListProvidersHandler(),
		commandhandlers.NewShowHelpHandler(),
		commandhandlers.NewShowVersionHandler(),
	)

	auxiliarCommandRouterFilter := filteradapter.NewRouterAsFilterAdapter(auxiliarCommandRouter)

	pipeline.AddFilter(filters.NewEmptyCommandValidationFilter())
	pipeline.AddFilter(auxiliarCommandRouterFilter)
	pipeline.AddFilter(filters.NewExecutorFilter())

	return pipeline.Start()
}
