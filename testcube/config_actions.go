package main

import (
	"github.com/reusee/fz"
)

func init() {
	fz.RegisterAction(ActionNoOP{})
}

type ActionNoOP struct{}
