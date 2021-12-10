package main

import "github.com/reusee/fz"

func (_ ConfigScope) ActionGenerators() fz.ActionGenerators {
	return fz.ActionGenerators{
		func() fz.Action {
			return ActionNoOP{}
		},
		//TODO
	}
}
