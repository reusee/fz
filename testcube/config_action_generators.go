package main

import "github.com/reusee/fz"

func (_ ConfigScope) ActionGenerators() fz.ActionGenerators {
	return fz.ActionGenerators{
		//TODO
		func() fz.Action {
			return fz.Seq()
		},
	}
}
