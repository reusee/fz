package fz

import "os"

func SaveConfigToFile(filename string) Operator {
	return Operator{
		AfterStop: func(
			writeConfig WriteConfig,
		) {
			f, err := os.Create(filename)
			ce(err)
			defer f.Close()
			ce(writeConfig(f))
		},
	}
}
