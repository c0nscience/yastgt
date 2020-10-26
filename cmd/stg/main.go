package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/c0nscience/yastgt/pkg/generate"
	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/reader"
	"github.com/c0nscience/yastgt/pkg/transform"
)

const (
	flagSvgFilePath = "svg"
	flagOutFilePath = "out"
	flagFillPNGFile = "fill"
	flagCurveSpeed  = "curveSpeed"
	flagLinearSpeed = "linearSpeed"
	flagGap         = "gap"
	flagThreshold   = "threshold"
)

func main() {

	app := &cli.App{
		Name:  "stg",
		Usage: "Generates GCode from SVG",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: flagSvgFilePath, Usage: "Path to the SVG file to generate GCode from.", Required: true},
			&cli.StringFlag{Name: flagOutFilePath, Usage: "Path to the output GCode file.", Required: true},
			&cli.StringFlag{Name: flagFillPNGFile, Usage: "PNG file containing the fill information."},
			&cli.Float64Flag{Name: flagCurveSpeed, Value: 3000.0, Usage: "Divisor to normalize the speed of curves."},
			&cli.Float64Flag{Name: flagLinearSpeed, Value: 4000.0, Usage: "Flat feed value for linear move commands."},
			&cli.Float64Flag{Name: flagGap, Value: 10.0, Usage: "Gap between fill lines."},
			&cli.Float64Flag{Name: flagThreshold, Value: 4.0, Usage: "Minimum line length for fill pattern."},
		},
		Action: func(c *cli.Context) error {
			curveSpeed := c.Float64(flagCurveSpeed)
			linearSpeed := c.Float64(flagLinearSpeed)
			svgFilePath := c.String(flagSvgFilePath)
			fillFilePath := c.String(flagFillPNGFile)
			outFilePath := c.String(flagOutFilePath)
			gap := c.Float64(flagGap)
			threshold := c.Float64(flagThreshold)

			b, err := ioutil.ReadFile(svgFilePath)
			if err != nil {
				return err
			}
			x, _ := reader.Unmarshal(b)

			s := parse.SVG(x)

			if len(fillFilePath) > 0 {
				f, err := os.Open(fillFilePath)
				if err != nil {
					return err
				}
				generate.SetGap(gap)
				generate.SetThreshold(threshold)
				fill := generate.FromPNG(f)
				s.Path = append(s.Path, fill...)
			}

			transform.SetG0Speed(linearSpeed)
			transform.SetG5Speed(curveSpeed)
			cmds := transform.Gcode(s)

			out, err := os.Create(outFilePath)
			if err != nil {
				return err
			}

			for _, c := range cmds {
				out.WriteString(string(c))
				out.WriteString("\n")
			}

			err = out.Close()
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
