package fz

import "encoding/xml"

type SequentialAction struct {
	Actions []Action
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
