package transform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c0nscience/yastgt/pkg/parse/transform"
)

func Test_Parse(t *testing.T) {
	t.Run("should return a matrix", func(t *testing.T) {
		s := "matrix(2,3,4,5,6,7)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 2.0, subj.At(0, 0), delta)
		assert.InDelta(t, 3.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 4.0, subj.At(0, 1), delta)
		assert.InDelta(t, 5.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 6.0, subj.At(0, 2), delta)
		assert.InDelta(t, 7.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
	t.Run("should return a matrix with space separated list", func(t *testing.T) {
		s := "matrix(2 3 4 5 6 7)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 2.0, subj.At(0, 0), delta)
		assert.InDelta(t, 3.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 4.0, subj.At(0, 1), delta)
		assert.InDelta(t, 5.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 6.0, subj.At(0, 2), delta)
		assert.InDelta(t, 7.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
	t.Run("should return a translation matrix with tx and ty", func(t *testing.T) {
		s := "translate(2,3)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 1.0, subj.At(0, 0), delta)
		assert.InDelta(t, 0.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 0.0, subj.At(0, 1), delta)
		assert.InDelta(t, 1.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 2.0, subj.At(0, 2), delta)
		assert.InDelta(t, 3.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
	t.Run("should return a translation matrix with tx and ty if space included", func(t *testing.T) {
		s := "translate(2, 3)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 1.0, subj.At(0, 0), delta)
		assert.InDelta(t, 0.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 0.0, subj.At(0, 1), delta)
		assert.InDelta(t, 1.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 2.0, subj.At(0, 2), delta)
		assert.InDelta(t, 3.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
	t.Run("should return a translation matrix with tx and ty if space and decimal digits included", func(t *testing.T) {
		s := "translate(2.64, 3.78)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 1.0, subj.At(0, 0), delta)
		assert.InDelta(t, 0.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 0.0, subj.At(0, 1), delta)
		assert.InDelta(t, 1.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 2.64, subj.At(0, 2), delta)
		assert.InDelta(t, 3.78, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
	t.Run("should return a translation matrix with ty as 0 if only one number found", func(t *testing.T) {
		s := "translate(2)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 1.0, subj.At(0, 0), delta)
		assert.InDelta(t, 0.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 0.0, subj.At(0, 1), delta)
		assert.InDelta(t, 1.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 2.0, subj.At(0, 2), delta)
		assert.InDelta(t, 0.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)
	})
	t.Run("should return a scaling matrix with sx and sy", func(t *testing.T) {
		s := "scale(4,5)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 4.0, subj.At(0, 0), delta)
		assert.InDelta(t, 0.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 0.0, subj.At(0, 1), delta)
		assert.InDelta(t, 5.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 0.0, subj.At(0, 2), delta)
		assert.InDelta(t, 0.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
	t.Run("should return a scaling matrix with sy is equal to sx", func(t *testing.T) {
		s := "scale(4)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 4.0, subj.At(0, 0), delta)
		assert.InDelta(t, 0.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 0.0, subj.At(0, 1), delta)
		assert.InDelta(t, 4.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 0.0, subj.At(0, 2), delta)
		assert.InDelta(t, 0.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)

	})
	t.Run("should return a rotate matrix with an angle", func(t *testing.T) {
		s := "rotate(40)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 0.7660444, subj.At(0, 0), delta)
		assert.InDelta(t, 0.6427876, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, -0.6427876, subj.At(0, 1), delta)
		assert.InDelta(t, 0.7660444, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 0.0, subj.At(0, 2), delta)
		assert.InDelta(t, 0.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)
	})
	t.Run("should return a rotate matrix with an angle and a rotation point", func(t *testing.T) {
		s := "rotate(40, 6, 7)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 0.7660444, subj.At(0, 0), delta)
		assert.InDelta(t, 0.6427876, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, -0.6427876, subj.At(0, 1), delta)
		assert.InDelta(t, 0.7660444, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 5.9032466, subj.At(0, 2), delta)
		assert.InDelta(t, -2.2190367, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)
	})
	t.Run("should return a skewX matrix with an angle", func(t *testing.T) {
		s := "skewX(40)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 1.0, subj.At(0, 0), delta)
		assert.InDelta(t, 0.0, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 0.839099, subj.At(0, 1), delta)
		assert.InDelta(t, 1.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 0.0, subj.At(0, 2), delta)
		assert.InDelta(t, 0.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)
	})
	t.Run("should return a skewY matrix with an angle", func(t *testing.T) {
		s := "skewY(40)"

		subj := transform.ParseTypes(s)[0]

		assert.InDelta(t, 1.0, subj.At(0, 0), delta)
		assert.InDelta(t, 0.839099, subj.At(1, 0), delta)
		assert.InDelta(t, 0.0, subj.At(2, 0), delta)

		assert.InDelta(t, 0.0, subj.At(0, 1), delta)
		assert.InDelta(t, 1.0, subj.At(1, 1), delta)
		assert.InDelta(t, 0.0, subj.At(2, 1), delta)

		assert.InDelta(t, 0.0, subj.At(0, 2), delta)
		assert.InDelta(t, 0.0, subj.At(1, 2), delta)
		assert.InDelta(t, 1.0, subj.At(2, 2), delta)
	})
	t.Run("should return no transforms for an unknown type", func(t *testing.T) {
		s := "unknown(5,4)"

		subj := transform.ParseTypes(s)
		assert.Empty(t, subj)
	})
	t.Run("should return multiple transforms from one string", func(t *testing.T) {
		s := "translate(2,3)skewY(40)"

		subj := transform.ParseTypes(s)

		assert.InDelta(t, 1.0, subj[0].At(0, 0), delta)
		assert.InDelta(t, 0.0, subj[0].At(1, 0), delta)
		assert.InDelta(t, 0.0, subj[0].At(2, 0), delta)

		assert.InDelta(t, 0.0, subj[0].At(0, 1), delta)
		assert.InDelta(t, 1.0, subj[0].At(1, 1), delta)
		assert.InDelta(t, 0.0, subj[0].At(2, 1), delta)

		assert.InDelta(t, 2.0, subj[0].At(0, 2), delta)
		assert.InDelta(t, 3.0, subj[0].At(1, 2), delta)
		assert.InDelta(t, 1.0, subj[0].At(2, 2), delta)

		assert.InDelta(t, 1.0, subj[1].At(0, 0), delta)
		assert.InDelta(t, 0.839099, subj[1].At(1, 0), delta)
		assert.InDelta(t, 0.0, subj[1].At(2, 0), delta)

		assert.InDelta(t, 0.0, subj[1].At(0, 1), delta)
		assert.InDelta(t, 1.0, subj[1].At(1, 1), delta)
		assert.InDelta(t, 0.0, subj[1].At(2, 1), delta)

		assert.InDelta(t, 0.0, subj[1].At(0, 2), delta)
		assert.InDelta(t, 0.0, subj[1].At(1, 2), delta)
		assert.InDelta(t, 1.0, subj[1].At(2, 2), delta)

	})
	t.Run("should return multiple transforms from one string with spaces", func(t *testing.T) {
		s := "translate(2,3) skewY(40)"

		subj := transform.ParseTypes(s)

		assert.InDelta(t, 1.0, subj[0].At(0, 0), delta)
		assert.InDelta(t, 0.0, subj[0].At(1, 0), delta)
		assert.InDelta(t, 0.0, subj[0].At(2, 0), delta)

		assert.InDelta(t, 0.0, subj[0].At(0, 1), delta)
		assert.InDelta(t, 1.0, subj[0].At(1, 1), delta)
		assert.InDelta(t, 0.0, subj[0].At(2, 1), delta)

		assert.InDelta(t, 2.0, subj[0].At(0, 2), delta)
		assert.InDelta(t, 3.0, subj[0].At(1, 2), delta)
		assert.InDelta(t, 1.0, subj[0].At(2, 2), delta)

		assert.InDelta(t, 1.0, subj[1].At(0, 0), delta)
		assert.InDelta(t, 0.839099, subj[1].At(1, 0), delta)
		assert.InDelta(t, 0.0, subj[1].At(2, 0), delta)

		assert.InDelta(t, 0.0, subj[1].At(0, 1), delta)
		assert.InDelta(t, 1.0, subj[1].At(1, 1), delta)
		assert.InDelta(t, 0.0, subj[1].At(2, 1), delta)

		assert.InDelta(t, 0.0, subj[1].At(0, 2), delta)
		assert.InDelta(t, 0.0, subj[1].At(1, 2), delta)
		assert.InDelta(t, 1.0, subj[1].At(2, 2), delta)

	})
}
