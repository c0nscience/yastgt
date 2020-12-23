package transform

import (
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

const (
	matrix    = "matrix"
	translate = "translate"
	scale     = "scale"
	rotate    = "rotate"
	skewX     = "skewX"
	skewY     = "skewY"
)

var re = regexp.MustCompile(`(matrix|translate|scale|rotate|skew[X|Y])\((.*?)\)`)

func ParseTypes(s string) []*mat.Dense {
	res := []*mat.Dense{}

	for len(s) > 0 {
		m, l := parseType(s)
		if m != nil {
			res = append(res, m)
		}
		s = s[l:]
	}

	return res
}

func parseType(s string) (*mat.Dense, int) {
	n, t, l := parse(s)

	var res *mat.Dense
	switch t {
	case matrix:
		res = New(n[0], n[1], n[2], n[3], n[4], n[5])
	case translate:
		res = handleTranslate(n)
	case scale:
		res = handleScale(n)
	case rotate:
		res = handleRotate(n)
	case skewX:
		res = NewSkewX(n[0])
	case skewY:
		res = NewSkewY(n[0])
	default:
		res = nil
	}
	return res, l
}

func parse(s string) ([]float64, string, int) {
	m := re.FindStringSubmatch(s)
	res := []float64{}
	if m == nil {
		return res, "", len(s)
	}
	if len(m) != 3 {
		return res, "", len(m[0])
	}

	prts := []string{}
	if strings.Contains(m[2], ",") {
		prts = strings.Split(m[2], ",")
	} else {
		prts = strings.Split(m[2], " ")
	}
	for _, p := range prts {
		res = append(res, toFloat(p))
	}
	return res, m[1], len(m[0])
}

func toFloat(s string) float64 {
	f, _ := strconv.ParseFloat(strings.Trim(s, " "), 64)
	return f
}

func handleTranslate(n []float64) *mat.Dense {
	if len(n) == 1 {
		return NewTranslation(n[0], 0)
	}

	return NewTranslation(n[0], n[1])
}

func handleScale(n []float64) *mat.Dense {
	if len(n) == 1 {
		return NewScaling(n[0], n[0])
	}
	return NewScaling(n[0], n[1])
}

func handleRotate(n []float64) *mat.Dense {
	if len(n) == 1 {
		return NewRotation(n[0])
	}

	var intRes mat.Dense
	intRes.Mul(NewTranslation(n[1], n[2]), NewRotation(n[0]))
	var res mat.Dense
	res.Mul(&intRes, NewTranslation(-n[1], -n[2]))
	return &res
}
