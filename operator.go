package fz

import (
	"reflect"

	"github.com/reusee/dscope"
)

type Operator struct {
	BeforeStart func() error
	BeforeDo    func() error
	AfterDo     func() error
	AfterStop   func() error
}

type Operators []Operator

var _ dscope.Reducer = Operators{}

func (c Operators) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	return dscope.Reduce(vs)
}
