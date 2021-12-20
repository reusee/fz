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

	defs := dscope.Methods(new(fz.Def))

	// actions
	defs = append(defs, &fz.MainAction{
		Action: fz.RandomActionTree([]fz.ActionMaker{
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
		}, 128),
	})

	// provide configs
	defs = append(defs, func() MaxClients {
		return 8
	}, func(
		maxClients MaxClients,
	) fz.ConfigItems {
		return fz.ConfigItems{maxClients}
	})

	scope := dscope.New(defs...)

	// overwrite
	defs = defs[:0]

	defs = append(defs, func() fz.EnableCPUProfile {
		return true
	})

	scope = scope.Fork(defs...)

	// config write
	var writeConfig fz.WriteConfig
	scope.Assign(&writeConfig)
	f, err := os.Create("config.xml")
	ce(err)
	ce(writeConfig(f))
	ce(f.Close())

	// config read
	var readConfig fz.ReadConfig
	scope.Assign(&readConfig)
	content, err := os.ReadFile("config.xml")
	ce(err)
	defs, err = readConfig(bytes.NewReader(content))
	ce(err)
	scope = scope.Fork(defs...)

	var kv *KV

	scope = scope.Fork(
		func(
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
		},
		&fz.Operators{
			fz.Operator{
				AfterStop: func() {
					pt("test done, %d kv operations\n", atomic.LoadInt64(&kv.numOps))
				},
			},
		},
	)

	scope.Call(func(
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
