package main

import "github.com/reusee/fz"

func (_ Def) MainConfig() fz.MainAction {
	return fz.MainAction{
		Action: fz.RandomActionTree([]fz.ActionMaker{
			func() fz.Action {
				return ActionNoOP{}
			},
		}, 64),
	}
}
