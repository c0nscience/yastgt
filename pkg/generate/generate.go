package generate

import (
	"fmt"
	"image/png"
	"math"
	"os"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
)

var dpi = 96.0
var gap = 10.0
var threshold = 4.0

func SetGap(f float64) {
	gap = f
}

func SetThreshold(f float64) {
	threshold = f
}

func SetDpi(f float64) {
	dpi = f
}

func FromPNG(f *os.File) []svg.Path {
	img, _ := png.Decode(f)
	bounds := img.Bounds()
	l := [][]*lineBounds{}
	gapAsPx := mmToPX(gap)
	bwd := false
	curr := []*lineBounds{}
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + gapAsPx {
		fnd := false
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			s := uint8(r)
			if s == 255 {
				if !fnd {
					curr = []*lineBounds{}
				}
				fnd = true

				if len(curr) == 0 {
					curr = append(curr, &lineBounds{y: y, from: x, to: x})
				}

				lb := curr[len(curr)-1]
				if lb.done {
					nlb := &lineBounds{y: y, from: x, to: x}
					curr = append(curr, nlb)
					lb = nlb
				}

				if lb.to < x {
					lb.to = x
				}
			} else {
				if fnd {
					lb := curr[len(curr)-1]
					if !lb.done {
						lb.done = true
					}
				}
			}
		}

		if fnd {
			if bwd {
				rev := make([]*lineBounds, len(curr))
				for i, lb := range curr {
					lb.flip()
					rev[len(curr)-i-1] = lb
				}
				curr = rev
			}
			bwd = !bwd
			l = append(l, curr)
		}

	}

	res := []svg.Path{}

	for _, v := range l {
		for _, lb := range v {
			from := pxToMM(lb.from)
			to := pxToMM(lb.to)
			if math.Abs(from-to) < threshold {
				continue
			}

			res = append(res, svg.Path{
				Points: []svg.PointI{
					svg.Point{X: from, Y: pxToMM(lb.y), MoveTo: true},
					svg.Point{X: to, Y: pxToMM(lb.y)},
				},
			})
		}
	}

	return res
}

type lineBounds struct {
	y    int
	from int
	to   int
	done bool
}

func (me *lineBounds) String() string {
	return fmt.Sprintf("[%d,%d]", me.from, me.to)
}

func (me *lineBounds) flip() {
	tmp := me.from
	me.from = me.to
	me.to = tmp
}

func pxToInch(px int) float64 {
	return float64(px) / dpi
}

func inchToMM(inch float64) float64 {
	return inch * 25.4
}

func pxToMM(px int) float64 {
	return inchToMM(pxToInch(px))
}

func mmToInch(mm float64) float64 {
	return mm / 25.4
}

func inchToPX(inch float64) int {
	return int(math.Ceil(inch * dpi))
}

func mmToPX(mm float64) int {
	return inchToPX(mmToInch(mm))
}
