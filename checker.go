package fz

import (
	"reflect"

	"github.com/reusee/dscope"
)

type Checker struct {
	BeforeStart func() error
	BeforeDo    func() error
	AfterDo     func() error
	AfterStop   func() error
}

type Checkers []Checker

var _ dscope.Reducer = Checkers{}

func (c Checkers) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	return dscope.Reduce(vs)
}
