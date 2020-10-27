package png_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c0nscience/yastgt/pkg/png"
)

func Test_Export(t *testing.T) {
	exp, _ := ioutil.ReadFile("../../resource/fill-test.png")
	t.Run("should export svg as png", func(t *testing.T) {
		svgFile := "../../resource/fill-test.svg"

		png.SetDpi(96)
		f, err := png.Export(svgFile)
		assert.NoError(t, err)
		defer os.Remove(f.Name())

		subj, err := ioutil.ReadFile(f.Name())
		assert.NoError(t, err)
		assert.Equal(t, exp, subj)
	})
}
