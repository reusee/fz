package main

import (
	"github.com/reusee/dscope"
	"github.com/reusee/fz"
)

type ConfigScope struct{}

type ConfigOverwriteScope struct{}

func NewConfigScope(parent Scope) Scope {
	configDefs := dscope.Methods(new(ConfigScope))
	configDefs = append(configDefs, dscope.Methods(new(fz.ConfigScope))...)
	configScope := parent.Fork(configDefs...)

	configOverwriteDefs := dscope.Methods(new(ConfigOverwriteScope))
	configScope = configScope.Fork(configOverwriteDefs...)

	return configScope
}
