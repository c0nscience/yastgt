package parse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func Test_Path(t *testing.T) {
	t.Run("should parse from xml", func(t *testing.T) {
		d := "M 1.5420259,10.163793 L 14,20 25,12 M 10.4375,31.906968 H 6 V 18 Z C 10,12 30,60 22,50 Z m -2,0 l -5,-1 h 5 v 7 c 0,0 5,-2 4,2"
		x := xml.Path{
			D: d,
		}
		subj := parse.Path(x)

		assert.Equal(t, []svg.PointI{
			&svg.Point{X: 1.5420259, Y: 10.163793, MoveTo: true}, // move to
			&svg.Point{X: 14, Y: 20},                             // lineto
			&svg.Point{X: 25, Y: 12},                             // lineto
			&svg.Point{X: 10.4375, Y: 31.906968, MoveTo: true},   // moveto
			&svg.Point{X: 6, Y: 31.906968},                       // horizontal lineto
			&svg.Point{X: 6, Y: 18},                              // vertical lineto
			&svg.Point{X: 10.4375, Y: 31.906968},                 // close path
			&svg.CubicPoint{ // bezier curve
				CP: &svg.Point{X: 22, Y: 50},
				P1: &svg.Point{X: 10, Y: 12},
				P2: &svg.Point{X: 30, Y: 60},
			},
			&svg.Point{X: 10.4375, Y: 31.906968}, // close path TODO this is not the correct point
		}, subj)

	})

	t.Run("should parse exponential values", func(t *testing.T) {
		d := "C 7e-4,4 6,7 3,5 4,5 7,5 1,5 L 1,2 5,2"

		x := xml.Path{
			D: d,
		}

		subj := parse.Path(x)

		assert.Equal(t, []svg.PointI{
			&svg.CubicPoint{
				CP: &svg.Point{X: 3, Y: 5},
				P1: &svg.Point{X: 0.0007, Y: 4},
				P2: &svg.Point{X: 6, Y: 7},
			},
			&svg.CubicPoint{
				CP: &svg.Point{X: 1, Y: 5},
				P1: &svg.Point{X: 4, Y: 5},
				P2: &svg.Point{X: 7, Y: 5},
			},
			&svg.Point{X: 1, Y: 2},
			&svg.Point{X: 5, Y: 2},
		}, subj)
	})

	t.Run("should handle consecutive move commands as draw moves", func(t *testing.T) {
		// given
		d := "M 10,20 20,20 30,30 40,30"
		x := xml.Path{
			D: d,
		}
		// when
		subj := parse.Path(x)

		// then
		assert.Equal(t, []svg.PointI{
			&svg.Point{X: 10, Y: 20, MoveTo: true},
			&svg.Point{X: 20, Y: 20},
			&svg.Point{X: 30, Y: 30},
			&svg.Point{X: 40, Y: 30},
		}, subj)
	})

	t.Run("should read point segments from space separated list", func(t *testing.T) {
		// given
		d := "M 10 20 20 20 30 30 40 30"
		x := xml.Path{
			D: d,
		}
		// when
		subj := parse.Path(x)

		// then
		assert.Equal(t, []svg.PointI{
			&svg.Point{X: 10, Y: 20, MoveTo: true},
			&svg.Point{X: 20, Y: 20},
			&svg.Point{X: 30, Y: 30},
			&svg.Point{X: 40, Y: 30},
		}, subj)

	})
}
