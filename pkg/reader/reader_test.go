package reader_test

import (
	"testing"

	"github.com/c0nscience/yastgt/pkg/reader"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
	"github.com/stretchr/testify/assert"
)

func Test_Unmarshal(t *testing.T) {
	// given
	data := `
	<svg
		xmlns:dc="http://purl.org/dc/elements/1.1/"
		xmlns:cc="http://creativecommons.org/ns#"
		xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
		xmlns:svg="http://www.w3.org/2000/svg"
		xmlns="http://www.w3.org/2000/svg"
		xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
		xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
		id="SVGRoot"
		version="1.1"
		viewBox="0 0 35 35"
		height="35cm"
		width="35cm"
		inkscape:version="1.0 (4035a4f, 2020-05-01)"
		sodipodi:docname="line_circle_rect_curve.svg">
		<g
			inkscape:label="Layer 1"
			inkscape:groupmode="layer"
			id="layer1">
			<path
				sodipodi:nodetypes="cc"
				style="fill:none;stroke:#000000;stroke-width:0.0281339px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
				d="M 1.5420259,10.163793 10.4375,31.906968"
				id="line"
				inkscape:label="Line" />
		</g>
		<path
			sodipodi:nodetypes="cc"
			style="fill:none;stroke:#000000;stroke-width:0.0281339px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
			d="M 4,5 9,3"
			id="line"
			inkscape:label="Line" />
	</svg>`

	// when
	subj, err := reader.Unmarshal([]byte(data))

	t.Run("path from root", func(t *testing.T) {
		p := subj.Path[0]

		// then
		assert.NoError(t, err)
		assert.Equal(t, xml.Path{D: "M 4,5 9,3"}, p)
	})

	t.Run("path from group", func(t *testing.T) {
		p := subj.G[0].Path[0]

		// then
		assert.NoError(t, err)
		assert.Equal(t, xml.Path{D: "M 1.5420259,10.163793 10.4375,31.906968"}, p)
	})

}
