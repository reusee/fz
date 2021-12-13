package main

import "github.com/reusee/dscope"

type CubeScope struct{}

func NewCubeScope(parent Scope) Scope {
	cubeDefs := dscope.Methods(new(CubeScope))
	cubeScope := parent.Fork(cubeDefs...)
	return cubeScope
}
