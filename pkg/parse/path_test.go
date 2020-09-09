package parse_test

import (
	"testing"

	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/svg"
	"github.com/stretchr/testify/assert"
)

func Test_Path(t *testing.T) {
	// given
	data := "M 1.5420259,10.163793 10.4375,31.906968 L 14,20 25,12 H 6 V 12.7638292"

	t.Run("find M", func(t *testing.T) {
		s := parse.FindM(data)

		assert.Equal(t, "M 1.5420259,10.163793 10.4375,31.906968", s)
		t.Run("parse M from string", func(t *testing.T) {
			// when
			subj := parse.M(s)

			// then
			assert.Len(t, subj, 2)
			assert.Equal(t, subj[0], svg.Point{X: 1.5420259, Y: 10.163793})
			assert.Equal(t, subj[1], svg.Point{X: 10.4375, Y: 31.906968})
		})
	})

	t.Run("find L", func(t *testing.T) {
		s := parse.FindL(data)

		assert.Equal(t, "L 14,20 25,12", s)

		t.Run("parse L from string", func(t *testing.T) {
			// when
			subj := parse.L(s)

			// then
			assert.Len(t, subj, 2)
			assert.Equal(t, subj[0], svg.Point{X: 14, Y: 20})
			assert.Equal(t, subj[1], svg.Point{X: 25, Y: 12})
		})
	})

	t.Run("find H", func(t *testing.T) {
		s := parse.FindH(data)

		assert.Equal(t, "H 6", s)

		t.Run("parse H from string", func(t *testing.T) {
			// when
			subj := parse.H(s)

			// then
			assert.Equal(t, float64(6), subj)
		})
	})

	t.Run("find V", func(t *testing.T) {
		s := parse.FindV(data)

		assert.Equal(t, "V 12.7638292", s)

		t.Run("parse V from string", func(t *testing.T) {
			// when
			subj := parse.V(s)

			// then
			assert.Equal(t, 12.7638292, subj)
		})
	})

}
