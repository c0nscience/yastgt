package parse_test

import (
	"testing"

	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
	"github.com/stretchr/testify/assert"
)

func Test_Path(t *testing.T) {
	d := "M 1.5420259,10.163793 L 14,20 25,12 M 10.4375,31.906968 H 6 V 18 Z C 10,12 30,60 22,50 m -2,0 l -5,-1 h 5 v 7 c 0,0 5,-2 4,2"
	xml := xml.Path{
		D: d,
	}

	t.Run("parse from xml", func(t *testing.T) {
		subj := parse.Path(xml)

		assert.Equal(t, []svg.PointI{
			svg.Point{X: 1.5420259, Y: 10.163793, MoveTo: true}, // move to
			svg.Point{X: 14, Y: 20},                             // lineto
			svg.Point{X: 25, Y: 12},                             // lineto
			svg.Point{X: 10.4375, Y: 31.906968, MoveTo: true},   // moveto
			svg.Point{X: 6, Y: 31.906968},                       // horizontal lineto
			svg.Point{X: 6, Y: 18},                              // vertical lineto
			svg.Point{X: 1.5420259, Y: 10.163793},               // close path
			svg.CubicPoint{ // bezier curve
				CP: svg.Point{X: 22, Y: 50},
				P1: svg.Point{X: 10, Y: 12},
				P2: svg.Point{X: 30, Y: 60},
			},
			svg.Point{X: -2, Y: 0, Rel: true, MoveTo: true}, // relative moveto

			svg.Point{X: -5, Y: -1, Rel: true}, // relative lineto
			svg.Point{X: 5, Y: -1, Rel: true},  // relative horizontal lineto
			svg.Point{X: 5, Y: 7, Rel: true},   // relative vertical lineto
			svg.CubicPoint{ // relative bezier curve
				CP:  svg.Point{X: 4, Y: 2},
				P1:  svg.Point{X: 0, Y: 0},
				P2:  svg.Point{X: 5, Y: -2},
				Rel: true,
			},
		}, subj.Points)

	})

}

func Test_ExponentialValue(t *testing.T) {
	d := "c 7e-4,4 6,7 3,5 4,5 7,5 1,5 l 1,2 5,2"

	xml := xml.Path{
		D: d,
	}

	t.Run("parse curve", func(t *testing.T) {
		subj := parse.Path(xml)

		assert.Equal(t, []svg.PointI{
			svg.CubicPoint{
				CP:  svg.Point{X: 3, Y: 5, Rel: false},
				P1:  svg.Point{X: 0.0007, Y: 4, Rel: false},
				P2:  svg.Point{X: 6, Y: 7, Rel: false},
				Rel: true,
			},
			svg.CubicPoint{
				CP:  svg.Point{X: 1, Y: 5, Rel: false},
				P1:  svg.Point{X: 4, Y: 5, Rel: false},
				P2:  svg.Point{X: 7, Y: 5, Rel: false},
				Rel: true,
			},
			svg.Point{X: 1, Y: 2, Rel: true},
			svg.Point{X: 5, Y: 2, Rel: true},
		}, subj.Points)
	})
}

func Test_CloseRelativeFirstMoveTo(t *testing.T) {
	// given
	d := "m 10,20 L 10,10 5,10 Z"
	xml := xml.Path{
		D: d,
	}
	// when
	subj := parse.Path(xml)

	// then
	assert.Equal(t, []svg.PointI{
		svg.Point{X: 10, Y: 20, MoveTo: true, Rel: true},
		svg.Point{X: 10, Y: 10},
		svg.Point{X: 5, Y: 10},
		svg.Point{X: 10, Y: 20},
	}, subj.Points)
}

func Test_CloseRelativeFirstMoveToWithRelativCloseCmd(t *testing.T) {
	// given
	d := "m 10,20 L 10,10 5,10 z"
	xml := xml.Path{
		D: d,
	}
	// when
	subj := parse.Path(xml)

	// then
	assert.Equal(t, []svg.PointI{
		svg.Point{X: 10, Y: 20, MoveTo: true, Rel: true},
		svg.Point{X: 10, Y: 10},
		svg.Point{X: 5, Y: 10},
		svg.Point{X: 10, Y: 20},
	}, subj.Points)
}

func Test_ConsecutiveRelativeMoveToCmds(t *testing.T) {
	// given
	d := "m 10,20 20,20 30,30 40,30"
	xml := xml.Path{
		D: d,
	}
	// when
	subj := parse.Path(xml)

	// then
	assert.Equal(t, []svg.PointI{
		svg.Point{X: 10, Y: 20, MoveTo: true, Rel: true},
		svg.Point{X: 20, Y: 20, Rel: true},
		svg.Point{X: 30, Y: 30, Rel: true},
		svg.Point{X: 40, Y: 30, Rel: true},
	}, subj.Points)
}

func Test_ConsecutiveMoveToCmds(t *testing.T) {
	// given
	d := "M 10,20 20,20 30,30 40,30"
	xml := xml.Path{
		D: d,
	}
	// when
	subj := parse.Path(xml)

	// then
	assert.Equal(t, []svg.PointI{
		svg.Point{X: 10, Y: 20, MoveTo: true},
		svg.Point{X: 20, Y: 20},
		svg.Point{X: 30, Y: 30},
		svg.Point{X: 40, Y: 30},
	}, subj.Points)
}
