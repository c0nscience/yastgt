package svg

type Point struct {
	X, Y float64
	Rel  bool
}

func (me Point) RelativeTo(p Point) Point {
	return Point{
		X: me.X - p.X,
		Y: me.Y - p.Y,
	}
}

type CubicPoint struct {
	CP  Point
	P1  Point
	P2  Point
	Rel bool
}
