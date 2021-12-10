package main

import "github.com/reusee/fz"

func (_ ExecuteScope) Operators() fz.Operators {
	return fz.Operators{
		fz.SaveConfig("config.xml"),
		fz.SaveCPUProfile("cpu.profile"),
		fz.SaveAllocsProfile("allocs.profile"),
	}
}
