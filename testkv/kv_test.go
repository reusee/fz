package testkv

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/reusee/dscope"
	"github.com/reusee/e4"
	"github.com/reusee/fz"
)

func TestKV(t *testing.T) {
	defer he(nil, e4.TestingFatal(t))

	global := dscope.New(dscope.Methods(new(fz.Global))...)

	configDefs := dscope.Methods(new(fz.ConfigScope))

	// actions
	configDefs = append(configDefs, &fz.ActionGenerators{
		func() fz.Action {
			return ActionSetThenGet{
				Key:   rand.Int63(),
				Value: rand.Int63(),
			}
		},
	})

	// configs
	configDefs = append(configDefs, func() (
		maxClients MaxClients,
		m fz.ConfigMap,
	) {
		maxClients = 8
		m = fz.ConfigMap{
			"MaxClients": maxClients,
		}
		return
	})

	configScope := global.Fork(configDefs...)

	executeDefs := dscope.Methods(new(fz.ExecuteScope))
	executeDefs = append(executeDefs, func(
		maxClients MaxClients,
	) (
		start fz.Start,
		stop fz.Stop,
		do fz.Do,
	) {

		var kv *KV

		// Start
		start = func() error {
			kv = NewKV(int(maxClients))
			return nil
		}

		// Stop
		stop = func() error {
			return nil
		}

		// Do
		do = func(action fz.Action) error {
			switch action := action.(type) {

			case ActionSetThenGet:
				kv.Set(action.Key, action.Value)
				kv.Get(action.Key)
			}

			panic(fmt.Errorf("unknown action: %T", action))
		}

		return
	})

	executeDefs = append(executeDefs, &fz.Operators{
		fz.Operator{
			AfterStop: func() {
				pt("test done\n")
			},
		},
	})

	executeScope := configScope.Fork(executeDefs...)

	executeScope.Call(func(
		execute fz.Execute,
	) {
		ce(execute())
	})

}

type MaxClients int

type ActionSetThenGet struct {
	Key   any
	Value any
}

var _ fz.Action = ActionSetThenGet{}

func (_ ActionSetThenGet) Type() fz.ActionType {
	return "set-then-get"
}
