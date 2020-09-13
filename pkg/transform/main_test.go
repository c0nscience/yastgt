package transform_test

import (
	"testing"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/transform"
	"github.com/c0nscience/yastgt/pkg/transform/gcode"
	"github.com/stretchr/testify/assert"
)

func Test_Transform(t *testing.T) {
	// given
	s := svg.SVG{
		Path: []svg.Path{
			{
				Points: []svg.PointI{
					svg.Point{X: 0.0, Y: 0.0},
					svg.Point{X: 10.0, Y: 0.0},
					svg.Point{X: 10.0, Y: 10.0},
					svg.Point{X: 0.0, Y: 10.0},
					svg.Point{X: 0.0, Y: 0.0},
					svg.CubicPoint{
						P1: svg.Point{X: 1, Y: 4},
						P2: svg.Point{X: 3, Y: 0},
						CP: svg.Point{X: 2, Y: 2},
					},
				},
			},
			{
				Points: []svg.PointI{
					svg.Point{X: 10, Y: 10},
					svg.Point{X: 5, Y: 2, Rel: true},
					svg.Point{X: 7, Y: 8, Rel: true},
					svg.Point{X: 2, Y: 2, Rel: true},
					svg.Point{X: 14, Y: 10, Rel: true},
					svg.Point{X: 42, Y: 21},
				},
			},
		},
	}
	exp := []gcode.Cmd{
		gcode.Cmd("G21"),
		gcode.Cmd("G90"),
		gcode.Cmd("G17"),
		gcode.Cmd("M400"),
		gcode.Cmd("M280 P0 S150"),
		gcode.Cmd("M400"),
		gcode.Cmd("G28 XY"),
		gcode.Cmd("G0 F5000"),
		gcode.Cmd("G0 X0.0 Y0.0"),
		gcode.Cmd("M400"),
		gcode.Cmd("M280 P0 S30"),
		gcode.Cmd("M400"),
		gcode.Cmd("G0 X10.0 Y0.0"),
		gcode.Cmd("G0 X10.0 Y10.0"),
		gcode.Cmd("G0 X0.0 Y10.0"),
		gcode.Cmd("G0 X0.0 Y0.0"),
		gcode.Cmd("G5 I-1.0 J2.0 P1.0 Q-2.0 X2.0 Y2.0"),
		gcode.Cmd("M400"),
		gcode.Cmd("M280 P0 S150"),
		gcode.Cmd("M400"),
		gcode.Cmd("G0 X10.0 Y10.0"),
		gcode.Cmd("M400"),
		gcode.Cmd("M280 P0 S30"),
		gcode.Cmd("M400"),
		gcode.Cmd("G91"),
		gcode.Cmd("G0 X5.0 Y2.0"),
		gcode.Cmd("G0 X7.0 Y8.0"),
		gcode.Cmd("G0 X2.0 Y2.0"),
		gcode.Cmd("G0 X14.0 Y10.0"),
		gcode.Cmd("G90"),
		gcode.Cmd("G0 X42.0 Y21.0"),
		gcode.Cmd("M400"),
		gcode.Cmd("M280 P0 S150"),
		gcode.Cmd("M400"),
		gcode.Cmd("G28 XY"),
	}

	// when
	cmd := transform.Gcode(s)

	// then
	assert.Equal(t, exp, cmd)
}
