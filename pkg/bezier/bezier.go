package bezier

import (
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"math"
)

/**
Code was translated from https://www.lemoda.net/maths/bezier-length/index.html into go
*/

func point(t, start, c1, c2, end float64) float64 {
	return start*(1.0-t)*(1.0-t)*(1.0-t) +
		3.0*c1*(1.0-t)*(1.0-t)*t +
		3.0*c2*(1.0-t)*t*t +
		end*t*t*t
}

func Length(start svg.Point, c svg.CubicPoint) float64 {
	var p svg.Point
	var prevP svg.Point
	steps := 10
	length := 0.0
	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)
		p.X = point(t, start.X, c.P1.X,
			c.P2.X, c.CP.X)
		p.Y = point(t, start.Y, c.P1.Y,
			c.P2.Y, c.CP.Y)
		if i > 0 {
			xDiff := p.X - prevP.X
			yDiff := p.Y - prevP.Y
			length += math.Sqrt(xDiff*xDiff + yDiff*yDiff)
		}
		prevP = p
	}
	return length
}
