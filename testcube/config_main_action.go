package main

import "github.com/reusee/fz"

func (_ ConfigScope) MainConfig(
	generators fz.ActionGenerators,
) fz.MainAction {
	return fz.MainAction{
		Action: fz.RandomActionTree(generators, 64),
	}
}
