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
		Path: []*svg.Path{
			&svg.Path{
				M: []svg.Point{
					svg.Point{X: 0.0, Y: 0.0},
				},
				L: []svg.Point{
					svg.Point{X: 10.0, Y: 0.0},
					svg.Point{X: 10.0, Y: 10.0},
					svg.Point{X: 0.0, Y: 10.0},
					svg.Point{X: 0.0, Y: 0.0},
				},
			},
		},
	}

	// when
	cmd := transform.ToGcode(s)

	// then
	assert.Len(t, cmd, 19)
	assert.Equal(t, gcode.Cmd("G21"), cmd[0])
	assert.Equal(t, gcode.Cmd("G90"), cmd[1])
	assert.Equal(t, gcode.Cmd("G17"), cmd[2])
	assert.Equal(t, gcode.Cmd("M400"), cmd[3])
	assert.Equal(t, gcode.Cmd("M280 P0 S150"), cmd[4])
	assert.Equal(t, gcode.Cmd("M400"), cmd[5])
	assert.Equal(t, gcode.Cmd("G28 XY"), cmd[6])
	assert.Equal(t, gcode.Cmd("G0 F5000"), cmd[7])
	assert.Equal(t, gcode.Cmd("G0 X0.0 Y0.0"), cmd[8])
	assert.Equal(t, gcode.Cmd("M400"), cmd[9])
	assert.Equal(t, gcode.Cmd("M280 P0 S30"), cmd[10])
	assert.Equal(t, gcode.Cmd("M400"), cmd[11])
	assert.Equal(t, gcode.Cmd("G0 X10.0 Y0.0"), cmd[12])
	assert.Equal(t, gcode.Cmd("G0 X10.0 Y10.0"), cmd[13])
	assert.Equal(t, gcode.Cmd("G0 X0.0 Y10.0"), cmd[14])
	assert.Equal(t, gcode.Cmd("G0 X0.0 Y0.0"), cmd[15])
	assert.Equal(t, gcode.Cmd("M400"), cmd[16])
	assert.Equal(t, gcode.Cmd("M280 P0 S150"), cmd[17])
	assert.Equal(t, gcode.Cmd("M400"), cmd[18])
}
