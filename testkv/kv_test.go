package testkv

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
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

	// config write
	var writeConfig fz.WriteConfig
	configScope.Assign(&writeConfig)
	f, err := os.Create("config.xml")
	ce(err)
	ce(writeConfig(f))
	ce(f.Close())

	// config read
	var readConfig fz.ReadConfig
	configScope.Assign(&readConfig)
	content, err := os.ReadFile("config.xml")
	ce(err)
	defs, err := readConfig(bytes.NewReader(content))
	ce(err)
	configScope = configScope.Fork(defs...)

	var kv *KV

	executeDefs := dscope.Methods(new(fz.ExecuteScope))
	executeDefs = append(executeDefs, func(
		maxClients MaxClients,
	) (
		start fz.Start,
		stop fz.Stop,
		do fz.Do,
	) {

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
				pt("test done, %d kv operations\n", atomic.LoadInt64(&kv.numOps))
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

func init() {
	fz.RegisterAction(ActionSet{})
	fz.RegisterAction(ActionGet{})
}

type ActionSet struct {
	Key   any
	Value any
}

type ActionGet struct {
	Key any
}
