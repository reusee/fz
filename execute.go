package fz

//TODO
type ProgramID string

type Run func() error

func (_ ExecuteScope) Run(
	start Start,
	stop Stop,
	do Do,
	testAction TestAction,
	checkers Checkers,
) Run {
	return func() (err error) {
		defer he(&err)

		for _, checker := range checkers {
			if checker.BeforeStart != nil {
				ce(checker.BeforeStart())
			}
		}

		ce(start())

		for _, checker := range checkers {
			if checker.BeforeDo != nil {
				ce(checker.BeforeDo())
			}
		}

		//TODO run testAction

		for _, checker := range checkers {
			if checker.AfterDo != nil {
				ce(checker.AfterDo())
			}
		}

		ce(stop())

		for _, checker := range checkers {
			if checker.AfterStop != nil {
				ce(checker.AfterStop())
			}
		}

		return
	}
}
