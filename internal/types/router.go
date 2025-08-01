package types

type IRouter interface {
	Route(ctx any, input ...any) (any, error)
}
