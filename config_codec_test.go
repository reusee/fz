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

	defs := dscope.Methods(new(Def))
	defs = append(defs, func() MainAction {
		return MainAction{
			Action: RandomActionTree([]ActionMaker{
				func() Action {
					return Seq()
				},
			}, 128),
		}
	})
	scope := dscope.New(defs...)

	scope.Call(func(
		write WriteConfig,
		read ReadConfig,
		createdAt CreatedAt,
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
			createdAt2 CreatedAt,
			id2 uuid.UUID,
			action2 MainAction,
		) {
			if createdAt2 != createdAt {
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
