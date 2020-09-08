package svg

import (
	"strconv"
	"strings"
)

type Path struct {
	D string `xml:"d,attr"`

	M []Point
	L []Point
	H []float64
	V []float64
}

func (me *Path) Parse() {
	parts := strings.Split(me.D, " ")

	isM := false //TODO: this is not scalable ... think of something diferent
	for _, p := range parts {
		switch p {
		case "M":
			isM = true
		default:
			if isM {
				cords := strings.Split(p, ",")
				x, _ := strconv.ParseFloat(cords[0], 64)
				y, _ := strconv.ParseFloat(cords[1], 64)
				pt := Point{X: x, Y: y}
				me.M = append(me.M, pt)
			}
		}
	}
}
