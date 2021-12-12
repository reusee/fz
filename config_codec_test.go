package fz

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/reusee/dscope"
	"github.com/reusee/e4"
)

func TestConfigCodec(t *testing.T) {
	defer he(nil, e4.TestingFatal(t))

	configDefs := dscope.Methods(new(ConfigScope))
	configDefs = append(configDefs, func() MainAction {
		return MainAction{
			Action: RandomActionTree([]ActionMaker{
				func() Action {
					return Seq()
				},
			}, 128),
		}
	})
	scope := dscope.New(configDefs...)

	scope.Call(func(
		write WriteConfig,
		read ReadConfig,
		createdTime CreatedTime,
		id uuid.UUID,
		scope dscope.Scope,
		action MainAction,
	) {
		buf := new(bytes.Buffer)
		ce(write(buf))

		decls, err := read(buf)
		ce(err)

		loaded := scope.Fork(decls...)
		loaded.Call(func(
			createdTime2 CreatedTime,
			id2 uuid.UUID,
			action2 MainAction,
		) {
			if createdTime2 != createdTime {
				t.Fatal()
			}
			if id2 != id {
				t.Fatal()
			}
			if !reflect.DeepEqual(action2.Action, action.Action) {
				t.Fatal()
			}
		})
	})

}
