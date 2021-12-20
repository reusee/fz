package fz

import (
	"encoding/xml"
	"fmt"
)

type SequentialAction struct {
	Actions []Action
}

func init() {
	RegisterAction(SequentialAction{})
}

func Seq(actions ...Action) SequentialAction {
	return SequentialAction{
		Actions: actions,
	}
}

var _ Action = SequentialAction{}

var _ xml.Marshaler = SequentialAction{}

func (s SequentialAction) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	defer he(&err)

	ce(e.EncodeToken(xml.StartElement{
		Name: xml.Name{
			Local: "SequentialAction",
		},
	}))

	for _, action := range s.Actions {
		ce(e.Encode(action))
	}

	ce(e.EncodeToken(xml.EndElement{
		Name: xml.Name{
			Local: "SequentialAction",
		},
	}))

	return
}

var _ xml.Unmarshaler = new(SequentialAction)

func (s *SequentialAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	defer he(&err)

	for {
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
		var action Action
		start := token.(xml.StartElement)
		ce(unmarshalAction(d, &start, &action))
		if action != nil {
			s.Actions = append(s.Actions, action)
		}
	}

}
