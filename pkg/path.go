package svg

type Path struct {
	D string `xml:"d,attr"`

	M []Point
	L []Point
	H []float64
	V []float64
}
