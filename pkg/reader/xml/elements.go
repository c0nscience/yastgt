package xml

type SVG struct {
	Height string   `xml:"height,attr"`
	Path   []Path   `xml:"path"`
	Rect   []Rect   `xml:"rect"`
	Line   []Line   `xml:"line"`
	Circle []Circle `xml:"circle"`
	G      []G      `xml:"g"`
}

type G struct {
	Path      []Path   `xml:"path"`
	Rect      []Rect   `xml:"rect"`
	Line      []Line   `xml:"line"`
	Circle    []Circle `xml:"circle"`
	G         []G      `xml:"g"`
	Transform string   `xml:"transform,attr"`
}

type Path struct {
	D         string `xml:"d,attr"`
	Transform string `xml:"transform,attr"`
}

type Rect struct {
	X         float64 `xml:"x,attr"`
	Y         float64 `xml:"y,attr"`
	Height    float64 `xml:"height,attr"`
	Width     float64 `xml:"width,attr"`
	Transform string  `xml:"transform,attr"`
}

type Line struct {
	X1        float64 `xml:"x1,attr"`
	Y1        float64 `xml:"y1,attr"`
	X2        float64 `xml:"x2,attr"`
	Y2        float64 `xml:"y2,attr"`
	Transform string  `xml:"transform,attr"`
}

type Circle struct {
	CX        float64 `xml:"cx,attr"`
	CY        float64 `xml:"cy,attr"`
	R         float64 `xml:"r,attr"`
	Transform string  `xml:"transform,attr"`
}
