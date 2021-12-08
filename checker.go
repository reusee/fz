package fz

import (
	"reflect"

	"github.com/reusee/dscope"
)

type Checker struct {
	BeforeDo func() error
	AfterDo  func() error
}

type Checkers []Checker

var _ dscope.Reducer = Checkers{}

func (c Checkers) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	return dscope.Reduce(vs)
}

func (_ ExecuteScope) Checkers() Checkers {
	return nil
}
