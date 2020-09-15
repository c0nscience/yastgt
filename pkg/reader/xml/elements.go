package xml

type SVG struct {
	Path []Path `xml:"path"`
	Rect []Rect `xml:"rect"`
	G    []G    `xml:"g"`
}

type G struct {
	Rect []Rect `xml:"rect"`
	Path []Path `xml:"path"`
}

type Path struct {
	D string `xml:"d,attr"`
}

type Rect struct {
	X      float64 `xml:"x,attr"`
	Y      float64 `xml:"y,attr"`
	Height float64 `xml:"height,attr"`
	Width  float64 `xml:"width,attr"`
}
