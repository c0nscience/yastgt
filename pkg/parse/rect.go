package parse

import (
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func Rect(r xml.Rect) []svg.PointI {
	return []svg.PointI{
		&svg.Point{X: r.X, Y: r.Y, MoveTo: true},
		&svg.Point{X: r.X + r.Width, Y: r.Y},
		&svg.Point{X: r.X + r.Width, Y: r.Y + r.Height},
		&svg.Point{X: r.X, Y: r.Y + r.Height},
		&svg.Point{X: r.X, Y: r.Y},
	}
}
