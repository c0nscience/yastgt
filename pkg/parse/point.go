package parse

import (
	"strconv"
	"strings"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
)

func Point(s string) svg.Point {
	crds := strings.Split(s, ",")
	x, _ := strconv.ParseFloat(crds[0], 64)
	y, _ := strconv.ParseFloat(crds[1], 64)
	return svg.Point{X: x, Y: y}
}

func MoveTo(s string) svg.Point {
	p := Point(s)
	p.MoveTo = true
	return p
}

func RelPoint(s string) svg.Point {
	p := Point(s)
	p.Rel = true
	return p
}

func RelMoveTo(s string) svg.Point {
	p := RelPoint(s)
	p.MoveTo = true
	return p
}
