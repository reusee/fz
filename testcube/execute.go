package main

import (
	"github.com/reusee/dscope"
	"github.com/reusee/fz"
)

type ExecuteScope func() Scope

func NewExecuteScope(parent Scope) Scope {
	executeDefs := dscope.Methods(new(ExecuteScope))
	executeDefs = append(executeDefs, dscope.Methods(new(fz.ExecuteScope))...)
	var executeScope Scope
	executeDefs = append(executeDefs, func() ExecuteScope {
		return func() Scope {
			return executeScope
		}
	})
	executeScope = parent.Fork(executeDefs...)
	return executeScope
}
