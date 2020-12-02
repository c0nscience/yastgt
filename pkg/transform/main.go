package transform

import (
	"github.com/c0nscience/yastgt/pkg/bezier"
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/transform/gcode"
)

var penUp = gcode.Servo("130")
var penDwn = gcode.Servo("90")
var g0Speed float64 = 2000
var g5Speed float64 = 100

func SetG0Speed(i float64) {
	g0Speed = i
}

func SetG5Speed(i float64) {
	g5Speed = i
}

func Gcode(svg svg.SVG) []gcode.Cmd {
	res := append([]gcode.Cmd{}, head()...)

	res = append(res, fromPath(svg.Path, svg.Height)...)

	res = append(res, penUp...)

	res = append(res, home()...)
	res = append(res, gcode.M18)
	return res
}

func head() []gcode.Cmd {
	res := []gcode.Cmd{gcode.G21, gcode.G90, gcode.G17}

	res = append(res, penUp...)
	res = append(res, home()...)

	return res
}

func fromPath(pths []svg.Path, h float64) []gcode.Cmd {
	res := []gcode.Cmd{}
	var cp svg.PointI
	for _, pth := range pths {
		for _, p := range pth.Points {
			p = p.ToPlotterCoord(h)
			switch pt := p.(type) {
			case svg.Point:
				if pt.MoveTo {
					res = append(res, penUp...)
				}

				res = append(res, gcode.G0(pt, g0Speed))

				if pt.MoveTo {
					res = append(res, penDwn...)
				}
			case svg.CubicPoint:
				length := bezier.Length(cp.CurrPt(), pt)
				res = append(res, gcode.G5(pt, cp.CurrPt(), g5Speed/length))
			}
			cp = p
		}
	}

	return res
}

func home() []gcode.Cmd {
	res := []gcode.Cmd{}

	res = append(res, "M906 X200")
	res = append(res, "M906 Y200")
	res = append(res, "M906 I1 Y200")
	res = append(res, gcode.G28("XY"))
	res = append(res, "M906 X800")
	res = append(res, "M906 Y800")
	res = append(res, "M906 I1 Y800")

	return res
}
