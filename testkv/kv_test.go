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
			key := rand.Int63()
			value := rand.Int63()
			return fz.Seq(
				ActionSet{
					Key:   key,
					Value: value,
				},
				ActionGet{
					Key: key,
				},
			)
		},
	})

	// provide configs
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

	// overwrite fz configs
	configScope = configScope.Fork(
		func() fz.EnableCPUProfile {
			return true
		},
	)

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

			case ActionSet:
				kv.Set(action.Key, action.Value)

			case ActionGet:
				kv.Get(action.Key)

			default:
				panic(fmt.Errorf("unknown action: %T", action))
			}
			return nil
		}

		return
	})

	// operators
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
		writeConfig fz.WriteConfig,
	) {
		ce(execute())

		// display config file
		//ce(writeConfig(os.Stdout))
	})

}

type MaxClients int

type ActionSet struct {
	Key   any
	Value any
}

type ActionGet struct {
	Key any
}
