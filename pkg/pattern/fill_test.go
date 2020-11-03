package pattern_test

import (
	"image/color"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/pattern"
)

var delta = 0.6

func Test_Fill(t *testing.T) {
	t.Run("should draw 45 degree pattern with 10 mm distance", func(t *testing.T) {
		// given
		f, _ := os.Open("../../resource/diag-fill.png")
		pattern.SetGap(10)
		pattern.SetDpi(96)
		pattern.SetDegrees(45)
		pattern.SetColor(color.NRGBA{R: 255, A: 255})
		pattern.SetThreshold(0)

		// when
		subj := pattern.Diagonal(f)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		t.Run("starting from max Y position", func(t *testing.T) {
			assert.InDelta(t, 41.7, x(0, 0), delta)
			assert.InDelta(t, 98.29, y(0, 0), delta)
			assert.True(t, moveTo(0, 0))
			assert.InDelta(t, 45, x(0, 1), delta)
			assert.InDelta(t, 101, y(0, 1), delta)

			assert.InDelta(t, 66, x(1, 0), delta)
			assert.InDelta(t, 108, y(1, 0), delta)
			assert.True(t, moveTo(1, 0))
			assert.InDelta(t, 48.8, x(1, 1), delta)
			assert.InDelta(t, 91, y(1, 1), delta)

			assert.InDelta(t, 55.8, x(2, 0), delta)
			assert.InDelta(t, 84, y(2, 0), delta)
			assert.True(t, moveTo(2, 0))
			assert.InDelta(t, 87, x(2, 1), delta)
			assert.InDelta(t, 115, y(2, 1), delta)

			assert.InDelta(t, 108, x(3, 0), delta)
			assert.InDelta(t, 122, y(3, 0), delta)
			assert.True(t, moveTo(3, 0))
			assert.InDelta(t, 62.9, x(3, 1), delta)
			assert.InDelta(t, 77, y(3, 1), delta)

			assert.InDelta(t, 70, x(4, 0), delta)
			assert.InDelta(t, 70, y(4, 0), delta)
			assert.True(t, moveTo(4, 0))
			assert.InDelta(t, 129.8, x(4, 1), delta)
			assert.InDelta(t, 129.8, y(4, 1), delta)
		})

		t.Run("clockwise over into x direction", func(t *testing.T) {
			assert.InDelta(t, 122.8, x(5, 0), delta)
			assert.InDelta(t, 108.81, y(5, 0), delta)
			assert.True(t, moveTo(5, 0))
			assert.InDelta(t, 77.1, x(5, 1), delta)
			assert.InDelta(t, 62.93, y(5, 1), delta)

			assert.InDelta(t, 84.14, x(6, 0), delta)
			assert.InDelta(t, 56.09, y(6, 0), delta)
			assert.True(t, moveTo(6, 0))
			assert.InDelta(t, 115.89, x(6, 1), delta)
			assert.InDelta(t, 87.71, y(6, 1), delta)

			assert.InDelta(t, 108.74, x(7, 0), delta)
			assert.InDelta(t, 66.54, y(7, 0), delta)
			assert.True(t, moveTo(7, 0))
			assert.InDelta(t, 91.55, x(7, 1), delta)
			assert.InDelta(t, 48.82, y(7, 1), delta)

			assert.InDelta(t, 98.29, x(8, 0), delta)
			assert.InDelta(t, 41.80, y(8, 0), delta)
			assert.True(t, moveTo(8, 0))
			assert.InDelta(t, 101.34, x(8, 1), delta)
			assert.InDelta(t, 44.98, y(8, 1), delta)

		})

	})

	t.Run("should draw fill pattern over multiple shapes", func(t *testing.T) {
		f, _ := os.Open("../../resource/diag-fill-multiple-shapes.png")
		pattern.SetGap(10)
		pattern.SetDpi(96)
		pattern.SetDegrees(45)
		pattern.SetColor(color.NRGBA{R: 255, A: 255})
		pattern.SetThreshold(0)

		// when
		subj := pattern.Diagonal(f)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		assert.InDelta(t, 41.7, x(0, 0), delta)
		assert.InDelta(t, 98.29, y(0, 0), delta)
		assert.True(t, moveTo(0, 0))
		assert.InDelta(t, 45, x(0, 1), delta)
		assert.InDelta(t, 101, y(0, 1), delta)

		assert.InDelta(t, 53, x(1, 0), delta)
		assert.InDelta(t, 108.8, y(1, 0), delta)
		assert.True(t, moveTo(1, 0))
		assert.InDelta(t, 126.72, x(1, 1), delta)
		assert.InDelta(t, 183, y(1, 1), delta)

		assert.InDelta(t, 55.8, x(4, 0), delta)
		assert.InDelta(t, 84, y(4, 0), delta)
		assert.True(t, moveTo(4, 0))
		assert.InDelta(t, 87, x(4, 1), delta)
		assert.InDelta(t, 115, y(4, 1), delta)

		assert.InDelta(t, 93.73, x(5, 0), delta)
		assert.InDelta(t, 121.98, y(5, 0), delta)
		assert.True(t, moveTo(5, 0))
		assert.InDelta(t, 140.78, x(5, 1), delta)
		assert.InDelta(t, 169.08, y(5, 1), delta)

	})

	t.Run("should draw 30 degree pattern with 10 mm distance", func(t *testing.T) {
		// given
		f, _ := os.Open("../../resource/diag-fill.png")
		pattern.SetGap(10)
		pattern.SetDpi(96)
		pattern.SetDegrees(30)
		pattern.SetColor(color.NRGBA{R: 255, A: 255})
		pattern.SetThreshold(0)

		// when
		subj := pattern.Diagonal(f)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		assert.InDelta(t, 44.9, x(0, 0), delta)
		assert.InDelta(t, 95.42, y(0, 0), delta)
		assert.True(t, moveTo(0, 0))
		assert.InDelta(t, 68, x(0, 1), delta)
		assert.InDelta(t, 109, y(0, 1), delta)

		assert.InDelta(t, 74.18, x(4, 0), delta)
		assert.InDelta(t, 66.04, y(4, 0), delta)
		assert.True(t, moveTo(4, 0))
		assert.InDelta(t, 117, x(4, 1), delta)
		assert.InDelta(t, 90.83, y(4, 1), delta)
	})
}

func xExtractor(p []svg.Path) func(int, int) float64 {
	return func(i int, j int) float64 {
		return p[i].Points[j].CurrPt().X
	}
}

func yExtractor(p []svg.Path) func(int, int) float64 {
	return func(i int, j int) float64 {
		return p[i].Points[j].CurrPt().Y
	}
}

func moveToExtractor(p []svg.Path) func(int, int) bool {
	return func(i int, j int) bool {
		return p[i].Points[j].CurrPt().MoveTo
	}
}
