package transform

import (
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/transform/gcode"
)

var penUp = gcode.Servo("150")
var penDwn = gcode.Servo("30")

func head() []gcode.Cmd {
	res := []gcode.Cmd{gcode.G21, gcode.G90, gcode.G17}

	res = append(res, penUp...)
	res = append(res, gcode.G28("XY"))
	res = append(res, gcode.G0F(5000))

	return res
}

func Gcode(svg svg.SVG) []gcode.Cmd {
	res := append([]gcode.Cmd{}, head()...)

	res = append(res, fromPath(svg.Path)...)

	res = append(res, gcode.G28("XY"))
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

func fromPath(pths []svg.Path) []gcode.Cmd {
	res := []gcode.Cmd{}
	for _, pth := range pths {
		pts := pth.Points
		if len(pts) >= 1 {
			pt := pts[0].CurrPt()
			detMode(pt, &res)

			res = append(res, gcode.G0(pt))
			pts = pts[1:]
		}

		res = append(res, penDwn...)

		for _, p := range pts {
			detMode(p, &res)
			switch pt := p.(type) {
			case svg.Point:
				res = append(res, gcode.G0(pt))
			case svg.CubicPoint:
				res = append(res, gcode.G5(pt))
			}
		}

		res = append(res, penUp...)
	}

	return res
}
