package fz

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
