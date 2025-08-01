package pipelinecontext

import "github.com/vitorlivi/go-ai-cli/internal/types"

type MainPipelineContext struct {
	Pipeline        types.IPipeline
	PipelineContext any
	ParsedPrompt    types.ParsedPrompt
	InteractionType types.InteractionType
}

func NewMainPipelineContext() *MainPipelineContext {
	return &MainPipelineContext{}
}

func (c *MainPipelineContext) SetParsedPrompt(parsedPrompt types.ParsedPrompt) {
	c.ParsedPrompt = parsedPrompt
}

func (c *MainPipelineContext) GetParsedPrompt() types.ParsedPrompt {
	return c.ParsedPrompt
}

func (c *MainPipelineContext) SetPipeline(pipeline types.IPipeline) {
	c.Pipeline = pipeline
}

func (c *MainPipelineContext) GetPipeline() types.IPipeline {
	return c.Pipeline
}

func (c *MainPipelineContext) SetPipelineContext(context any) {
	c.PipelineContext = context
}

func (c *MainPipelineContext) GetPipelineContext() any {
	return c.PipelineContext
}

func (c *MainPipelineContext) SetInteractionType(interactionType types.InteractionType) {
	c.InteractionType = interactionType
}

func (c *MainPipelineContext) GetInteractionType() types.InteractionType {
	return c.InteractionType
}
