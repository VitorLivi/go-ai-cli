package types

type IFilter interface {
	Process(context any, input ...any) (any, error)
}

type ParsedPrompt struct {
	Model         string
	Prompt        string
	ListModels    bool
	ListProviders bool
	ShowHelp      bool
	ShowVersion   bool
	Timeout       int
	MaxTokens     int
	Temperature   float32
}

type HasParsedPrompt interface {
	SetParsedPrompt(ParsedPrompt)
	GetParsedPrompt() ParsedPrompt
}
