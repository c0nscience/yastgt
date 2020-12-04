package transform

import (
	"math"

	"gonum.org/v1/gonum/mat"

	"github.com/c0nscience/yastgt/pkg/unit"
)

func New(a, b, c, d, e, f float64) *mat.Dense {
	return mat.NewDense(3, 3, []float64{
		a, c, e,
		b, d, f,
		0, 0, 1,
	})
}

func NewTranslation(tx, ty float64) *mat.Dense {
	return New(1, 0, 0, 1, tx, ty)
}

func NewScaling(sx, sy float64) *mat.Dense {
	return New(sx, 0, 0, sy, 0, 0)
}

func NewRotation(a float64) *mat.Dense {
	rad := unit.DegToRad(a)
	return New(math.Cos(rad), math.Sin(rad), -math.Sin(rad), math.Cos(rad), 0, 0)
}

func NewSkewX(a float64) *mat.Dense {
	rad := unit.DegToRad(a)
	return New(1, 0, math.Tan(rad), 1, 0, 0)
}

func NewSkewY(a float64) *mat.Dense {
	rad := unit.DegToRad(a)
	return New(1, math.Tan(rad), 0, 1, 0, 0)
}
