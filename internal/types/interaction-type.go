package types

type InteractionType string

type InteractionTypesStruct struct {
	Auxiliary InteractionType
	Main      InteractionType
}

type HasInteractionType interface {
	SetInteractionType(InteractionType)
	GetInteractionType() InteractionType
}
