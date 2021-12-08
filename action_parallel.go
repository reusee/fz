package fz

import "encoding/xml"

type ParallelAction struct {
	Actions []Action
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
