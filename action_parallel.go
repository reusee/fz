package fz

import (
	"encoding/xml"
	"io"
)

type ParallelAction struct {
	Actions []Action
}

func init() {
	RegisterAction(ParallelAction{})
}

func Par(actions ...Action) ParallelAction {
	return ParallelAction{
		Actions: actions,
	}
}

var _ Action = ParallelAction{}

func (_ ParallelAction) Type() ActionType {
	return "parallel"
}

var _ xml.Marshaler = ParallelAction{}

func (s ParallelAction) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	defer he(&err)

	ce(e.EncodeToken(xml.StartElement{
		Name: xml.Name{
			Local: "ParallelAction",
		},
	}))

	for _, action := range s.Actions {
		ce(e.Encode(action))
	}

	ce(e.EncodeToken(xml.EndElement{
		Name: xml.Name{
			Local: "ParallelAction",
		},
	}))

	return
}

var _ xml.Unmarshaler = new(ParallelAction)

func (s *ParallelAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
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
