package fz

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/reusee/e4"
	"github.com/reusee/sb"
)

func TestActionCodec(t *testing.T) {
	defer he(nil, e4.TestingFatal(t))

	type Case struct {
		Value Action
		XML   string
	}

	cases := []Case{
		{
			Value: Seq(),
			XML:   "<SequentialAction></SequentialAction>",
		},
		{
			Value: Seq(Seq()),
			XML:   "<SequentialAction><SequentialAction></SequentialAction></SequentialAction>",
		},
		{
			Value: Seq(Seq(Seq())),
			XML:   "<SequentialAction><SequentialAction><SequentialAction></SequentialAction></SequentialAction></SequentialAction>",
		},
		{
			Value: Seq(Seq(Seq()), Par()),
			XML:   "<SequentialAction><SequentialAction><SequentialAction></SequentialAction></SequentialAction><ParallelAction></ParallelAction></SequentialAction>",
		},
	}

	for i, c := range cases {
		// encode
		buf := new(bytes.Buffer)
		ce(xml.NewEncoder(buf).Encode(c.Value))
		if !bytes.Equal(buf.Bytes(), []byte(c.XML)) {
			t.Fatalf("bad case: %d", i)
		}
		// decode
		var action Action
		ce(unmarshalAction(
			xml.NewDecoder(strings.NewReader(c.XML)),
			&action,
		))
		if sb.MustCompare(
			sb.Marshal(action),
			sb.Marshal(c.Value),
		) != 0 {
			t.Fatalf("bad case: %d", i)
		}
	}

}
