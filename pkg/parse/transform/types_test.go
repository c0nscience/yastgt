package transform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c0nscience/yastgt/pkg/parse/transform"
)

const delta = 1e-6

func Test_Types(t *testing.T) {
	t.Run("should create a transform matrix", func(t *testing.T) {
		a := 5.0
		b := 6.0
		c := 7.0
		d := 8.0
		e := 9.0
		f := 10.0

		subj := transform.New(a, b, c, d, e, f)

		assert.Equal(t, 5.0, subj.At(0, 0))
		assert.Equal(t, 7.0, subj.At(0, 1))
		assert.Equal(t, 9.0, subj.At(0, 2))

		assert.Equal(t, 6.0, subj.At(1, 0))
		assert.Equal(t, 8.0, subj.At(1, 1))
		assert.Equal(t, 10.0, subj.At(1, 2))

		assert.Equal(t, 0.0, subj.At(2, 0))
		assert.Equal(t, 0.0, subj.At(2, 1))
		assert.Equal(t, 1.0, subj.At(2, 2))
	})
	t.Run("should create a translate matrix", func(t *testing.T) {
		tx := 10.0
		ty := 11.0

		subj := transform.NewTranslation(tx, ty)

		assert.Equal(t, 1.0, subj.At(0, 0))
		assert.Equal(t, 0.0, subj.At(1, 0))
		assert.Equal(t, 0.0, subj.At(2, 0))

		assert.Equal(t, 0.0, subj.At(0, 1))
		assert.Equal(t, 1.0, subj.At(1, 1))
		assert.Equal(t, 0.0, subj.At(2, 1))

		assert.Equal(t, 10.0, subj.At(0, 2))
		assert.Equal(t, 11.0, subj.At(1, 2))
		assert.Equal(t, 1.0, subj.At(2, 2))
	})
	t.Run("should create a scale matrix", func(t *testing.T) {
		sx := 13.0
		sy := 14.0

		subj := transform.NewScaling(sx, sy)

		assert.Equal(t, 13.0, subj.At(0, 0))
		assert.Equal(t, 0.0, subj.At(1, 0))
		assert.Equal(t, 0.0, subj.At(2, 0))

		assert.Equal(t, 0.0, subj.At(0, 1))
		assert.Equal(t, 14.0, subj.At(1, 1))
		assert.Equal(t, 0.0, subj.At(2, 1))

		assert.Equal(t, 0.0, subj.At(0, 2))
		assert.Equal(t, 0.0, subj.At(1, 2))
		assert.Equal(t, 1.0, subj.At(2, 2))
	})
	t.Run("should create a rotate matrix", func(t *testing.T) {
		a := 35.0

		subj := transform.NewRotation(a)

		assert.InDelta(t, 0.819152, subj.At(0, 0), delta)
		assert.InDelta(t, 0.573576, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, -0.573576, subj.At(0, 1), delta)
		assert.InDelta(t, 0.819152, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 0.0, subj.At(0, 2), delta)
		assert.InDelta(t, 0.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
	t.Run("should create a skewX matrix", func(t *testing.T) {
		a := 55.0

		subj := transform.NewSkewX(a)

		assert.InDelta(t, 1.0, subj.At(0, 0), delta)
		assert.InDelta(t, 0.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 1.4281480, subj.At(0, 1), delta)
		assert.InDelta(t, 1.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 0.0, subj.At(0, 2), delta)
		assert.InDelta(t, 0.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
	t.Run("should create a skewY matrix", func(t *testing.T) {
		a := 55.0

		subj := transform.NewSkewY(a)

		assert.InDelta(t, 1.0, subj.At(0, 0), delta)
		assert.InDelta(t, 1.4281480, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 0.0, subj.At(0, 1), delta)
		assert.InDelta(t, 1.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 0.0, subj.At(0, 2), delta)
		assert.InDelta(t, 0.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
}
