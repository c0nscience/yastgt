package parse

import (
	"strconv"

	"gonum.org/v1/gonum/mat"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/parse/transform"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func SVG(xml xml.SVG) svg.SVG {
	h, _ := strconv.ParseFloat(xml.Height[:len(xml.Height)-2], 64)
	res := svg.SVG{
		Height: h,
		Path:   fromPath(xml.Path, []*mat.Dense{}),
	}

	res.Path = append(res.Path, fromGroup(xml.G, []*mat.Dense{})...)

	return res
}

func fromGroup(grps []xml.G, trans []*mat.Dense) []svg.Path {
	res := []svg.Path{}
	for _, g := range grps {
		trans = append(trans, transform.ParseTypes(g.Transform)...)
		res = append(res, fromGroup(g.G, trans)...)
	}

	for _, g := range grps {
		res = append(res, fromPath(g.Path, trans)...)
	}

	return res
}

func fromPath(pths []xml.Path, trans []*mat.Dense) []svg.Path {
	res := []svg.Path{}
	for _, p := range pths {
		trans = append(trans, transform.ParseTypes(p.Transform)...)
		res = append(res, Path(p))
	}
	return transform.Apply(trans, res)
}
