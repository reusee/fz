package fz

type Target interface {
	Start() error
	Do(action Action) error
	Stop() error
}
