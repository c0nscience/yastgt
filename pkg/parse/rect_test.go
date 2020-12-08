package parse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func Test_Rect(t *testing.T) {
	t.Run("should return a path describing the rectangle", func(t *testing.T) {
		r := xml.Rect{
			X:      23.4,
			Y:      56.5,
			Height: 20,
			Width:  40,
		}

		subj := parse.Rect(r)

		assert.Equal(t, &svg.Point{X: 23.4, Y: 56.5, MoveTo: true}, subj[0])
		assert.Equal(t, &svg.Point{X: 63.4, Y: 56.5}, subj[1])
		assert.Equal(t, &svg.Point{X: 63.4, Y: 76.5}, subj[2])
		assert.Equal(t, &svg.Point{X: 23.4, Y: 76.5}, subj[3])
		assert.Equal(t, &svg.Point{X: 23.4, Y: 56.5}, subj[4])
	})
}
