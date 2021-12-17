package main

import "github.com/reusee/dscope"

type CubeNodeScope struct{}

func NewCubeNodeScope(parent Scope, defs ...any) Scope {
	defs = append(defs, dscope.Methods(new(CubeNodeScope))...)
	return parent.Fork(defs...)
}

type CubeNodeID int
