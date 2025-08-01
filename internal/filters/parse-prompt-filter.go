package filters

import (
	"errors"
	"flag"
	"fmt"

	"github.com/vitorlivi/go-ai-cli/internal/types"
)

type ParsePromptFilter struct {
}

func NewParsePromptFilter() *ParsePromptFilter {
	return &ParsePromptFilter{}
}

func (p *ParsePromptFilter) Process(context any, input ...any) (any, error) {
	var (
		model         = flag.String("model", "", "AI model to use")
		modelShort    = flag.String("m", "", "AI model to use (short)")
		prompt        = flag.String("prompt", "", "Prompt to send to AI")
		promptShort   = flag.String("p", "", "Prompt to send to AI (short)")
		listModels    = flag.Bool("list-models", false, "List available models")
		listProviders = flag.Bool("list-providers", false, "List available providers")
		showHelp      = flag.Bool("help", false, "Show help")
		helpShort     = flag.Bool("h", false, "Show help (short)")
		showVersion   = flag.Bool("version", false, "Show version")
		versionShort  = flag.Bool("v", false, "Show version (short)")
		timeout       = flag.Int("timeout", 0, "Timeout in seconds")
		maxTokens     = flag.Int("max-tokens", 0, "Maximum number of tokens")
		temperature   = flag.Float64("temperature", 1.0, "Temperature for text generation")
	)

	fmt.Print("AFTER PARSE PROMPT FILTER")

	flag.Parse()

	effectiveModel := *model
	if *modelShort != "" {
		effectiveModel = *modelShort
	}

	effectivePrompt := *prompt
	if *promptShort != "" {
		effectivePrompt = *promptShort
	}

	effectiveShowHelp := *showHelp || *helpShort
	effectiveShowVersion := *showVersion || *versionShort

	if ctx, ok := context.(types.HasParsedPrompt); ok {
		ctx.SetParsedPrompt(
			types.ParsedPrompt{
				Model:         effectiveModel,
				Prompt:        effectivePrompt,
				ListModels:    *listModels,
				ListProviders: *listProviders,
				ShowHelp:      effectiveShowHelp,
				ShowVersion:   effectiveShowVersion,
				Timeout:       *timeout,
				MaxTokens:     *maxTokens,
				Temperature:   *temperature,
			})

		return ctx.GetParsedPrompt(), nil
	}

	return nil, errors.New("parse prompt filter context does not implement HasParsedPrompt")
}
