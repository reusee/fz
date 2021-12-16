package main

import (
	"bytes"
	"encoding"
	"io"
	"reflect"

	"github.com/matrixorigin/matrixcube/config"
	"github.com/pelletier/go-toml/v2"
)

type UpdateCubeConfigTOML func(r io.Reader) (result []byte, err error)

func (_ CubeScope) UpdateCubeConfigTOML() UpdateCubeConfigTOML {
	return func(r io.Reader) (result []byte, err error) {
		defer he(&err)

		config := reflect.New(
			tomlOnlyType(reflect.TypeOf((*config.Config)(nil)).Elem()),
		)
		ce(toml.NewDecoder(r).Decode(config.Interface()))
		buf := new(bytes.Buffer)
		ce(toml.NewEncoder(buf).Encode(config.Elem().Interface()))
		result = buf.Bytes()

		return
	}
}

func tomlOnlyType(t reflect.Type) reflect.Type {
	if t.Kind() != reflect.Struct {
		return t
	}
	var newFields []reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("toml") == "" {
			continue
		}
		if field.Type.Implements(textMarshalerType) ||
			field.Type.Implements(textUnmarshalerType) {
			// skip
		} else {
			field.Type = tomlOnlyType(field.Type)
		}
		newFields = append(newFields, field)
	}
	return reflect.StructOf(newFields)
}

var textMarshalerType = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()

var textUnmarshalerType = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
