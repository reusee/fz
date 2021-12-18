package main

import (
	"github.com/reusee/dscope"
	"github.com/reusee/fz"
)

type ConfigScope func() Scope

func NewConfigScope(parent Scope) Scope {
	defs := dscope.Methods(new(ConfigScope))
	defs = append(defs, dscope.Methods(new(fz.ConfigScope))...)
	var scope Scope
	defs = append(defs, func() ConfigScope {
		return func() Scope {
			return scope
		}
	})
	scope = parent.Fork(defs...)
	return scope
}
