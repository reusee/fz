package main

import (
	"fmt"

	"github.com/reusee/dscope"
	"github.com/reusee/fz"
)

type Global struct{}

type ConfigScope struct{}

type ExecuteScope struct{}

var (
	pt = fmt.Printf
)

func main() {
	globalDefs := dscope.Methods(new(fz.Global))
	globalDefs = append(globalDefs, dscope.Methods(new(Global))...)
	globalScope := dscope.New(globalDefs...)

	configDefs := dscope.Methods(new(fz.ConfigScope))
	configDefs = append(configDefs, dscope.Methods(new(ConfigScope))...)
	configScope := globalScope.Fork(configDefs...)

	executeDefs := dscope.Methods(new(fz.ExecuteScope))
	executeDefs = append(executeDefs, dscope.Methods(new(ExecuteScope))...)

	executeDefs = append(executeDefs, func() (
		start fz.Start,
		do fz.Do,
		stop fz.Stop,
	) {
		start = func() error {
			pt("start\n")
			return nil
		}
		stop = func() error {
			pt("stop\n")
			return nil
		}
		do = func(action fz.Action) error {
			pt("do action %+v\n", action)
			return nil
		}
		return
	})

	executeScope := configScope.Fork(executeDefs...)

	executeScope.Call(func(
		run fz.Run,
	) {
		if err := run(); err != nil {
			panic(err)
		}
	})
}
