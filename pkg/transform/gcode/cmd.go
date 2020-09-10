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
)

func G28(s string) Cmd {
	return Cmd("G28 " + s)
}

func M280(s string) Cmd {
	return Cmd("M280 P0 S" + s)
}

func G0(p svg.Point) Cmd {
	x := fmt.Sprintf("%.1f", p.X)
	y := fmt.Sprintf("%.1f", p.Y)
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
