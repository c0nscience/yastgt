package parse

import (
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func Line(l xml.Line) []svg.PointI {
	return []svg.PointI{
		&svg.Point{
			X:      l.X1,
			Y:      l.Y1,
			MoveTo: true,
		},
		&svg.Point{
			X: l.X2,
			Y: l.Y2,
		},
	}
}
