package fz

import "encoding/xml"

func nextTokenSkipCharData(d *xml.Decoder) (xml.Token, error) {
read:
	token, err := d.Token()
	if err != nil {
		return nil, we(err)
	}
	if _, ok := token.(xml.CharData); ok {
		goto read
	}
	return token, nil
}
