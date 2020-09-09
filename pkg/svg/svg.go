package svg

type SVG struct {
	G []*G `xml:"g"`
}

func (me *SVG) Populate() {
	for _, g := range me.G {
		g.Populate()
	}
}

type G struct {
	Path []*Path `xml:"path"`
}

func (me *G) Populate() {
	for _, p := range me.Path {
		p.Populate()
	}
}
