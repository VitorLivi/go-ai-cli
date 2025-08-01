package commandhandlers

import (
	"fmt"
)

type ShowHelpHandler struct{}

func NewShowHelpHandler() *ShowHelpHandler {
	return &ShowHelpHandler{}
}

func (p *ShowHelpHandler) Handle() (any, error) {
	fmt.Print(`
AI CLI Tool - Chat with AI models from the command line

Usage:
  ai [options]

Options:
  --model, -m <model>     Use specific model (default: gpt-3.5-turbo)
  --prompt, -p <text>     Prompt to send to the AI
  --list-models           List available models
  --list-providers        List available providers
  --help, -h              Show this help message
  --version, -v           Show version information
  --temperature <float>   Set temperature for response generation (0.0-1.0)
  --max-tokens <int>      Set maximum tokens for response
  --timeout <duration>    Set timeout for requests (e.g., 30s, 1m)

Examples:
  ai --prompt "Hello, how are you?"
  ai -p "Explain quantum computing" -m gpt-4
  ai --list-models
  ai --list-providers
  ai -p "Write a poem" --temperature 0.9 --max-tokens 500
  ai --health-check

Environment Variables:
  OPENAI_API_KEY          Required for ChatGPT/GPT models
  GEMINI_API_KEY          Required for Gemini models
`)

	return nil, nil
}
