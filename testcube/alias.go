package main

import (
	"fmt"

	"github.com/reusee/e4"
)

var (
	ce = e4.Check.With(e4.WrapStacktrace)
	we = e4.Wrap.With(e4.WrapStacktrace)
	he = e4.Handle
	pt = fmt.Printf
)
