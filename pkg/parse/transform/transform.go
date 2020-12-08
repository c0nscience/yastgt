package transform

import (
	"gonum.org/v1/gonum/mat"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
)

func Apply(trans []*mat.Dense, pts []svg.PointI) {
	if len(trans) == 0 {
		return
	}

	t := fold(trans)
	for _, s := range pts {
		s.Transform(t)
	}
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
