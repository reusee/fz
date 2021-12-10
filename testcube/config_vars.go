package main

import (
	_ "embed"

	"github.com/reusee/fz"
)

type CubeConfigTOML string

//go:embed cube-config.toml
var defaultCubeConfigTOML CubeConfigTOML

func (_ ConfigScope) CubeConfigTOML() CubeConfigTOML {
	return defaultCubeConfigTOML
}

func (_ ConfigScope) CubeConfigTOMLMap(
	t CubeConfigTOML,
) fz.ConfigMap {
	return fz.ConfigMap{
		"CubeConfig": t,
	}
}
