package pattern

import (
	"image"
	"image/color"
	"math"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
	"github.com/c0nscience/yastgt/pkg/unit"
)

var gap = 10.0
var dpi = 96.0
var degrees = 45.0
var clr = color.Color(color.NRGBA{R: 255})
var threshold = 4.0

func SetGap(f float64) {
	gap = f
}

func SetDpi(f float64) {
	dpi = f
}

func SetDegrees(f float64) {
	degrees = f
}

func SetColor(c color.Color) {
	clr = c
}

func SetThreshold(f float64) {
	threshold = f
}

var maxX float64
var maxY float64
var bwd bool

func Diagonal(img image.Image) []svg.Path {
	bounds := img.Bounds()
	maxX = float64(bounds.Max.X)
	maxY = float64(bounds.Max.Y)
	bwd = false

	var lines [][]*line
	if degrees <= 90 {
		lines = clockwise(img, bounds)
	} else {
		lines = counterClockwise(img, bounds)
	}

	res := []svg.Path{}
	pxToMM := unit.PxToMM(dpi)
	for _, trace := range lines {
		for _, line := range trace {
			res = append(res, svg.Path{
				Points: []svg.PointI{
					svg.Point{X: pxToMM(px(line.start.x)), Y: pxToMM(px(line.start.y)), MoveTo: true},
					svg.Point{X: pxToMM(px(line.end.x)), Y: pxToMM(px(line.end.y))},
				},
			})

		}
	}

	return res
}

func clockwise(img image.Image, bounds image.Rectangle) [][]*line {
	dist := distance(gap, math.Abs(90-degrees))
	distPx := unit.MmToPX(dist, dpi)

	lines := walk(
		(bounds.Max.Y/distPx)*distPx,
		func(t int) bool { return t >= bounds.Min.Y },
		func(i int) vector { return vector{y: float64(i)} },
		func(i int) int { return i - distPx },
		img,
	)

	dist = distance(gap, degrees)
	distPx = unit.MmToPX(dist, dpi)
	lines = append(lines, walk(
		distPx,
		func(t int) bool {
			return t >= bounds.Min.X && t <= bounds.Max.X
		},
		func(i int) vector { return vector{x: float64(i)} },
		func(i int) int { return i + distPx },
		img,
	)...)
	return lines
}

func counterClockwise(img image.Image, bounds image.Rectangle) [][]*line {
	dist := distance(gap, math.Abs(90-degrees))
	distPx := unit.MmToPX(dist, dpi)

	lines := walk(
		(bounds.Max.Y/distPx)*distPx,
		func(t int) bool { return t >= bounds.Min.Y },
		func(i int) vector { return vector{x: float64(bounds.Max.X), y: float64(i)} },
		func(i int) int { return i - distPx },
		img,
	)

	dist = distance(gap, 180-degrees)
	distPx = unit.MmToPX(dist, dpi)
	lines = append(lines, walk(
		bounds.Max.X-distPx,
		func(t int) bool { return t >= bounds.Min.X && t <= bounds.Max.X },
		func(i int) vector { return vector{x: float64(i)} },
		func(i int) int { return i - distPx },
		img,
	)...)
	return lines
}

func walk(min int, cond func(int) bool, v func(int) vector, incr func(int) int, img image.Image) [][]*line {
	lines := [][]*line{}
	for t := min; cond(t); t = incr(t) {
		res := trace(v(t), img)
		if len(res) > 0 {
			if bwd {
				rev := make([]*line, len(res))
				for i, l := range res {
					l.invertDirection()
					rev[len(res)-i-1] = l
				}
				res = rev
			}
			lines = append(lines, res)
			bwd = !bwd
		}
	}
	return lines
}

func trace(start vector, img image.Image) []*line {
	curr := []*line{}
	dir := fromAngle(degrees)
	for p := start; p.x >= float64(img.Bounds().Min.X) && p.y >= float64(img.Bounds().Min.Y) && p.x <= maxX && p.y <= maxY; p = p.plus(dir) {
		c := img.At(px(p.x), px(p.y))
		if clr == c {
			if len(curr) == 0 {
				curr = append(curr, &line{start: p, end: p})
			}

			l := curr[len(curr)-1]

			if l.done {
				l = &line{
					start: p,
					end:   p,
					done:  false,
				}
				curr = append(curr, l)
			}

			if l.end != p {
				l.end = p
			}
		} else {
			if len(curr) == 0 {
				continue
			}

			l := curr[len(curr)-1]
			if l.done {
				continue
			}

			if l.length() < float64(unit.MmToPX(threshold, dpi)) {
				curr = curr[:len(curr)-1]
				continue
			}

			l.done = true
		}
	}
	return curr
}

func distance(dist, deg float64) float64 {
	return dist / math.Sin(unit.DegToRad(deg))
}

func px(f float64) int {
	return int(math.Round(f))
}

type line struct {
	start vector
	end   vector
	done  bool
}

func (l *line) length() float64 {
	return math.Sqrt(math.Pow(l.start.x-l.end.x, 2) + math.Pow(l.start.y-l.end.y, 2))
}

func (l *line) invertDirection() {
	tmp := l.start
	l.start = l.end
	l.end = tmp
}

type vector struct {
	x float64
	y float64
}

func fromAngle(d float64) vector {
	return vector{
		x: precision(math.Cos(unit.DegToRad(d)), 5),
		y: precision(math.Sin(unit.DegToRad(d)), 5),
	}
}

func precision(f float64, i int) float64 {
	return math.Round(f*math.Pow10(i)) / math.Pow10(i)
}

func (v vector) plus(v2 vector) vector {
	return vector{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}
}
