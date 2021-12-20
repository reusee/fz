package fz

import (
	"os"
	"runtime/pprof"
)

func SaveCPUProfile(filename string) Operator {
	var f *os.File
	return Operator{
		BeforeStart: func() {
			var err error
			f, err = os.Create(filename)
			ce(err)
			pprof.StartCPUProfile(f)
		},
		AfterStop: func() {
			pprof.StopCPUProfile()
			ce(f.Close())
			f = nil
		},
		Finally: func() {
			if f != nil {
				f.Close()
			}
		},
	}
}

type EnableCPUProfile bool

func (_ Def) EnableCPUProfile() EnableCPUProfile {
	return false
}

func (_ Def) EnableCPUProfileConfig(
	enable EnableCPUProfile,
) ConfigItems {
	return ConfigItems{enable}
}

func (_ Def) CPUProfile(
	enable EnableCPUProfile,
) Operators {
	if !enable {
		return nil
	}
	return Operators{
		SaveCPUProfile("cpu-profile"),
	}
}
