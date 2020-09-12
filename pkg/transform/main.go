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

var abs = true

func Gcode(svg svg.SVG) []gcode.Cmd {
	res := append([]gcode.Cmd{}, head()...)

	res = append(res, fromPath(svg.Path)...)

	res = append(res, gcode.G28("XY"))
	return res
}

func fromPath(pths []svg.Path) []gcode.Cmd {
	res := []gcode.Cmd{}
	for _, pth := range pths {
		pts := pth.Points
		if len(pts) >= 1 {
			pt, ok := pts[0].(svg.Point)
			if !ok {
				pt = pts[0].(svg.CubicPoint).CP
			}
			res = append(res, gcode.G0(pt))
			pts = pts[1:]
		}

		res = append(res, penDwn...)

		for _, p := range pts {
			switch pt := p.(type) {
			case svg.Point:
				if abs && pt.Rel {
					res = append(res, gcode.G91)
					abs = false
				}
				if !abs && !pt.Rel {
					res = append(res, gcode.G90)
					abs = true
				}
				res = append(res, gcode.G0(pt))
			case svg.CubicPoint:
				if pt.Rel && abs {
					res = append(res, gcode.G91)
				}
				if !pt.Rel && !abs {
					res = append(res, gcode.G90)
				}
				res = append(res, gcode.G5(pt))
			}
		}

		res = append(res, penUp...)
	}

	return res
}
