package transform

import (
	"gonum.org/v1/gonum/mat"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
)

func Apply(trans []*mat.Dense, paths []svg.Path) []svg.Path {
	if len(trans) == 0 {
		return paths
	}

	t := fold(trans)
	res := []svg.Path{}
	for _, p := range paths {
		pts := []svg.PointI{}
		for _, pt := range p.Points {
			pts = append(pts, pt.Transform(t))
		}
		res = append(res, svg.Path{Points: pts})
	}
	return res
}

func fold(trans []*mat.Dense) *mat.Dense {
	res := trans[0]
	for _, t := range trans[1:] {
		res = mul(res, t)
	}

	return res
}

func mul(m1 *mat.Dense, m2 *mat.Dense) *mat.Dense {
	var res mat.Dense
	res.Mul(m1, m2)
	return &res
}
