package parse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func Test_Circle(t *testing.T) {
	t.Run("should return a svg circle", func(t *testing.T) {
		c := xml.Circle{
			CX: 34.2,
			CY: 21.15,
			R:  5.6,
		}

		subj := parse.Circle(c).(*svg.CirclePoint)

		assert.InDelta(t, 39.8, subj.CP.X, delta)
		assert.InDelta(t, 21.15, subj.CP.Y, delta)
		assert.True(t, subj.CP.MoveTo)

		assert.InDelta(t, 34.2, subj.P.X, delta)
		assert.InDelta(t, 21.15, subj.P.Y, delta)
		assert.InDelta(t, 5.6, subj.R, delta)
	})
}
