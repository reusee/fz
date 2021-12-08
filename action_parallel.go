package fz

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
