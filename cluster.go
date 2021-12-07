package fz

type Cluster interface {
	Start() error
	Do(action Action) error
	Stop() error
}
