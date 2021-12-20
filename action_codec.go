package fz

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
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

func unmarshalAction(d *xml.Decoder, start *xml.StartElement, target *Action) (err error) {

	if start == nil {
		token, err := nextTokenSkipCharData(d)
		if err != nil {
			return we(err)
		}
		s, ok := token.(xml.StartElement)
		if !ok {
			return we(fmt.Errorf("execpting end element"))
		}
		start = &s
	}

	v, ok := registeredActionTypes.Load(start.Name.Local)
	if !ok {
		return we.With(
			e4.Info("action: %s", start.Name.Local),
		)(ErrActionNotRegistered)
	}
	t := v.(reflect.Type)

	ptr := reflect.New(t)
	ce(d.DecodeElement(ptr.Interface(), start))
	*target = ptr.Elem().Interface().(Action)

	return
}
