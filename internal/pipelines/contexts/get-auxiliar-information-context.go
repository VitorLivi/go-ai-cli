package pipelinecontext

import "github.com/vitorlivi/go-ai-cli/internal/types"

type GetAuxiliarInformationPipelineContext struct {
	ParsedPrompt   types.ParsedPrompt
	AIClientConfig types.AIClientConfig
	CommandHandler types.ICommandHandler[any]
}

func (c *GetAuxiliarInformationPipelineContext) SetParsedPrompt(parsedPrompt types.ParsedPrompt) {
	c.ParsedPrompt = parsedPrompt
}

func (c *GetAuxiliarInformationPipelineContext) GetParsedPrompt() types.ParsedPrompt {
	return c.ParsedPrompt
}

func (c *GetAuxiliarInformationPipelineContext) SetCommandHandler(commandHandler types.ICommandHandler[any]) {
	c.CommandHandler = commandHandler
}

func (c *GetAuxiliarInformationPipelineContext) GetCommandHandler() types.ICommandHandler[any] {
	return c.CommandHandler
}

func NewGetAuxiliarInformationPipelineContext() *GetAuxiliarInformationPipelineContext {
	return &GetAuxiliarInformationPipelineContext{}
}
