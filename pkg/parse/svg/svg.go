package svg

type SVG struct {
	Path []Path
	G    []G //TODO we can flatten the structure here ...
}

type G struct {
	Path []Path
}
