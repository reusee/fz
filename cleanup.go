package fz

import (
	"reflect"

	"github.com/reusee/dscope"
)

type Cleanup func()

var _ dscope.Reducer = Cleanup(nil)

func (_ Cleanup) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	return dscope.Reduce(vs)
}

func (_ Def) DumbCleanup() Cleanup {
	return func() {}
}
