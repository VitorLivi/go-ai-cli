package pipelinecontext

import "github.com/vitorlivi/go-ai-cli/internal/types"

type GetAuxiliarInformationPipelineContext struct {
	ParsedPrompt types.ParsedPrompt
}

func (c *GetAuxiliarInformationPipelineContext) SetParsedPrompt(parsedPrompt types.ParsedPrompt) {
	c.ParsedPrompt = parsedPrompt
}

func (c *GetAuxiliarInformationPipelineContext) GetParsedPrompt() types.ParsedPrompt {
	return c.ParsedPrompt
}

func NewGetAuxiliarInformationPipelineContext() GetAuxiliarInformationPipelineContext {
	return GetAuxiliarInformationPipelineContext{}
}
