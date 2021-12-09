package fz

import (
	"encoding/xml"
	"reflect"

	"github.com/reusee/dscope"
)

type Action interface {
	Type() ActionType
}

type ActionType string

type ActionGenerators []func() Action

var _ dscope.Reducer = ActionGenerators{}

func (_ ActionGenerators) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	return dscope.Reduce(vs)
}

func (_ ConfigScope) ActionGenerators() ActionGenerators {
	return nil
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

func (_ ConfigScope) DefaultAction(
	generators ActionGenerators,
) (
	action TestAction,
	m ConfigMap,
) {
	action.Action = RandomActionTree(generators, 1024)
	m = ConfigMap{
		"TestAction": action,
	}
	return
}
