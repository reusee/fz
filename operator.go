package fz

import (
	"reflect"

	"github.com/reusee/dscope"
)

type Operator struct {
	BeforeStart any
	BeforeDo    any
	AfterDo     any
	AfterStop   any
	Finally     any
}

type Operators []Operator

var _ dscope.Reducer = Operators{}

func (c Operators) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	return dscope.Reduce(vs)
}
