package svg

type PointI interface {
	CurrPt() Point
	Relative() bool
}

var _ PointI = CubicPoint{}

type Point struct {
	X, Y   float64
	Rel    bool
	MoveTo bool
}

func (me Point) CurrPt() Point {
	return me
}

func (me Point) Relative() bool {
	return me.Rel
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

func (me CubicPoint) CurrPt() Point {
	return me.CP
}

func (me CubicPoint) Relative() bool {
	return me.Rel
}
