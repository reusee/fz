package fz

import (
	"math/rand"
	"reflect"

	"github.com/reusee/dscope"
)

type ConfigScope struct{}

type ConfigEntries map[string]any

var _ dscope.Reducer = ConfigEntries{}

func (_ ConfigEntries) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	return dscope.Reduce(vs)
}

// demo

type Foo int64

func (_ ConfigScope) Foo() (
	foo Foo,
	entries ConfigEntries,
) {
	foo = Foo(rand.Int63())
	entries = ConfigEntries{
		"foo": foo,
	}
	return
}

type Bar int64

func (_ ConfigScope) Bar(
	foo Foo,
) (
	bar Bar,
	entries ConfigEntries,
) {
	bar = Bar(foo * 2)
	entries = ConfigEntries{
		"bar": bar,
	}
	return
}
