package svg

type SVG struct {
	Path []*Path `xml:"path"`
	Rect []*Rect `xml:"rect"`
	G    []*G    `xml:"g"`
}

type G struct {
	Path []*Path `xml:"path"`
}
