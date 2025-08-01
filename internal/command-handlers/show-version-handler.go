package commandhandlers

import (
	"fmt"

	"github.com/vitorlivi/go-ai-cli/internal/data"
)

type ShowVersionHandler struct{}

func NewShowVersionHandler() *ShowVersionHandler {
	return &ShowVersionHandler{}
}

func (p *ShowVersionHandler) Handle() (any, error) {
	fmt.Printf("AI CLI Tool version %s\n", data.VERSION)

	return nil, nil
}
