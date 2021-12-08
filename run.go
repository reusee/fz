package fz

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

		ce(start())

		//TODO run testAction

		ce(stop())

		for _, checker := range checkers {
			ce(checker())
		}

		return
	}
}
