package fz

import (
	"sync"

	"github.com/reusee/dscope"
)

type (
	Start func() error
	Stop  func() error
	Do    func(action Action) error
)

type Execute func() error

func (_ ExecuteScope) Execute(
	start Start,
	stop Stop,
	do Do,
	testAction TestAction,
	ops Operators,
	call dscope.Call,
	doAction DoAction,
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
		ce(doAction(testAction.Action))

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

type DoAction func(action Action) error

func (_ ExecuteScope) DoAction(
	do Do,
) (
	doAction DoAction,
) {

	doAction = func(action Action) error {
		switch action := action.(type) {

		case SequentialAction:
			// sequential action
			for _, action := range action.Actions {
				if err := doAction(action); err != nil {
					return err
				}
			}

		case ParallelAction:
			// parallel action
			wg := new(sync.WaitGroup)
			errCh := make(chan error, 1)
			for _, action := range action.Actions {
				select {
				case err := <-errCh:
					return err
				default:
				}
				action := action
				wg.Add(1)
				go func() {
					defer wg.Done()
					err := doAction(action)
					if err != nil {
						select {
						case errCh <- err:
						default:
						}
					}
				}()
			}
			wg.Wait()

		default:
			// send to target
			if err := do(action); err != nil {
				return err
			}

		}
		return nil
	}

	return
}
