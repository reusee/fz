package main

import (
	"github.com/reusee/dscope"
	"github.com/reusee/fz"
)

type Global struct{}

type ConfigScope struct{}

type ConfigOverwriteScope struct{}

type ExecuteScope struct{}

func run() (err error) {
	defer he(&err)

	globalDefs := dscope.Methods(new(Global))
	globalDefs = append(globalDefs, dscope.Methods(new(fz.Global))...)
	global := dscope.New(globalDefs...)

	configDefs := dscope.Methods(new(ConfigScope))
	configDefs = append(configDefs, dscope.Methods(new(fz.ConfigScope))...)
	configScope := global.Fork(configDefs...)

	configOverwriteDefs := dscope.Methods(new(ConfigOverwriteScope))
	configScope = configScope.Fork(configOverwriteDefs...)

	executeDefs := dscope.Methods(new(ExecuteScope))
	executeDefs = append(executeDefs, dscope.Methods(new(fz.ExecuteScope))...)
	executeScope := configScope.Fork(executeDefs...)

	var execute fz.Execute
	executeScope.Assign(&execute)
	ce(execute())

	return
}

func main() {
	run()
}
