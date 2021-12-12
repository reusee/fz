package fz

import (
	"encoding/xml"
)

type Action interface {
}

type ActionMaker = func() Action

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
