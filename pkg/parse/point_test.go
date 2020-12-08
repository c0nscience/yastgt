package parse_test

import (
	"testing"

	"github.com/c0nscience/yastgt/pkg/parse"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/stretchr/testify/assert"
)

func Test_Point(t *testing.T) {
	// given
	data := "4,9"

	// when
	subj := parse.Point(data)

	// then
	assert.Equal(t, &svg.Point{X: 4, Y: 9}, subj)
}
