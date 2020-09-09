package parser_test

import (
	"testing"

	"github.com/c0nscience/yastgt/pkg/parser"
	"github.com/c0nscience/yastgt/pkg/svg"
	"github.com/stretchr/testify/assert"
)

func Test_PathParser(t *testing.T) {
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

	subj, err := parser.SVGfrom(data)

	// then
	assert.NoError(t, err)
	g := subj.G[0]
	path := g.Path[0]

	assert.Len(t, path.M, 2)
	assert.Equal(t, path.M[0], svg.Point{X: 1.5420259, Y: 10.163793})
	assert.Equal(t, path.M[1], svg.Point{X: 10.4375, Y: 31.906968})

}
