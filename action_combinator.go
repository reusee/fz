package fz

import "math/rand"

// compound action types

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

// compound action generators

func Rand(args ...Action) Action {
	actions := make([]Action, len(args))
	copy(actions, args)
	rand.Shuffle(len(actions), func(i, j int) {
		actions[i], actions[j] = actions[j], actions[i]
	})
	return Seq(actions...)
}

func Interleave(args ...Action) Action {
	var actions []Action
	for _, arg := range args {
		seq, ok := arg.(SequentialAction)
		if !ok {
			actions = append(actions, arg)
			continue
		}
		steps := seq.Actions
		var left []Action
		right := actions
		for len(steps) > 0 {
			step := steps[0]
			steps = steps[1:]
			cut := rand.Intn(len(right))
			left = append(left, right[:cut]...)
			left = append(left, step)
			right = right[cut:]
		}
		actions = append(left, right...)
	}
	return Seq(actions...)
}
