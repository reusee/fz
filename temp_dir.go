package fz

import (
	"os"
)

type TempDir string

func (_ Def) TempDir() (
	dir TempDir,
	cleanup Cleanup,
) {
	d, err := os.MkdirTemp(os.TempDir(), "")
	ce(err)
	dir = TempDir(d)
	return dir, func() {
		os.RemoveAll(d)
	}
}
