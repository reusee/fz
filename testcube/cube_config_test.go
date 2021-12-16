package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestUpdateCubeConfig(t *testing.T) {
	NewTestScope().Call(func(
		update UpdateCubeConfig,
	) {

		config, err := update(strings.NewReader(""))
		ce(err)
		config, err = update(bytes.NewReader(config))
		ce(err)

	})
}
