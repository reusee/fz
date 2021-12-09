package fz

import (
	"encoding/xml"
	"io"
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

func (_ SequentialAction) Type() ActionType {
	return "sequence"
}

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
		var action Action
		err := unmarshalAction(d, &action)
		if is(err, io.EOF) {
			err = nil
			break
		}
		s.Actions = append(s.Actions, action)
	}

	return
}
