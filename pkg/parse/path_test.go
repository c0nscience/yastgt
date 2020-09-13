package parse_test

import (
	"testing"

	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
	"github.com/stretchr/testify/assert"
)

func Test_Path(t *testing.T) {
	d := "M 1.5420259,10.163793 10.4375,31.906968 L 14,20 25,12 H 6 V 18 Z C 10,12 30,60 22,50 m -2,0 -10,2 l -5,-1 h 5 v 7 c 0,0 5,-2 4,2"
	xml := xml.Path{
		D: d,
	}

	t.Run("parse from xml", func(t *testing.T) {
		subj := parse.Path(xml)

		assert.Len(t, subj.Points, 14)

		t.Run("moveto points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: 1.5420259, Y: 10.163793}, subj.Points[0])
			assert.Equal(t, svg.Point{X: 10.4375, Y: 31.906968}, subj.Points[1])
		})

		t.Run("lineto points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: 14, Y: 20}, subj.Points[2])
			assert.Equal(t, svg.Point{X: 25, Y: 12}, subj.Points[3])
		})

		t.Run("horizontal lineto points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: 6, Y: 12}, subj.Points[4])
		})

		t.Run("vertical lineto points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: 6, Y: 18}, subj.Points[5])
		})

		t.Run("closepath points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: 1.5420259, Y: 10.163793}, subj.Points[6])
		})

		t.Run("closepath points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: 1.5420259, Y: 10.163793}, subj.Points[6])
		})

		t.Run("curveto points", func(t *testing.T) {
			assert.Equal(t, svg.CubicPoint{
				CP: svg.Point{X: 22, Y: 50},
				P1: svg.Point{X: 10, Y: 12},
				P2: svg.Point{X: 30, Y: 60},
			}, subj.Points[7])
		})

		t.Run("relative moveto points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: -2, Y: 0, Rel: true}, subj.Points[8])
			assert.Equal(t, svg.Point{X: -10, Y: 2, Rel: true}, subj.Points[9])
		})

		t.Run("relative lineto points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: -5, Y: -1, Rel: true}, subj.Points[10])
		})

		t.Run("relative horizontal lineto points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: 5, Y: -1, Rel: true}, subj.Points[11])
		})

		t.Run("relative vertical lineto points", func(t *testing.T) {
			assert.Equal(t, svg.Point{X: 5, Y: 7, Rel: true}, subj.Points[12])
		})

		t.Run("relative curveto points", func(t *testing.T) {
			assert.Equal(t, svg.CubicPoint{
				CP:  svg.Point{X: 4, Y: 2},
				P1:  svg.Point{X: 0, Y: 0},
				P2:  svg.Point{X: 5, Y: -2},
				Rel: true,
			}, subj.Points[13])
		})
	})

}
