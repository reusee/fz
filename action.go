package fz

import (
	"encoding/xml"
	"reflect"

	"github.com/reusee/dscope"
)

type Action interface {
}

type ActionGenerators []func() Action

var _ dscope.Reducer = ActionGenerators{}

func (_ ActionGenerators) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	return dscope.Reduce(vs)
}

type TestAction struct {
	Action Action
}

var _ xml.Unmarshaler = new(TestAction)

func (t *TestAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	if err := unmarshalAction(d, &t.Action); err != nil {
		return we(err)
	}
	if err := d.Skip(); err != nil {
		return we(err)
	}
	return nil
}

func (_ ConfigScope) TestAction(
	generators ActionGenerators,
) (
	action TestAction,
) {
	if len(generators) == 0 {
		panic("no ActionGenerators provided")
	}
	action.Action = RandomActionTree(generators, 128)
	return
}

func (_ ConfigScope) TestActionConfig(
	action TestAction,
) ConfigMap {
	return ConfigMap{
		"TestAction": action,
	}
}
