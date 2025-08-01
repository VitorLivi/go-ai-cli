package pipelines

import (
	"github.com/vitorlivi/go-ai-cli/internal/common"
	"github.com/vitorlivi/go-ai-cli/internal/filters"
	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type TalkToAIPipeline struct{}

func NewTalkToAIPipeline() *TalkToAIPipeline {
	return &TalkToAIPipeline{}
}

func (f *TalkToAIPipeline) Start(context any) (any, error) {
	var pipeline = common.NewPipeline[types.InteractionType](context)

	pipeline.AddFilter(filters.NewEmptyCommandValidationFilter())
	pipeline.AddFilter(filters.NewLoadAIClientConfigFilter())
	pipeline.AddFilter(filters.NewAiClientFilter())
	pipeline.AddFilter(filters.NewTalkToAIFilter())

	return pipeline.Start()
}
