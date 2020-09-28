package generate_test

import (
	"os"
	"testing"

	"github.com/c0nscience/yastgt/pkg/generate"
	"github.com/stretchr/testify/assert"
)

var delta = 0.6

func Test_FromPNG(t *testing.T) {
	// given
	f, _ := os.Open("../../resource/fill-test.png")
	// when

	//152*303px
	subj := generate.FromPNG(f)

	// then
	assert.Len(t, subj, 8)
	assert.InDelta(t, 55, subj[0].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 30, subj[0].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[0].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 95, subj[0].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 30, subj[0].Points[1].CurrPt().Y, delta)

	assert.InDelta(t, 95, subj[1].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 40, subj[1].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[1].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 55, subj[1].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 40, subj[1].Points[1].CurrPt().Y, delta)

	assert.InDelta(t, 55, subj[2].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 50, subj[2].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[2].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 95, subj[2].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 50, subj[2].Points[1].CurrPt().Y, delta)

	assert.InDelta(t, 95, subj[3].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 60, subj[3].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[3].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 55, subj[3].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 60, subj[3].Points[1].CurrPt().Y, delta)

	assert.InDelta(t, 55, subj[4].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 70, subj[4].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[4].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 95, subj[4].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 70, subj[4].Points[1].CurrPt().Y, delta)

	assert.InDelta(t, 95, subj[5].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 80, subj[5].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[5].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 55, subj[5].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 80, subj[5].Points[1].CurrPt().Y, delta)

	assert.InDelta(t, 55, subj[6].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 90, subj[6].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[6].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 95, subj[6].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 90, subj[6].Points[1].CurrPt().Y, delta)

	assert.InDelta(t, 95, subj[7].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 100, subj[7].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[7].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 55, subj[7].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 100, subj[7].Points[1].CurrPt().Y, delta)
}
func Test_FromPNGWithMultipleShapesInALine(t *testing.T) {
	// given
	f, _ := os.Open("../../resource/fill-test2.png")
	// when

	//152*303px
	subj := generate.FromPNG(f)

	// then
	assert.InDelta(t, 55, subj[0].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 30, subj[0].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[0].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 95, subj[0].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 30, subj[0].Points[1].CurrPt().Y, delta)

	assert.InDelta(t, 130, subj[1].Points[0].CurrPt().X, delta)
	assert.InDelta(t, 30, subj[1].Points[0].CurrPt().Y, delta)
	assert.True(t, subj[1].Points[0].CurrPt().MoveTo)
	assert.InDelta(t, 170, subj[1].Points[1].CurrPt().X, delta)
	assert.InDelta(t, 30, subj[1].Points[1].CurrPt().Y, delta)

}
