package fz

import (
	"bytes"
	"testing"

	"github.com/google/uuid"
	"github.com/reusee/dscope"
	"github.com/reusee/e4"
)

func TestConfigCodec(t *testing.T) {
	defer he(nil, e4.TestingFatal(t))

	scope := dscope.New(dscope.Methods(new(ConfigScope))...)
	scope.Call(func(
		write WriteConfig,
		read ReadConfig,
		createdTime CreatedTime,
		id uuid.UUID,
	) {
		buf := new(bytes.Buffer)
		ce(write(buf))

		decls, err := read(buf)
		ce(err)

		loaded := dscope.New(decls...)
		loaded.Call(func(
			createdTime2 CreatedTime,
			id2 uuid.UUID,
		) {
			if createdTime2 != createdTime {
				t.Fatal()
			}
			if id2 != id {
				t.Fatal()
			}
		})
	})

}
