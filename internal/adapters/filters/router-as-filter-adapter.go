package filteradapter

import (
	"github.com/vitorlivi/go-ai-cli/internal/types"
	common "github.com/vitorlivi/go-ai-cli/internal/types"
)

type RouterAsFilterAdapter struct {
	router common.IRouter
}

func NewRouterAsFilterAdapter(router types.IRouter) *RouterAsFilterAdapter {
	return &RouterAsFilterAdapter{
		router: router,
	}
}

func (r *RouterAsFilterAdapter) Process(ctx any, input ...any) (any, error) {
	return r.router.Route(ctx, input...)
}
