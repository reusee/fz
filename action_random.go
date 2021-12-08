package fz

import "math/rand"

func RandomAction(args ...Action) Action {
	actions := make([]Action, len(args))
	copy(actions, args)
	rand.Shuffle(len(actions), func(i, j int) {
		actions[i], actions[j] = actions[j], actions[i]
	})
	return Seq(actions...)
}
