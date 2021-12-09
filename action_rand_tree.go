package fz

import "math/rand"

func RandomActionTree(
	generators ActionGenerators,
	numActions int,
) Action {

	var split func([]Action) Action
	split = func(actions []Action) Action {
		if len(actions) == 0 {
			return nil
		}
		if len(actions) <= 10 {
			return leafRandTreeFuncs[rand.Intn(len(leafRandTreeFuncs))](actions)
		}
		nSplit := 3
		step := len(actions) / nSplit
		var actions2 []Action
		for i := step; i < len(actions); i += step {
			action := split(actions[:i])
			if action != nil {
				actions2 = append(actions2, action)
			}
			actions = actions[i:]
		}
		action := split(actions)
		if action != nil {
			actions2 = append(actions2, action)
		}
		return split(actions2)
	}

	var actions []Action
	for i := 0; i < numActions; i++ {
		actions = append(actions, generators[rand.Intn(len(generators))]())
	}

	return split(actions)
}

var leafRandTreeFuncs = []func([]Action) Action{
	func(actions []Action) Action {
		return Seq(actions...)
	},
	func(actions []Action) Action {
		return RandomAction(actions...)
	},
	func(actions []Action) Action {
		return Par(actions...)
	},
	func(actions []Action) Action {
		return InterleaveAction(actions...)
	},
}
