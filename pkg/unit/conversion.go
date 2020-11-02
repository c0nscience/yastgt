package unit

import "math"

func pxToInch(px int, dpi float64) float64 {
	return float64(px) / dpi
}

func inchToMM(inch float64) float64 {
	return inch * 25.4
}

func PxToMM(dpi float64) func(int) float64 {
	return func(px int) float64 {
		return inchToMM(pxToInch(px, dpi))
	}
}

func mmToInch(mm float64) float64 {
	return mm / 25.4
}

func inchToPX(inch, dpi float64) int {
	return int(math.Round(inch * dpi))
}

func MmToPX(mm float64, dpi float64) int {
	return inchToPX(mmToInch(mm), dpi)
}

func DegToRad(d float64) float64 {
	return d * (math.Pi / 180)
}
