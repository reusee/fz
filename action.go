package fz

import (
	"encoding/xml"
	"fmt"
)

type Action interface {
}

type ActionMaker = func() Action

type MainAction struct {
	Action Action
}

var _ xml.Unmarshaler = new(MainAction)

func (t *MainAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	token, err := nextTokenSkipCharData(d)
	if err != nil {
		return we(err)
	}
	if end, ok := token.(xml.EndElement); ok {
		if end.Name != start.Name {
			return we(xml.UnmarshalError(fmt.Sprintf(
				"expecting end of %s, got %s", start.Name.Local, end.Name.Local)))
		}
		return nil
	}
	start = token.(xml.StartElement)
	if err := unmarshalAction(d, &start, &t.Action); err != nil {
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
