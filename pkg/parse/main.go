package parse

import (
	"strconv"

	"gonum.org/v1/gonum/mat"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/parse/transform"
	"github.com/c0nscience/yastgt/pkg/reader/xml"
)

func SVG(xml xml.SVG) svg.SVG {
	h, _ := strconv.ParseFloat(xml.Height[:len(xml.Height)-2], 64)
	p := []svg.PointI{}
	p = append(p, fromPath(xml.Path, []*mat.Dense{})...)
	p = append(p, fromLine(xml.Line, []*mat.Dense{})...)
	p = append(p, fromRect(xml.Rect, []*mat.Dense{})...)
	p = append(p, fromCircle(xml.Circle, []*mat.Dense{})...)
	p = append(p, fromGroup(xml.G, []*mat.Dense{})...)

	return svg.SVG{
		Height: h,
		Points: p,
	}
}

func fromGroup(grps []xml.G, trans []*mat.Dense) []svg.PointI {
	res := []svg.PointI{}
	for _, g := range grps {
		t := trans
		t = append(t, transform.ParseTypes(g.Transform)...)
		res = append(res, fromGroup(g.G, t)...)
		res = append(res, fromPath(g.Path, t)...)
		res = append(res, fromLine(g.Line, t)...)
		res = append(res, fromRect(g.Rect, t)...)
		res = append(res, fromCircle(g.Circle, t)...)
	}

	return res
}

func fromPath(pths []xml.Path, trans []*mat.Dense) []svg.PointI {
	res := []svg.PointI{}
	for _, p := range pths {
		pts := Path(p)
		transform.Apply(append(trans, transform.ParseTypes(p.Transform)...), pts)
		res = append(res, pts...)
	}
	return res
}

func fromLine(lines []xml.Line, trans []*mat.Dense) []svg.PointI {
	res := []svg.PointI{}
	for _, l := range lines {
		pts := Line(l)
		transform.Apply(append(trans, transform.ParseTypes(l.Transform)...), pts)
		res = append(res, pts...)
	}

	return res
}

func fromRect(rects []xml.Rect, trans []*mat.Dense) []svg.PointI {
	res := []svg.PointI{}
	for _, r := range rects {
		pts := Rect(r)
		transform.Apply(append(trans, transform.ParseTypes(r.Transform)...), pts)
		res = append(res, pts...)
	}
	return res
}

func fromCircle(circles []xml.Circle, trans []*mat.Dense) []svg.PointI {
	res := []svg.PointI{}
	for _, c := range circles {
		pt := Circle(c)
		transform.Apply(append(trans, transform.ParseTypes(c.Transform)...), []svg.PointI{pt})
		res = append(res, pt)
	}
	return res
}
