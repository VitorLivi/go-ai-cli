package types

type IPipeline interface {
	Start(context any) (any, error)
}

type HasPipeline interface {
	SetPipeline(IPipeline)
	SetPipelineContext(context any)
	GetPipeline() IPipeline
	GetPipelineContext() any
}
