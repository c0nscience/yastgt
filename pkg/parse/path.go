package parse

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
)

var mRe = regexp.MustCompile(`M [0-9,\. ]*`)

func FindM(s string) string {
	return strings.Trim(mRe.FindString(s), " ")
}

func M(s string) []svg.Point {
	pts := strings.Split(s, " ")[1:]
	res := []svg.Point{}

	for _, pt := range pts {
		res = append(res, Point(pt))
	}

	return res
}

var lRe = regexp.MustCompile(`L [0-9,\. ]*`)

func FindL(s string) string {
	return strings.Trim(lRe.FindString(s), " ")
}

func L(s string) []svg.Point {
	pts := strings.Split(s, " ")[1:]
	res := []svg.Point{}

	for _, pt := range pts {
		res = append(res, Point(pt))
	}

	return res
}

var hRe = regexp.MustCompile(`H [0-9]*\.?[0-9]+`)

func FindH(s string) string {
	return strings.Trim(hRe.FindString(s), " ")
}

func H(s string) (f float64) {
	f, _ = strconv.ParseFloat(strings.Split(s, " ")[1], 64)
	return
}

var vRe = regexp.MustCompile(`V [0-9]*\.?[0-9]+`)

func FindV(s string) string {
	return strings.Trim(vRe.FindString(s), " ")
}

func V(s string) (f float64) {
	f, _ = strconv.ParseFloat(strings.Split(s, " ")[1], 64)
	return
}
