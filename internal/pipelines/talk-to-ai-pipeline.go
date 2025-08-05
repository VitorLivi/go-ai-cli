package pipelines

import (
	"github.com/vitorlivi/go-ai-cli/internal/common"
	"github.com/vitorlivi/go-ai-cli/internal/filters"
)

type TalkToAIPipeline struct{}

func NewTalkToAIPipeline() *TalkToAIPipeline {
	return &TalkToAIPipeline{}
}

func (f *TalkToAIPipeline) Start(context any) (any, error) {
	var pipeline = common.NewPipeline[string](context)

	pipeline.AddFilter(filters.NewEmptyCommandValidationFilter())
	pipeline.AddFilter(filters.NewLoadAIClientConfigFilter())
	pipeline.AddFilter(filters.NewAiClientFilter())
	pipeline.AddFilter(filters.NewTalkToAIFilter())
	pipeline.AddFilter(filters.NewAIResponseFilter())

	return pipeline.Start()
}
