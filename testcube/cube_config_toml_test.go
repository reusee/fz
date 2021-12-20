package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestUpdateCubeConfig(t *testing.T) {
	NewScope().Call(func(
		update UpdateCubeConfigTOML,
	) {

		config, err := update(strings.NewReader(""))
		ce(err)
		config, err = update(bytes.NewReader(config))
		ce(err)

	})
}
