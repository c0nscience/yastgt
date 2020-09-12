package gcode

import (
	"fmt"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
)

type Cmd string

const (
	G21  Cmd = "G21"
	G90  Cmd = "G90"
	G91  Cmd = "G91"
	G17  Cmd = "G17"
	M400 Cmd = "M400"
)

func G28(s string) Cmd {
	return Cmd("G28 " + s)
}

func M280(s string) Cmd {
	return Cmd("M280 P0 S" + s)
}

func G0(p svg.Point) Cmd {
	x := floatToString(p.X)
	y := floatToString(p.Y)
	return Cmd("G0 X" + x + " Y" + y)
}

func G0F(f int) Cmd {
	return Cmd("G0 F" + fmt.Sprintf("%d", f))
}

func Servo(s string) []Cmd {
	return []Cmd{
		M400,
		M280(s),
		M400,
	}
}

func floatToString(f float64) string {
	return fmt.Sprintf("%.1f", f)
}

func G5(p svg.CubicPoint) Cmd {
	x := floatToString(p.CP.X)
	y := floatToString(p.CP.Y)
	p1x := floatToString(p.P1.X)
	p1y := floatToString(p.P1.Y)
	p2x := floatToString(p.P2.X)
	p2y := floatToString(p.P2.Y)
	if !p.Rel {
		p1 := p.P1.RelativeTo(p.CP)
		p2 := p.P2.RelativeTo(p.CP)
		p1x = floatToString(p1.X)
		p1y = floatToString(p1.Y)
		p2x = floatToString(p2.X)
		p2y = floatToString(p2.Y)
	}
	return Cmd(fmt.Sprintf("G5 I%s J%s P%s Q%s X%s Y%s", p1x, p1y, p2x, p2y, x, y))
}
