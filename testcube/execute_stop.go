package main

import "github.com/reusee/fz"

func (_ ExecuteScope) Stop() fz.Stop {
	return func() error {
		//TODO
		return nil
	}
}
