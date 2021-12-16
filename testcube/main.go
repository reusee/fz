package main

import (
	"github.com/reusee/fz"
)

func run() (err error) {
	defer he(&err)

	globalScope := NewGlobalScope()
	configScope := NewConfigScope(globalScope)
	cubeScope := NewCubeScope(configScope)
	executeScope := NewExecuteScope(cubeScope)

	var execute fz.Execute
	executeScope.Assign(&execute)
	ce(execute())

	return
}

func main() {
	run()
}

func NewTestScope() Scope {
	return NewExecuteScope(
		NewCubeScope(
			NewConfigScope(
				NewGlobalScope(),
			),
		),
	)
}
