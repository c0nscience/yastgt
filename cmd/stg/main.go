package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/reader"
	"github.com/c0nscience/yastgt/pkg/transform"
)

func main() {
	args := os.Args
	g0 := flag.Int64("g0s", int64(2000), "G0 feed rate")
	g5 := flag.Int64("g5s", int64(100), "G5 feed rate")
	b, err := ioutil.ReadFile(args[1])

	x, err := reader.Unmarshal(b)
	if err != nil {
		log.Fatalf("Could not read file %s. Error: %s", args[1], err.Error())
	}

	s := parse.SVG(x)

	transform.SetG0Speed(*g0)
	transform.SetG5Speed(*g5)
	cmds := transform.Gcode(s)

	var buff bytes.Buffer
	for _, c := range cmds {
		buff.WriteString(string(c))
		buff.WriteString("\n")
	}

	err = ioutil.WriteFile(args[2], buff.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Could not write file %s. Error: %s", args[2], err.Error())
	}
}
