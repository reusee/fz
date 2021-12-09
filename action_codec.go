package fz

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"

	"github.com/reusee/e4"
)

var registeredActionTypes sync.Map

func RegisterAction(value any) {
	t := reflect.TypeOf(value)
	name := t.Name()
	if name == "" {
		panic(fmt.Errorf("Action must be named type: %T", value))
	}
	registeredActionTypes.Store(name, t)
}

var ErrActionNotRegistered = errors.New("action not registered")

func unmarshalAction(d *xml.Decoder, target *Action) (err error) {

	var data struct {
		XMLName xml.Name
		Raw     []byte `xml:",innerxml"`
	}
	if err := d.Decode(&data); err != nil {
		return we(err)
	}

	v, ok := registeredActionTypes.Load(data.XMLName.Local)
	if !ok {
		return we.With(
			e4.Info("action: %s", data.XMLName.Local),
		)(ErrActionNotRegistered)
	}
	t := v.(reflect.Type)

	ptr := reflect.New(t)
	valueDecoder := xml.NewDecoder(
		io.MultiReader(
			strings.NewReader("<"),
			strings.NewReader(data.XMLName.Local),
			strings.NewReader(">"),
			bytes.NewReader(data.Raw),
			strings.NewReader("</"),
			strings.NewReader(data.XMLName.Local),
			strings.NewReader(">"),
		),
	)
	ce(valueDecoder.Decode(ptr.Interface()))
	*target = ptr.Elem().Interface().(Action)

	return
}
