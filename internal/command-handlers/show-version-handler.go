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
	message := fmt.Sprintf("AI CLI Tool version %s", data.VERSION)
	fmt.Println(message)

	return &message, nil
}
