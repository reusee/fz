package main

import (
	"github.com/reusee/dscope"
	"github.com/reusee/fz"
)

type Global struct{}

func NewGlobalScope() Scope {
	globalDefs := dscope.Methods(new(Global))
	globalDefs = append(globalDefs, dscope.Methods(new(fz.Global))...)
	global := dscope.New(globalDefs...)
	return global
}
