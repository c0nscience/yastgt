package generate

import (
	"fmt"
	"image/png"
	"math"
	"os"
	"sort"

	"github.com/c0nscience/yastgt/pkg/parse/svg"
)

var dpi = 96.0
var gap = 10.0

func SetGap(f float64) {
	gap = f
}

func FromPNG(f *os.File) []svg.Path {
	img, _ := png.Decode(f)
	bounds := img.Bounds()
	l := map[int][]*lineBounds{}
	gapAsPx := mmToPX(gap)
	bwd := false
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + gapAsPx {
		fnd := false
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			s := uint8(r)
			if s == 255 {
				lbs, ok := l[y]
				fnd = true
				if !ok {
					l[y] = []*lineBounds{}
				}

				if len(lbs) == 0 {
					lbs = append(l[y], &lineBounds{from: x, to: x})
					l[y] = lbs
				}

				lb := lbs[len(lbs)-1]
				if lb.done {
					nlb := &lineBounds{from: x, to: x}
					l[y] = append(l[y], nlb)
					lb = nlb
				}

				if lb.to < x {
					lb.to = x
				}
			} else {
				lbs, ok := l[y]
				if ok {
					lb := lbs[len(lbs)-1]
					if !lb.done {
						lb.done = true
					}
				}
			}
		}

		if fnd {
			if bwd {
				rev := make([]*lineBounds, len(l[y]))
				for i, lb := range l[y] {
					lb.flip()
					rev[len(l[y])-i-1] = lb
				}
				l[y] = rev
			}
			bwd = !bwd
		}

	}

	res := []svg.Path{}

	for y, v := range l {
		ymm := pxToMM(y)
		for _, lb := range v {
			res = append(res, svg.Path{
				Points: []svg.PointI{
					svg.Point{X: pxToMM(lb.from), Y: ymm, MoveTo: true},
					svg.Point{X: pxToMM(lb.to), Y: ymm},
				},
			})
		}
	}

	//TODO: connect paths
	sort.Slice(res, byY(res))

	return res
}

type lineBounds struct {
	from int
	to   int
	done bool
}

func (me *lineBounds) String() string {
	return fmt.Sprintf("[%d,%d]", me.from, me.from)
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

func byY(a []svg.Path) func(int, int) bool {
	return func(i, j int) bool {
		return a[i].Points[0].CurrPt().Y < a[j].Points[0].CurrPt().Y
	}
}
