package main

import (
	"github.com/reusee/dscope"
	"github.com/reusee/fz"
)

type Def struct{}

type Def2 struct{}

func NewScope() Scope {
	var defs []any
	defs = append(defs, dscope.Methods(new(fz.Def))...)
	defs = append(defs, dscope.Methods(new(Def))...)
	scope := dscope.New(defs...)
	defs = defs[:0]
	defs = append(defs, dscope.Methods(new(Def2))...)
	scope = scope.Fork(defs...)
	return scope
}
