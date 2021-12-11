package fz

import (
	"bytes"
	"encoding/xml"
	"io"
	"reflect"
	"sort"
	"strings"
)

type NamedConfigItem struct {
	Name  string
	Value any
}

func (_ ConfigScope) NamedConfigItems(
	items ConfigItems,
) (
	nameds []NamedConfigItem,
	m map[string]NamedConfigItem,
) {
	m = make(map[string]NamedConfigItem)
	for _, item := range items {
		name := reflect.TypeOf(item).Name()
		item := NamedConfigItem{
			Name:  name,
			Value: item,
		}
		nameds = append(nameds, item)
		m[name] = item
	}
	sort.Slice(nameds, func(i, j int) bool {
		a := nameds[i]
		b := nameds[j]
		if w1, w2 := configWeights[a.Name], configWeights[b.Name]; w1 != w2 {
			return w1 < w2
		}
		return a.Name < b.Name
	})
	return
}

type WriteConfig func(w io.Writer) error

func (_ ConfigScope) WriteConfig(
	nameds []NamedConfigItem,
) WriteConfig {

	return func(w io.Writer) (err error) {
		defer he(&err)

		encoder := xml.NewEncoder(w)
		encoder.Indent("", "    ")
		for _, named := range nameds {
			ce(encoder.EncodeElement(named.Value, xml.StartElement{
				Name: xml.Name{
					Local: named.Name,
				},
			}))
		}

		return
	}
}

var configWeights = map[string]int{
	"MainAction": 1,
}

type ReadConfig func(r io.Reader) ([]any, error)

func (_ ConfigScope) ReadConfig(
	nameds map[string]NamedConfigItem,
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

			item, ok := nameds[data.XMLName.Local]
			if !ok {
				// unknown config key
				continue
			}
			ptr := reflect.New(reflect.TypeOf(item.Value))
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
