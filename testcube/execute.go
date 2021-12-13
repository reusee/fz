package main

import (
	"github.com/reusee/dscope"
	"github.com/reusee/fz"
)

type ExecuteScope struct{}

func NewExecuteScope(parent Scope) Scope {
	executeDefs := dscope.Methods(new(ExecuteScope))
	executeDefs = append(executeDefs, dscope.Methods(new(fz.ExecuteScope))...)
	executeScope := parent.Fork(executeDefs...)
	return executeScope
}
