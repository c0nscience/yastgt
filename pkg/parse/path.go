package parse

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

var re = regexp.MustCompile(`[MLHVCZQTAmlhvczqta]\s?[\-0-9,\. e]*`)

func Path(p xml.Path) []svg.PointI {
	d := p.D
	res := []svg.PointI{}

	var initSubPathPt *svg.Point
	wasClosed := false
	for len(d) > 0 {
		m := re.FindString(d)
		prts := splitIntoPairs(m)
		currPt := cp(res)
		if wasClosed {
			initSubPathPt = currPt
			wasClosed = false
		}
		switch prts[0] {
		case "M":
			for i, prt := range prts[1:] {
				p := Point(prt)
				if i == 0 {
					p = MoveTo(prt)
					initSubPathPt = p
				}
				res = append(res, p)
			}
		case "L":
			for _, prt := range prts[1:] {
				res = append(res, Point(prt))
			}
		case "H":
			for _, prt := range prts[1:] {
				x, _ := strconv.ParseFloat(prt, 64)
				y := 0.0
				if currPt != nil {
					y = currPt.Y
				}
				res = append(res, &svg.Point{X: x, Y: y})
			}
		case "V":
			for _, prt := range prts[1:] {
				y, _ := strconv.ParseFloat(prt, 64)
				res = append(res, &svg.Point{X: currPt.X, Y: y})
			}
		case "Z":
			if initSubPathPt != nil {
				p := &svg.Point{
					X: initSubPathPt.CurrPt().X,
					Y: initSubPathPt.CurrPt().Y,
				}
				currPt = p
				res = append(res, p)
				wasClosed = true
			}
		case "C":
			rst := prts[1:]
			for i := 0; i < (len(rst) / 3); i++ {
				cp := &svg.CubicPoint{
					P1: Point(rst[i*3]),
					P2: Point(rst[i*3+1]),
					CP: Point(rst[i*3+2]),
				}
				res = append(res, cp)
			}
		default:
			fmt.Printf("Could not match: %s\n", strings.Join(prts, " "))
		}

		d = d[len(m):]

		if len(m) == 0 {
			d = ""
		}
	}

	return res
}

func cp(arr []svg.PointI) *svg.Point {
	if len(arr) == 0 {
		return nil
	}
	lst := arr[len(arr)-1]
	return lst.CurrPt()
}

func splitIntoPairs(m string) []string {
	m = strings.Trim(m, " ")
	prts := strings.Split(m, " ")

	if strings.ContainsAny(m, "HVLhvl") {
		return prts
	}

	if strings.Contains(m, ",") {
		return prts
	}
	res := []string{}
	if (len(prts)-1)%2 == 0 {
		res = append(res, prts[0])
		tmp := prts[1:]
		for i := 0; i < len(tmp); i = i + 2 {
			res = append(res, strings.Join(tmp[i:i+2], ","))
		}
	}
	return res
}
