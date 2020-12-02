package parse

import (
	"strconv"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func SVG(xml xml.SVG) svg.SVG {
	h, _ := strconv.ParseFloat(xml.Height[:len(xml.Height)-2], 64)
	res := svg.SVG{
		Height: h,
		Path:   fromPath(xml.Path),
	}

	for _, g := range xml.G {
		res.Path = append(res.Path, fromPath(g.Path)...)
	}

	return res
}

func fromPath(pths []xml.Path) []svg.Path {
	res := []svg.Path{}
	for _, p := range pths {
		res = append(res, Path(p))
	}
	return res
}
