package fz

import (
	"constraints"
	"fmt"
	"math/rand"
)

func RandBetween[T constraints.Integer, T2 constraints.Integer](target *T, lower T2, upper T2) {
	if lower == upper {
		*target = T(lower)
		return
	}
	if lower > upper {
		panic(fmt.Errorf("bad argument: %d %d", lower, upper))
	}
	gap := upper - lower
	i := rand.Intn(int(gap + 1))
	*target = T(lower) + T(i)
	return
}

func RandBool[T ~bool](target *T) {
	*target = rand.Intn(2) == 0
}
