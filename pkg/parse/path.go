package parse

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

var re = regexp.MustCompile(`[MLHVZCmlhvcz]\s{0,1}[\-0-9,\. e]*`)

func Path(p xml.Path) svg.Path {
	d := p.D
	res := svg.Path{
		Points: []svg.PointI{},
	}

	for len(d) > 0 {
		m := re.FindString(d)
		prts := strings.Split(strings.Trim(m, " "), " ")
		cp := cp(res.Points)
		switch prts[0] {
		case "M":
			for i, prt := range prts[1:] {
				if i == 0 {
					res.Points = append(res.Points, MoveTo(prt))
				} else {
					res.Points = append(res.Points, Point(prt))
				}
			}
		case "L":
			for _, prt := range prts[1:] {
				res.Points = append(res.Points, Point(prt))
			}
		case "m":
			for i, prt := range prts[1:] {
				if i == 0 {
					res.Points = append(res.Points, RelMoveTo(prt))
				} else {
					res.Points = append(res.Points, RelPoint(prt))
				}
			}
		case "l":
			for _, prt := range prts[1:] {
				res.Points = append(res.Points, RelPoint(prt))
			}
		case "H":
			for _, prt := range prts[1:] {
				x, _ := strconv.ParseFloat(prt, 64)
				res.Points = append(res.Points, svg.Point{X: x, Y: cp.Y})
			}
		case "h":
			for _, prt := range prts[1:] {
				x, _ := strconv.ParseFloat(prt, 64)
				res.Points = append(res.Points, svg.Point{X: x, Y: cp.Y, Rel: true})
			}
		case "V":
			for _, prt := range prts[1:] {
				y, _ := strconv.ParseFloat(prt, 64)
				res.Points = append(res.Points, svg.Point{X: cp.X, Y: y})
			}
		case "v":
			for _, prt := range prts[1:] {
				y, _ := strconv.ParseFloat(prt, 64)
				res.Points = append(res.Points, svg.Point{X: cp.X, Y: y, Rel: true})
			}
		case "Z", "z":
			fp := res.Points[0].CurrPt()
			fp.MoveTo = false
			fp.Rel = false
			res.Points = append(res.Points, fp)
		case "C":
			rst := prts[1:]
			for i := 0; i < (len(rst) / 3); i++ {
				cp := svg.CubicPoint{
					P1: Point(rst[i*3]),
					P2: Point(rst[i*3+1]),
					CP: Point(rst[i*3+2]),
				}
				res.Points = append(res.Points, cp)
			}
		case "c":
			rst := prts[1:]
			for i := 0; i < (len(rst) / 3); i++ {
				cp := svg.CubicPoint{
					P1:  Point(rst[i*3]),
					P2:  Point(rst[i*3+1]),
					CP:  Point(rst[i*3+2]),
					Rel: true,
				}
				res.Points = append(res.Points, cp)
			}
		}

		d = d[len(m):]

		if len(m) == 0 {
			d = ""
		}
	}

	return res
}

func cp(arr []svg.PointI) svg.Point {
	if len(arr) == 0 {
		return svg.Point{}
	}
	lst := arr[len(arr)-1]
	return lst.CurrPt()
}
