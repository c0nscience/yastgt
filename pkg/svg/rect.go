package svg

// <rect x="1" y="1" width="1198" height="398"
// 		fill="none" stroke="blue" stroke-width="2"/>

type Rect struct {
	X      float64 `xml:"x:attr"`
	Y      float64 `xml:"y:attr"`
	Width  float64 `xml:"width:attr"`
	Height float64 `xml:"height:attr"`
}
