package fz

type Action interface {
	Type() ActionType
}

type ActionType string
