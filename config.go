package fz

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/reusee/dscope"
)

type ConfigScope struct{}

type ConfigMap map[string]any

var _ dscope.Reducer = ConfigMap{}

func (_ ConfigMap) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	ret := make(ConfigMap)
	for _, value := range vs {
		m := value.Interface().(ConfigMap)
		for k, v := range m {
			if _, ok := ret[k]; ok {
				panic(fmt.Errorf("duplicated config key: %s", k))
			}
			ret[k] = v
		}
	}
	return reflect.ValueOf(ret)
}

// built-ins

type CreatedTime string

func (_ ConfigScope) CreateTime() (
	t CreatedTime,
	m ConfigMap,
) {
	t = CreatedTime(time.Now().Format(time.RFC3339))
	m = ConfigMap{
		"CreatedTime": t,
	}
	return
}

func (_ ConfigScope) UUID() (
	id uuid.UUID,
	m ConfigMap,
) {
	id = uuid.New()
	m = ConfigMap{
		"ConfigID": id,
	}
	return
}

type TestAction struct {
	Action Action
}

var _ xml.Unmarshaler = new(TestAction)

func (t *TestAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	if err := unmarshalAction(d, &t.Action); err != nil {
		return we(err)
	}
	if err := d.Skip(); err != nil {
		return we(err)
	}
	return nil
}

func (_ ConfigScope) DefaultAction() (
	action TestAction,
	m ConfigMap,
) {
	action = TestAction{
		Action: Seq(),
	}
	m = ConfigMap{
		"TestAction": action,
	}
	return
}
