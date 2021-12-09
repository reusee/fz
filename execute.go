package fz

//TODO
type ProgramID string

type Execute func() error

func (_ ExecuteScope) Execute(
	start Start,
	stop Stop,
	do Do,
	testAction TestAction,
	ops Operators,
) Execute {
	return func() (err error) {
		defer he(&err)

		for _, op := range ops {
			if op.BeforeStart != nil {
				ce(op.BeforeStart())
			}
		}

		if start == nil {
			panic("Start not provided")
		}
		ce(start())

		for _, op := range ops {
			if op.BeforeDo != nil {
				ce(op.BeforeDo())
			}
		}

		if do == nil {
			panic("Do not provided")
		}
		//TODO run testAction

		for _, op := range ops {
			if op.AfterDo != nil {
				ce(op.AfterDo())
			}
		}

		if stop == nil {
			panic("Stop not provided")
		}
		ce(stop())

		for _, op := range ops {
			if op.AfterStop != nil {
				ce(op.AfterStop())
			}
		}

		return
	}
}
