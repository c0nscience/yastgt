package transform

import (
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/transform/gcode"
)

var penUp = gcode.Servo("150")
var penDwn = gcode.Servo("30")
var g0Speed int64 = 2000
var g5Speed int64 = 100

func SetG0Speed(i int64) {
	g0Speed = i
}

func SetG5Speed(i int64) {
	g5Speed = i
}

func Gcode(svg svg.SVG) []gcode.Cmd {
	res := append([]gcode.Cmd{}, head()...)

	res = append(res, fromPath(svg.Path)...)

	res = append(res, penUp...)

	res = append(res, home()...)
	return res
}

func head() []gcode.Cmd {
	res := []gcode.Cmd{gcode.G21, gcode.G90, gcode.G17}

	res = append(res, penUp...)
	res = append(res, home()...)

	return res
}

func fromPath(pths []svg.Path) []gcode.Cmd {
	res := []gcode.Cmd{}
	var cp svg.PointI
	for _, pth := range pths {
		for _, p := range pth.Points {
			detMode(p, &res)
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
				res = append(res, gcode.G5(pt, cp.CurrPt(), g5Speed))
			}
			cp = p
		}
	}

	return res
}

var abs = true

func detMode(p svg.PointI, res *[]gcode.Cmd) {
	if abs && p.Relative() {
		*res = append(*res, gcode.G91)
		abs = false
	}
	if !abs && !p.Relative() {
		*res = append(*res, gcode.G90)
		abs = true
	}
}

func home() []gcode.Cmd {
	res := []gcode.Cmd{}

	res = append(res, gcode.Cmd("M906 X200"))
	res = append(res, gcode.Cmd("M906 Y200"))
	res = append(res, gcode.Cmd("M906 I1 Y200"))
	res = append(res, gcode.G28("XY"))
	res = append(res, gcode.Cmd("M906 X800"))
	res = append(res, gcode.Cmd("M906 Y800"))
	res = append(res, gcode.Cmd("M906 I1 Y800"))

	return res
}
