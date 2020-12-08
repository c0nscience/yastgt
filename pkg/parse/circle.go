package parse

import (
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

/**
Point on a circle calculation:
x = a + r * cos(t)
y = b + r * sin(t)
with t as [0,2*pi] determining the angle of the line from the middle to the point on the circle
*/
func Circle(c xml.Circle) svg.PointI {
	return &svg.CirclePoint{
		P: &svg.Point{
			X: c.CX,
			Y: c.CY,
		},
		CP:&svg.Point{
			X:      c.CX + c.R,
			Y:      c.CY,
			MoveTo: true,
		},
		R: c.R,
	}
}
