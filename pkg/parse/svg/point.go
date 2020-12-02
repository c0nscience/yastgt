package svg

import "fmt"

type PointI interface {
	fmt.Stringer
	CurrPt() Point
	ToPlotterCoord(h float64) PointI
}

var _ PointI = Point{}

type Point struct {
	X, Y   float64
	Rel    bool
	MoveTo bool
	Idx    int
}

func (me Point) ToPlotterCoord(h float64) PointI {
	me.Y = h - me.Y
	return me
}

func (me Point) String() string {
	return fmt.Sprintf("(%.2f,%.2f)", me.X, me.Y)
}

func (me Point) CurrPt() Point {
	return me
}

func (me Point) RelativeTo(p Point) Point {
	return Point{
		X: me.X - p.X,
		Y: me.Y - p.Y,
	}
}

var _ PointI = CubicPoint{}

type CubicPoint struct {
	CP  Point
	P1  Point
	P2  Point
	Rel bool
}

func (me CubicPoint) String() string {
	return fmt.Sprintf("(%s, %s, %s)", me.CP, me.P1, me.P2)
}

func (me CubicPoint) CurrPt() Point {
	return me.CP
}

func (me CubicPoint) ToPlotterCoord(h float64) PointI {
	me.CP.Y = h - me.CP.Y
	me.P1.Y = h - me.P1.Y
	me.P2.Y = h - me.P2.Y
	return me
}
