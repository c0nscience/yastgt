package transform

import (
	"fmt"

	"github.com/c0nscience/yastgt/pkg/bezier"
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/transform/gcode"
)

var g0Speed float64 = 2000
var g5Speed float64 = 100
var penUpAngle int64 = 130
var penDownAngle int64 = 130

func SetG0Speed(i float64) {
	g0Speed = i
}

func SetG5Speed(i float64) {
	g5Speed = i
}

func SetPenUpAngle(f int64) {
	penUpAngle = f
}

func SetPenDownAngle(f int64) {
	penDownAngle = f
}

func penUp() []gcode.Cmd {
	return gcode.Servo(fmt.Sprintf("%d", penUpAngle))
}

func penDown() []gcode.Cmd {
	return gcode.Servo(fmt.Sprintf("%d", penDownAngle))
}

func Gcode(svg svg.SVG) []gcode.Cmd {
	res := append([]gcode.Cmd{}, head()...)

	res = append(res, fromPoints(svg.Points, svg.Height)...)

	res = append(res, penUp()...)

	res = append(res, home()...)
	res = append(res, gcode.M18)
	return res
}

func head() []gcode.Cmd {
	res := []gcode.Cmd{gcode.G21, gcode.G90, gcode.G17}

	res = append(res, penUp()...)
	res = append(res, home()...)

	return res
}

func fromPoints(pts []svg.PointI, h float64) []gcode.Cmd {
	res := []gcode.Cmd{}
	var cp svg.PointI
	for _, p := range pts {
		p.ToPlotterCoord(h)
		switch pt := p.(type) {
		case *svg.Point:
			if pt.MoveTo {
				res = append(res, penUp()...)
			}

			res = append(res, gcode.G0(pt, g0Speed))

			if pt.MoveTo {
				res = append(res, penDown()...)
			}
		case *svg.CubicPoint:
			length := bezier.Length(cp.CurrPt(), pt)
			res = append(res, gcode.G5(pt, cp.CurrPt(), g5Speed/length))
		case *svg.CirclePoint:
			res = append(res, penUp()...)
			res = append(res, gcode.G0(pt.CP, g0Speed))
			res = append(res, penDown()...)
			res = append(res, gcode.G2(pt))
		}
		cp = p
	}

	return res
}

func home() []gcode.Cmd {
	return []gcode.Cmd{
		"M906 X200",
		"M906 Y200",
		"M906 I1 Y200",
		gcode.G28("XY"),
		"M906 X800",
		"M906 Y800",
		"M906 I1 Y800",
	}

}
