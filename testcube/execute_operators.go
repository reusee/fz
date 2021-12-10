package main

import "github.com/reusee/fz"

func (_ ExecuteScope) Operators() fz.Operators {
	return fz.Operators{
		fz.SaveConfigToFile("config.xml"),
		fz.NewCPUProfiler("cpu.profile"),
		fz.SaveAllocsProfile("allocs.profile"),
	}
}
