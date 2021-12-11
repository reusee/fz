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

func (_ ConfigScope) EnableCPUProfile() EnableCPUProfile {
	return false
}

func (_ ConfigScope) EnableCPUProfileConfig(
	enable EnableCPUProfile,
) ConfigItems {
	return ConfigItems{enable}
}

func (_ ExecuteScope) CPUProfile(
	enable EnableCPUProfile,
) Operators {
	if !enable {
		return nil
	}
	return Operators{
		SaveCPUProfile("cpu-profile"),
	}
}
