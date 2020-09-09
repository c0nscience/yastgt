package svg

type Path struct {
	D string `xml:"d,attr"`

	M []Point
	L []Point
	H []float64
	V []float64
}

//TODO: extract a populator ??
func (me *Path) Populate() {
	// parts := strings.Split(me.D, " ")

	// arr := func() *[]interface{} { return nil }
	// //TODO extract the parts differently

	// for _, p := range parts {
	// 	switch p {
	// 	case "M":
	// 		arr = func() *[]interface{} {
	// 			return &me.M
	// 		}
	// 	default:
	// 		cords := strings.Split(p, ",")
	// 		x, _ := strconv.ParseFloat(cords[0], 64)
	// 		y, _ := strconv.ParseFloat(cords[1], 64)
	// 		pt := Point{X: x, Y: y}
	// 		a := []Point(arr())
	// 		a = append(*a, pt)
	// 	}
	// }
}
