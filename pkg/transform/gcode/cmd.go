package gcode

import (
	"fmt"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
)

type Cmd string

const (
	G21  Cmd = "G21"
	G90  Cmd = "G90"
	G17  Cmd = "G17"
	M400 Cmd = "M400"
	M18  Cmd = "M18"
)

func G28(s string) Cmd {
	return Cmd("G28 " + s)
}

func M280(s string) Cmd {
	return Cmd("M280 P0 S" + s)
}

func G0(p svg.Point, i float64) Cmd {
	x := floatToString(p.X)
	y := floatToString(p.Y)
	return Cmd(fmt.Sprintf("G0 F%.2f X%s Y%s", i, x, y))
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

func G5(p svg.CubicPoint, cp svg.Point, i float64) Cmd {
	x := floatToString(p.CP.X)
	y := floatToString(p.CP.Y)
	p1 := p.P1.RelativeTo(cp)
	p2 := p.P2.RelativeTo(p.CP)
	p1x := floatToString(p1.X)
	p1y := floatToString(p1.Y)
	p2x := floatToString(p2.X)
	p2y := floatToString(p2.Y)
	return Cmd(fmt.Sprintf("G5 F%.2f I%s J%s P%s Q%s X%s Y%s", i, p1x, p1y, p2x, p2y, x, y))
}
