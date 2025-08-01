package types

type IPipe[O any] interface {
	AddFilter(filter IFilter)
	Start() (*O, error)
}
