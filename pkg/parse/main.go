package parse

import (
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func ToSvg(xml xml.SVG) svg.SVG {
	res := svg.SVG{
		Path: fromPath(xml.Path),
		G:    []svg.G{},
	}

	for _, xg := range xml.G {
		g := svg.G{
			Path: fromPath(xg.Path),
		}
		res.G = append(res.G, g)
	}

	return res
}

func fromPath(pths []xml.Path) []svg.Path {
	res := []svg.Path{}
	for _, p := range pths {
		path := svg.Path{
			M: M(FindM(p.D)),
			L: L(FindL(p.D)),
		}

		res = append(res, path)
	}
	return res
}
