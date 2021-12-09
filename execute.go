package fz

import "github.com/reusee/dscope"

//TODO
type ProgramID string

type Execute func() error

func (_ ExecuteScope) Execute(
	start Start,
	stop Stop,
	do Do,
	testAction TestAction,
	ops Operators,
	call dscope.Call,
) Execute {
	return func() (err error) {
		defer he(&err)

		defer func() {
			for _, op := range ops {
				if op.Finally != nil {
					call(op.Finally)
				}
			}
		}()

		for _, op := range ops {
			if op.BeforeStart != nil {
				call(op.BeforeStart)
			}
		}

		if start == nil {
			panic("Start not provided")
		}
		ce(start())

		for _, op := range ops {
			if op.BeforeDo != nil {
				call(op.BeforeDo)
			}
		}

		if do == nil {
			panic("Do not provided")
		}
		//TODO run testAction

		for _, op := range ops {
			if op.AfterDo != nil {
				call(op.AfterDo)
			}
		}

		if stop == nil {
			panic("Stop not provided")
		}
		ce(stop())

		for _, op := range ops {
			if op.AfterStop != nil {
				call(op.AfterStop)
			}
		}

		return
	}
}
