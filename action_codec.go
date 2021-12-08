package fz

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func unmarshalAction(d *xml.Decoder, target *Action) (err error) {
	var data struct {
		XMLName xml.Name
		Raw     []byte `xml:",innerxml"`
	}
	if err := d.Decode(&data); err != nil {
		return we(err)
	}

	switch data.XMLName.Local {

	case "SequentialAction":
		var action SequentialAction
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
		ce(valueDecoder.Decode(&action))
		*target = action

	case "ParallelAction":
		var action ParallelAction
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
		ce(valueDecoder.Decode(&action))
		*target = action

	default:
		panic(fmt.Errorf("unknown action: %s", data.XMLName.Local))
	}

	return
}
