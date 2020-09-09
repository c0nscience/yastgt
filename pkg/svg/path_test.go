package svg_test

import (
	"testing"

	"github.com/c0nscience/yastgt/pkg/parser"
	"github.com/c0nscience/yastgt/pkg/svg"
	"github.com/stretchr/testify/assert"
)

func Test_PopulatePath(t *testing.T) {
	// given
	subj := svg.Path{}
	// when
	parser.M("M 1.5420259,10.163793 10.4375,31.906968", &subj.M)

	// then
	assert.Len(t, subj.M, 2)
	assert.Equal(t, subj.M[0], svg.Point{X: 1.5420259, Y: 10.163793})
	assert.Equal(t, subj.M[1], svg.Point{X: 10.4375, Y: 31.906968})
}
