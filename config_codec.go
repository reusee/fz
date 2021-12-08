package fz

import (
	"bytes"
	"encoding/xml"
	"io"
	"reflect"
	"sort"
	"strings"
)

type WriteConfig func(w io.Writer) error

func (_ ConfigScope) WriteConfig(
	m ConfigMap,
) WriteConfig {
	return func(w io.Writer) (err error) {
		defer he(&err)

		var keys []string
		for key := range m {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		encoder := xml.NewEncoder(w)
		encoder.Indent("", "    ")
		for _, key := range keys {
			value := m[key]
			ce(encoder.EncodeElement(value, xml.StartElement{
				Name: xml.Name{
					Local: key,
				},
			}))
		}

		return
	}
}

type ReadConfig func(r io.Reader) ([]any, error)

func (_ ConfigScope) ReadConfig(
	m ConfigMap,
) ReadConfig {
	return func(r io.Reader) (decls []any, err error) {
		defer he(&err)

		for {
			var data struct {
				XMLName xml.Name
				Raw     []byte `xml:",innerxml"`
			}
			decoder := xml.NewDecoder(r)
			err = decoder.Decode(&data)
			if err == io.EOF {
				err = nil
				break
			}
			ce(err)

			value, ok := m[data.XMLName.Local]
			if !ok {
				// unknown config key
				continue
			}
			ptr := reflect.New(reflect.TypeOf(value))
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
			decls = append(decls, ptr.Interface())
		}

		return
	}
}
