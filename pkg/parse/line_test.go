package parse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

const delta = 1e-6

func Test_Line(t *testing.T) {
	t.Run("should return a path with the line coordinates", func(t *testing.T) {
		l := xml.Line{
			X1: 13.5,
			Y1: 10.4,
			X2: 16.0,
			Y2: 19.0,
		}

		subj := parse.Line(l)

		assert.InDelta(t, 13.5, subj[0].CurrPt().X, delta)
		assert.InDelta(t, 10.4, subj[0].CurrPt().Y, delta)
		assert.True(t, subj[0].CurrPt().MoveTo)
		assert.InDelta(t, 16.0, subj[1].CurrPt().X, delta)
		assert.InDelta(t, 19.0, subj[1].CurrPt().Y, delta)
	})
}
