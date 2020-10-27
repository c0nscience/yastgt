package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/c0nscience/yastgt/pkg/generate"
	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/png"
	"github.com/c0nscience/yastgt/pkg/reader"
	"github.com/c0nscience/yastgt/pkg/transform"
)

const (
	flagSvgFilePath  = "svg"
	flagOutFilePath  = "out"
	flagCurveSpeed   = "curveSpeed"
	flagLinearSpeed  = "linearSpeed"
	flagGap          = "gap"
	flagThreshold    = "threshold"
	flagDpi          = "dpi"
	flagInkscapePath = "inkscape"
	flagNoFill       = "no-fill"
	flagPadding      = "padding"
)

func main() {

	app := &cli.App{
		Name:  "stg",
		Usage: "Generates GCode from SVG",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: flagSvgFilePath, Usage: "Path to the SVG file to generate GCode from.", Required: true},
			&cli.StringFlag{Name: flagOutFilePath, Usage: "Path to the output GCode file.", Required: true},
			&cli.Float64Flag{Name: flagCurveSpeed, Value: 3000.0, Usage: "Divisor to normalize the speed of curves."},
			&cli.Float64Flag{Name: flagLinearSpeed, Value: 4000.0, Usage: "Flat feed value for linear move commands."},
			&cli.Float64Flag{Name: flagGap, Value: 10.0, Usage: "Gap between fill lines."},
			&cli.Float64Flag{Name: flagThreshold, Value: 4.0, Usage: "Minimum line length for fill pattern."},
			&cli.Float64Flag{Name: flagDpi, Value: 96.0, Usage: "DPI of the rasterized SVG image. Used to calculate the fill pattern."},
			&cli.StringFlag{Name: flagInkscapePath, Value: "", Usage: "The path to a inkscape commandline binary version >= 1.x"},
			&cli.BoolFlag{Name: flagNoFill, Value: false, Usage: "Set to disable filling the shapes with patterns."},
			&cli.Float64Flag{Name: flagPadding, Value: 0.0, Usage: "Set a padding in mm for fill pattern."},
		},
		Action: func(c *cli.Context) error {
			curveSpeed := c.Float64(flagCurveSpeed)
			linearSpeed := c.Float64(flagLinearSpeed)
			svgFilePath := c.String(flagSvgFilePath)
			outFilePath := c.String(flagOutFilePath)
			gap := c.Float64(flagGap)
			threshold := c.Float64(flagThreshold)
			dpi := c.Float64(flagDpi)
			inkscapePath := c.String(flagInkscapePath)
			noFill := c.Bool(flagNoFill)
			padding := c.Float64(flagPadding)

			b, err := ioutil.ReadFile(svgFilePath)
			if err != nil {
				return err
			}
			x, _ := reader.Unmarshal(b)

			s := parse.SVG(x)

			if !noFill {
				png.SetDpi(int(dpi))
				png.SetInkscapePath(inkscapePath)
				f, err := png.Export(svgFilePath)
				if err != nil {
					return err
				}
				defer os.Remove(f.Name())

				generate.SetPadding(padding)
				generate.SetGap(gap)
				generate.SetThreshold(threshold)
				generate.SetDpi(dpi)
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
