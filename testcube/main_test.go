package main

import (
	"testing"

	"github.com/reusee/e4"
)

func TestRun(t *testing.T) {
	defer he(nil, e4.TestingFatal(t))
	ce(run())
}
