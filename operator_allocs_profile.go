package fz

import (
	"os"
	"runtime/pprof"
)

func SaveAllocsProfile(filename string) Operator {
	return Operator{
		AfterStop: func() {
			profile := pprof.Lookup("allocs")
			if profile == nil {
				return
			}
			f, err := os.Create(filename)
			ce(err)
			defer f.Close()
			ce(profile.WriteTo(f, 0))
		},
	}
}
