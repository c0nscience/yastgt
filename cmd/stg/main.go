package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/reader"
	"github.com/c0nscience/yastgt/pkg/transform"
)

func main() {
	args := os.Args
	b, err := ioutil.ReadFile(args[1])

	x, err := reader.Unmarshal(b)
	if err != nil {
		log.Fatalf("Could not read file %s. Error: %s", args[1], err.Error())
	}

	s := parse.ToSvg(x)

	cmds := transform.ToGcode(s)

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
