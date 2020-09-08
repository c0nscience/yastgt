package svg

type SVG struct {
	G []G `xml:"g"`
}

type G struct {
	Path []Path `xml:"path"`
}
