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

type MainAction struct {
	Action Action
}

var _ xml.Unmarshaler = new(MainAction)

func (t *MainAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	if err := unmarshalAction(d, &t.Action); err != nil {
		return we(err)
	}
	if err := d.Skip(); err != nil {
		return we(err)
	}
	return nil
}

func (_ ConfigScope) MainActionConfig(
	action MainAction,
) ConfigItems {
	return ConfigItems{action}
}
