package types

type ICommandHandler[O any] interface {
	Handle() (O, error)
}

type HasCommandHandler interface {
	SetCommandHandler(ICommandHandler[any])
	GetCommandHandler() ICommandHandler[any]
}
