package common

import "github.com/vitorlivi/go-ai-cli/internal/types"

type Pipe[R any] struct {
	context any
	filters []types.IFilter
}

func NewPipeline[R any](context any, filters ...types.IFilter) types.IPipe[R] {
	return &Pipe[R]{
		context: context,
		filters: filters,
	}
}

func (p *Pipe[R]) AddFilter(filter types.IFilter) {
	p.filters = append(p.filters, filter)
}

func (p *Pipe[R]) Start() (*R, error) {
	lastFilterResponse := any(nil)
	for _, filter := range p.filters {
		res, err := filter.Process(p.context, lastFilterResponse)

		if err != nil {
			return nil, err
		}

		lastFilterResponse = res
	}

	return lastFilterResponse.(*R), nil
}
