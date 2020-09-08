package svg_test

import (
	"encoding/xml"
	"testing"

	"github.com/c0nscience/yastgt/pkg/svg"
	"github.com/stretchr/testify/assert"
)

func Test_Path(t *testing.T) {
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
		</svg>
	`

	// when
	subj := svg.SVG{}
	err := xml.Unmarshal([]byte(data), &subj)

	// then
	assert.NoError(t, err)
	g := subj.G
	assert.Len(t, g, 1)
	p := g[0].Path[0]
	assert.Equal(t, p.D, "M 1.5420259,10.163793 10.4375,31.906968")

	t.Run("parse path", func(t *testing.T) {
		// when
		p.Parse()

		t.Run("validate M", func(t *testing.T) {
			// then
			assert.Len(t, p.M, 2)
			assert.Equal(t, p.M[0], svg.Point{X: 1.5420259, Y: 10.163793})
			assert.Equal(t, p.M[1], svg.Point{X: 10.4375, Y: 31.906968})
		})
	})
}
