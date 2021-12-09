package fz

import "math/rand"

func InterleaveAction(args ...Action) Action {
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
			if len(right) > 0 {
				cut := rand.Intn(len(right))
				left = append(left, right[:cut]...)
				right = right[cut:]
			}
			left = append(left, step)
		}
		actions = append(left, right...)
	}
	return Seq(actions...)
}
