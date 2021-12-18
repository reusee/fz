package main

import "github.com/reusee/dscope"

type CubeNodeScope func() Scope

func NewCubeNodeScope(parent Scope, defs ...any) Scope {
	defs = append(defs, dscope.Methods(new(CubeNodeScope))...)
	var scope Scope
	defs = append(defs, func() CubeNodeScope {
		return func() Scope {
			return scope
		}
	})
	scope = parent.Fork(defs...)
	return scope
}

type CubeNodeID int
