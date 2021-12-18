package main

import "github.com/reusee/dscope"

type CubeScope func() Scope

func NewCubeScope(parent Scope) Scope {
	defs := dscope.Methods(new(CubeScope))
	var scope Scope
	defs = append(defs, func() CubeScope {
		return func() Scope {
			return scope
		}
	})
	scope = parent.Fork(defs...)
	return scope
}
