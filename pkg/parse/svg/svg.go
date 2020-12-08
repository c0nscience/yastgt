package svg

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type SVG struct {
	Height float64
	Points []PointI
}

type PointI interface {
	CurrPt() *Point
	ToPlotterCoord(h float64)
	Transform(t *mat.Dense)
}

var _ PointI = &Point{}

type Point struct {
	X, Y   float64
	MoveTo bool
	Idx    int
}

func (me *Point) Transform(m *mat.Dense) {
	v := mat.NewDense(3, 1, []float64{
		me.X,
		me.Y,
		1,
	})

	var res mat.Dense
	res.Mul(m, v)
	me.X = res.At(0, 0)
	me.Y = res.At(1, 0)
}

func (me *Point) ToPlotterCoord(h float64) {
	me.Y = h - me.Y
}

func (me *Point) CurrPt() *Point {
	return me
}

func (me *Point) RelativeTo(p *Point) Point {
	return Point{
		X: me.X - p.X,
		Y: me.Y - p.Y,
	}
}

func (me Point) String() string {
	return fmt.Sprintf("(%.2f,%.2f)", me.X, me.Y)
}

var _ PointI = &CubicPoint{}

type CubicPoint struct {
	CP, P1, P2 *Point
}

func (me *CubicPoint) Transform(m *mat.Dense) {
	me.CP.Transform(m)
	me.P1.Transform(m)
	me.P2.Transform(m)
}

func (me *CubicPoint) CurrPt() *Point {
	return me.CP
}

func (me *CubicPoint) ToPlotterCoord(h float64) {
	me.CP.Y = h - me.CP.Y
	me.P1.Y = h - me.P1.Y
	me.P2.Y = h - me.P2.Y
}

func (me CubicPoint) String() string {
	return fmt.Sprintf("(%s, %s, %s)", me.CP, me.P1, me.P2)
}

var _ PointI = &CirclePoint{}

type CirclePoint struct {
	CP, P *Point
	R     float64
}

func (c *CirclePoint) CurrPt() *Point {
	return c.P
}

func (c *CirclePoint) ToPlotterCoord(h float64) {
	c.P.ToPlotterCoord(h)
	c.CP.ToPlotterCoord(h)
}

func (c *CirclePoint) Transform(t *mat.Dense) {
	c.P.Transform(t)
	c.CP.Transform(t)
}
