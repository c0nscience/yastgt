package reader

import (
	goxml "encoding/xml"

	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func Unmarshal(b []byte) (xml.SVG, error) {
	res := xml.SVG{}
	err := goxml.Unmarshal(b, &res)
	if err != nil {
		return xml.SVG{}, err
	}
	return res, nil
}
