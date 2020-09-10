package xml

type SVG struct {
	Path []Path `xml:"path"`
	G    []G    `xml:"g"`
}

type G struct {
	Path []Path `xml:"path"`
}

type Path struct {
	D string `xml:"d,attr"`
}
