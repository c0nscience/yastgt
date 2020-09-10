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

func ToGcode(svg svg.SVG) []gcode.Cmd {
	res := append([]gcode.Cmd{}, head()...)

	for _, g := range svg.G {
		res = append(res, fromPath(g.Path)...)
	}
	res = append(res, fromPath(svg.Path)...)

	res = append(res, gcode.G28("XY"))
	return res
}

func fromPath(pths []svg.Path) []gcode.Cmd {
	res := []gcode.Cmd{}
	for _, pth := range pths {
		mp := pth.M
		if len(pth.M) >= 1 {
			res = append(res, gcode.G0(mp[0]))
			mp = mp[1:]
		}

		res = append(res, penDwn...)

		for _, p := range mp {
			res = append(res, gcode.G0(p))
		}

		for _, p := range pth.L {
			res = append(res, gcode.G0(p))
		}

		res = append(res, penUp...)
	}

	return res
}
