package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/c0nscience/yastgt/pkg/generate"
	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/reader"
	"github.com/c0nscience/yastgt/pkg/transform"
)

func main() {
	fillInfo := flag.String("f", "", "file with fill information")
	g0 := flag.Int64("g0s", int64(4000), "G0 feed rate")
	g5 := flag.Int64("g5s", int64(100), "G5 feed rate")
	flag.Parse()

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("couldnt read input")
	}
	x, _ := reader.Unmarshal(b)

	s := parse.SVG(x)

	if len(*fillInfo) > 0 {
		f, err := os.Open(*fillInfo)
		if err == nil {
			fill := generate.FromPNG(f)
			s.Path = append(s.Path, fill...)
		}
	}

	transform.SetG0Speed(*g0)
	transform.SetG5Speed(*g5)
	cmds := transform.Gcode(s)

	var buff bytes.Buffer
	for _, c := range cmds {
		buff.WriteString(string(c))
		buff.WriteString("\n")
	}

	os.Stdout.Write(buff.Bytes())
}
