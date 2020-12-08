package transform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/parse/transform"
)

func Test_Apply(t *testing.T) {
	t.Run("should apply the given transform to all points of a path", func(t *testing.T) {
		subj := []svg.PointI{
			&svg.Point{
				X: 2,
				Y: 3,
			},
			&svg.Point{
				X: 7,
				Y: 4,
			},
			&svg.CubicPoint{
				CP: &svg.Point{X: 2, Y: 11},
				P1: &svg.Point{X: 14, Y: 18},
				P2: &svg.Point{X: 21, Y: 26},
			},
		}
		transform.Apply([]*mat.Dense{
			transform.NewTranslation(5, 5),
		}, subj)

		assert.Equal(t, &svg.Point{X: 7, Y: 8}, subj[0])
		assert.Equal(t, &svg.Point{X: 12, Y: 9}, subj[1])
		assert.Equal(t, &svg.CubicPoint{
			CP: &svg.Point{X: 7, Y: 16},
			P1: &svg.Point{X: 19, Y: 23},
			P2: &svg.Point{X: 26, Y: 31},
		}, subj[2])
	})
	t.Run("should apply multiple transforms", func(t *testing.T) {
		subj := []svg.PointI{
			&svg.Point{
				X: 2,
				Y: 3,
			},
			&svg.Point{
				X: 7,
				Y: 4,
			},
			&svg.CubicPoint{
				CP: &svg.Point{X: 2, Y: 11},
				P1: &svg.Point{X: 14, Y: 18},
				P2: &svg.Point{X: 21, Y: 26},
			},
		}

		transform.Apply([]*mat.Dense{
			transform.NewTranslation(5, 5),
			transform.NewRotation(30),
		}, subj)

		assert.Equal(t, &svg.Point{X: 5.232050807568878, Y: 8.598076211353316}, subj[0])
		assert.Equal(t, &svg.Point{X: 9.062177826491071, Y: 11.964101615137753}, subj[1])
		assert.Equal(t, &svg.CubicPoint{
			CP: &svg.Point{X: 1.2320508075688785, Y: 15.526279441628827},
			P1: &svg.Point{X: 8.124355652982144, Y: 27.588457268119896},
			P2: &svg.Point{X: 10.186533479473214, Y: 38.0166604983954},
		}, subj[2])
	})

	t.Run("should return same path if not transformation is provided", func(t *testing.T) {
		subj := []svg.PointI{
			&svg.Point{
				X: 2,
				Y: 3,
			},
			&svg.Point{
				X: 7,
				Y: 4,
			},
			&svg.CubicPoint{
				CP: &svg.Point{X: 2, Y: 11},
				P1: &svg.Point{X: 14, Y: 18},
				P2: &svg.Point{X: 21, Y: 26},
			},
		}
		transform.Apply([]*mat.Dense{}, subj)

		assert.Equal(t, &svg.Point{X: 2, Y: 3}, subj[0])
		assert.Equal(t, &svg.Point{X: 7, Y: 4}, subj[1])
		assert.Equal(t, &svg.CubicPoint{
			CP: &svg.Point{X: 2, Y: 11},
			P1: &svg.Point{X: 14, Y: 18},
			P2: &svg.Point{X: 21, Y: 26},
		}, subj[2])

	})
}
