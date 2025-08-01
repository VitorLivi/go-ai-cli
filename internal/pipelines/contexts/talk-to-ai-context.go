package pipelinecontext

import (
	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type TalkToAIPipelineContext struct {
	ParsedPrompt    types.ParsedPrompt
	AiClientConfig  types.AIClientConfig
	InteractionType types.InteractionType
	AIClient        types.IAIClient
}

func NewTalkToAIPipelineContext() *TalkToAIPipelineContext {
	return &TalkToAIPipelineContext{}
}

func (c *TalkToAIPipelineContext) SetParsedPrompt(parsedPrompt types.ParsedPrompt) {
	c.ParsedPrompt = parsedPrompt
}

func (c *TalkToAIPipelineContext) GetParsedPrompt() types.ParsedPrompt {
	return c.ParsedPrompt
}

func (c *TalkToAIPipelineContext) SetAIClientConfig(aiClientConfig types.AIClientConfig) {
	c.AiClientConfig = aiClientConfig
}

func (c *TalkToAIPipelineContext) GetAIClientConfig() types.AIClientConfig {
	return c.AiClientConfig
}

func (c *TalkToAIPipelineContext) SetAIClient(aiClient types.IAIClient) {
	c.AIClient = aiClient
}

func (c *TalkToAIPipelineContext) GetAIClient() types.IAIClient {
	return c.AIClient
}

func (c *TalkToAIPipelineContext) SetInteractionType(interactionType types.InteractionType) {
	c.InteractionType = interactionType
}

func (c *TalkToAIPipelineContext) GetInteractionType() types.InteractionType {
	return c.InteractionType
}
