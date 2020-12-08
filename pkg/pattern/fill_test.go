package pattern_test

import (
	"image/color"
	gpng "image/png"
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
		img, _ := gpng.Decode(f)

		// when
		subj := pattern.Diagonal(img)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		t.Run("starting from max Y position", func(t *testing.T) {
			assert.InDelta(t, 41.7, x(0), delta)
			assert.InDelta(t, 98.29, y(0), delta)
			assert.True(t, moveTo(0))
			assert.InDelta(t, 45, x(1), delta)
			assert.InDelta(t, 101, y(1), delta)

			assert.InDelta(t, 66, x(2), delta)
			assert.InDelta(t, 108, y(2), delta)
			assert.True(t, moveTo(2))
			assert.InDelta(t, 48.8, x(3), delta)
			assert.InDelta(t, 91, y(3), delta)

			assert.InDelta(t, 55.8, x(4), delta)
			assert.InDelta(t, 84, y(4), delta)
			assert.True(t, moveTo(4))
			assert.InDelta(t, 87, x(5), delta)
			assert.InDelta(t, 115, y(5), delta)

			assert.InDelta(t, 108, x(6), delta)
			assert.InDelta(t, 122, y(6), delta)
			assert.True(t, moveTo(6))
			assert.InDelta(t, 62.9, x(7), delta)
			assert.InDelta(t, 77, y(7), delta)

			assert.InDelta(t, 70, x(8), delta)
			assert.InDelta(t, 70, y(8), delta)
			assert.True(t, moveTo(8))
			assert.InDelta(t, 129.8, x(9), delta)
			assert.InDelta(t, 129.8, y(9), delta)
		})

		t.Run("clockwise over into x direction", func(t *testing.T) {
			assert.InDelta(t, 122.8, x(10), delta)
			assert.InDelta(t, 108.81, y(10), delta)
			assert.True(t, moveTo(10))
			assert.InDelta(t, 77.1, x(11), delta)
			assert.InDelta(t, 62.93, y(11), delta)

			assert.InDelta(t, 84.14, x(12), delta)
			assert.InDelta(t, 56.09, y(12), delta)
			assert.True(t, moveTo(12))
			assert.InDelta(t, 115.89, x(13), delta)
			assert.InDelta(t, 87.71, y(13), delta)

			assert.InDelta(t, 108.74, x(14), delta)
			assert.InDelta(t, 66.54, y(14), delta)
			assert.True(t, moveTo(14))
			assert.InDelta(t, 91.55, x(15), delta)
			assert.InDelta(t, 48.82, y(15), delta)

			assert.InDelta(t, 98.29, x(16), delta)
			assert.InDelta(t, 41.80, y(16), delta)
			assert.True(t, moveTo(16))
			assert.InDelta(t, 101.34, x(17), delta)
			assert.InDelta(t, 44.98, y(17), delta)

		})

	})

	t.Run("should draw fill pattern over multiple shapes", func(t *testing.T) {
		f, _ := os.Open("../../resource/diag-fill-multiple-shapes.png")
		pattern.SetGap(10)
		pattern.SetDpi(96)
		pattern.SetDegrees(45)
		pattern.SetColor(color.NRGBA{R: 255, A: 255})
		pattern.SetThreshold(0)
		img, _ := gpng.Decode(f)

		// when
		subj := pattern.Diagonal(img)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		assert.InDelta(t, 41.7, x(0), delta)
		assert.InDelta(t, 98.29, y(0), delta)
		assert.True(t, moveTo(0))
		assert.InDelta(t, 45, x(1), delta)
		assert.InDelta(t, 101, y(1), delta)

		assert.InDelta(t, 53, x(2), delta)
		assert.InDelta(t, 108.8, y(2), delta)
		assert.True(t, moveTo(2))
		assert.InDelta(t, 126.72, x(3), delta)
		assert.InDelta(t, 183, y(3), delta)

		assert.InDelta(t, 55.8, x(8), delta)
		assert.InDelta(t, 84, y(8), delta)
		assert.True(t, moveTo(8))
		assert.InDelta(t, 87, x(9), delta)
		assert.InDelta(t, 115, y(9), delta)

		assert.InDelta(t, 93.73, x(10), delta)
		assert.InDelta(t, 121.98, y(10), delta)
		assert.True(t, moveTo(10))
		assert.InDelta(t, 140.78, x(11), delta)
		assert.InDelta(t, 169.08, y(11), delta)

	})

	t.Run("should draw 30 degree pattern with 10 mm distance", func(t *testing.T) {
		// given
		f, _ := os.Open("../../resource/diag-fill.png")
		pattern.SetGap(10)
		pattern.SetDpi(96)
		pattern.SetDegrees(30)
		pattern.SetColor(color.NRGBA{R: 255, A: 255})
		pattern.SetThreshold(0)
		img, _ := gpng.Decode(f)

		// when
		subj := pattern.Diagonal(img)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		assert.InDelta(t, 44.9, x(0), delta)
		assert.InDelta(t, 95.42, y(0), delta)
		assert.True(t, moveTo(0))
		assert.InDelta(t, 68, x(1), delta)
		assert.InDelta(t, 109, y(1), delta)

		assert.InDelta(t, 74.18, x(8), delta)
		assert.InDelta(t, 66.04, y(8), delta)
		assert.True(t, moveTo(8))
		assert.InDelta(t, 117, x(9), delta)
		assert.InDelta(t, 90.83, y(9), delta)
	})

	t.Run("should draw 90 degree pattern with 10 mm distance", func(t *testing.T) {
		f, _ := os.Open("../../resource/diag-fill.png")
		pattern.SetGap(10)
		pattern.SetDpi(96)
		pattern.SetDegrees(90)
		pattern.SetColor(color.NRGBA{R: 255, A: 255})
		pattern.SetThreshold(0)
		img, _ := gpng.Decode(f)

		// when
		subj := pattern.Diagonal(img)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		assert.InDelta(t, 50.27, x(0), delta)
		assert.InDelta(t, 89.92, y(0), delta)
		assert.True(t, moveTo(0))
		assert.InDelta(t, 50.27, x(1), delta)
		assert.InDelta(t, 103.35, y(1), delta)

	})

	t.Run("should draw 0 degree pattern with 10 mm distance", func(t *testing.T) {
		f, _ := os.Open("../../resource/diag-fill-multiple-shapes-bigger.png")
		pattern.SetGap(10)
		pattern.SetDpi(96)
		pattern.SetDegrees(0)
		pattern.SetColor(color.NRGBA{R: 255, A: 255})
		pattern.SetThreshold(0)
		img, _ := gpng.Decode(f)

		// when
		subj := pattern.Diagonal(img)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		assert.InDelta(t, 284.8, x(18), delta)
		assert.InDelta(t, 251.2, y(18), delta)
		assert.True(t, moveTo(18))
		assert.InDelta(t, 105.5, x(19), delta)
		assert.InDelta(t, 251.2, y(19), delta)

	})

	t.Run("should draw 120 degree pattern with 10 mm distance", func(t *testing.T) {
		f, _ := os.Open("../../resource/diag-fill-multiple-shapes-bigger.png")
		pattern.SetGap(20)
		pattern.SetDpi(96)
		pattern.SetDegrees(120)
		pattern.SetColor(color.NRGBA{R: 255, A: 255})
		pattern.SetThreshold(0)
		img, _ := gpng.Decode(f)

		// when
		subj := pattern.Diagonal(img)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		assert.Len(t, subj, 66)

		assert.InDelta(t, 335, x(0), delta)
		assert.InDelta(t, 185, y(0), delta)
		assert.True(t, moveTo(0))
		assert.InDelta(t, 313.5, x(1), delta)
		assert.InDelta(t, 223, y(1), delta)

		assert.InDelta(t, 153, x(12), delta)
		assert.InDelta(t, 301, y(12), delta)
		assert.True(t, moveTo(12))
		assert.InDelta(t, 201, x(13), delta)
		assert.InDelta(t, 218, y(13), delta)

	})

	t.Run("should draw 150 degree pattern with 10 mm distance", func(t *testing.T) {
		f, _ := os.Open("../../resource/diag-fill-multiple-shapes-bigger.png")
		pattern.SetGap(20)
		pattern.SetDpi(96)
		pattern.SetDegrees(150)
		pattern.SetColor(color.NRGBA{R: 255, A: 255})
		pattern.SetThreshold(0)
		img, _ := gpng.Decode(f)

		// when
		subj := pattern.Diagonal(img)
		x := xExtractor(subj)
		y := yExtractor(subj)
		moveTo := moveToExtractor(subj)

		assert.Len(t, subj, 66)

		assert.InDelta(t, 137, x(6), delta)
		assert.InDelta(t, 284, y(6), delta)
		assert.True(t, moveTo(6))
		assert.InDelta(t, 325.5, x(7), delta)
		assert.InDelta(t, 175.7, y(7), delta)

		assert.InDelta(t, 105, x(34), delta)
		assert.InDelta(t, 187, y(34), delta)
		assert.True(t, moveTo(34))
		assert.InDelta(t, 67, x(35), delta)
		assert.InDelta(t, 209.5, y(35), delta)

	})
}

func xExtractor(p []svg.PointI) func(int) float64 {
	return func(i int) float64 {
		return p[i].CurrPt().X
	}
}

func yExtractor(p []svg.PointI) func(int) float64 {
	return func(i int) float64 {
		return p[i].CurrPt().Y
	}
}

func moveToExtractor(p []svg.PointI) func(int) bool {
	return func(i int) bool {
		return p[i].CurrPt().MoveTo
	}
}
