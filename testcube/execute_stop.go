package main

import "github.com/reusee/fz"

func (_ Def2) Stop() fz.Stop {
	return func() error {
		//TODO
		return nil
	}
}
