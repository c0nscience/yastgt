package parser

import (
	"encoding/xml"

	"github.com/c0nscience/yastgt/pkg/svg"
)

func SVGfrom(data string) (*svg.SVG, error) {
	res := &svg.SVG{}
	err := xml.Unmarshal([]byte(data), res)

	if err != nil {
		return nil, err
	}

	res.Populate()
	return res, nil
}

// type svgXml struct {
// 	G []*gXml `xml:"g"`
// }

// type gXml struct {
// 	Path []*pathXml `xml:"path"`
// }

// type pathXml struct {
// 	D string `xml:"d,attr"`
// }
