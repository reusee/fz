package fz

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/reusee/dscope"
)

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

func (_ ConfigScope) CreateTime() CreatedTime {
	return CreatedTime(time.Now().Format(time.RFC3339))
}

func (_ ConfigScope) CreatedTimeConfig(
	t CreatedTime,
) ConfigMap {
	return ConfigMap{
		"CreatedTime": t,
	}
}

func (_ ConfigScope) UUID() uuid.UUID {
	return uuid.New()
}

func (_ ConfigScope) UUIDConfig(
	id uuid.UUID,
) ConfigMap {
	return ConfigMap{
		"ConfigID": id,
	}
}
