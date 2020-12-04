package main

import (
	"fmt"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/c0nscience/yastgt/pkg/fillpng"
	"github.com/c0nscience/yastgt/pkg/parse"
	"github.com/c0nscience/yastgt/pkg/pattern"
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
	flagFill         = "fill"
	//flagPadding      = "padding"
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
			&cli.StringSliceFlag{Name: flagFill, Value: nil, Usage: "Defines the fill pattern via 'degrees,red,green,blue' with the color values in standard 0-255. If not specified no fill data is generated."},
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
			fills := c.StringSlice(flagFill)

			b, err := ioutil.ReadFile(svgFilePath)
			if err != nil {
				return err
			}
			x, _ := reader.Unmarshal(b) //svg.Read

			s := parse.SVG(x) //svg.Parse

			if fills != nil {
				fillpng.SetDpi(int(dpi))
				fillpng.SetInkscapePath(inkscapePath)
				f, err := fillpng.Export(svgFilePath)
				if err != nil {
					return err
				}
				defer os.Remove(f.Name())

				img, err := png.Decode(f)
				if err != nil {
					fmt.Printf("Coud not decode image. Error: %s", err.Error())
				}

				pattern.SetGap(gap)
				pattern.SetDpi(dpi)
				pattern.SetThreshold(threshold)
				for _, str := range fills {
					prts := strings.Split(str, ",")
					degrees, _ := strconv.ParseFloat(prts[0], 64)
					r, _ := strconv.ParseInt(prts[1], 10, 64)
					g, _ := strconv.ParseInt(prts[2], 10, 64)
					b, _ := strconv.ParseInt(prts[3], 10, 64)

					pattern.SetColor(color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255})
					pattern.SetDegrees(degrees)
					s.Path = append(s.Path, pattern.Diagonal(img)...)
				}
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
