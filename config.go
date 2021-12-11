package fz

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/reusee/dscope"
)

type ConfigItems []any

var _ dscope.Reducer = ConfigItems{}

func (_ ConfigItems) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	var ret ConfigItems
	names := make(map[string]struct{})
	for _, value := range vs {
		items := value.Interface().(ConfigItems)
		for _, item := range items {
			name := reflect.TypeOf(item).Name()
			if name == "" {
				panic(fmt.Errorf("config item must be named: %T", item))
			}
			if _, ok := names[name]; ok {
				panic(fmt.Errorf("duplicated config: %s", name))
			}
			names[name] = struct{}{}
			ret = append(ret, item)
		}
	}
	return reflect.ValueOf(ret)
}

// built-ins

type CreatedTime string

func (_ ConfigScope) CreateTime() CreatedTime {
	return CreatedTime(time.Now().Format(time.RFC3339))
}

func (_ ConfigScope) CreatedTimeConfig(
	t CreatedTime,
) ConfigItems {
	return ConfigItems{t}
}

func (_ ConfigScope) UUID() uuid.UUID {
	return uuid.New()
}

func (_ ConfigScope) UUIDConfig(
	id uuid.UUID,
) ConfigItems {
	return ConfigItems{id}
}
