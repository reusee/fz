package main

import (
	"fmt"

	"github.com/reusee/fz"
)

func (_ Def2) Do() fz.Do {
	return func(action fz.Action) error {

		switch action := action.(type) {

		case ActionNoOP:

		default:
			panic(fmt.Errorf("unknown action: %#v", action))

		}

		return nil
	}
}
