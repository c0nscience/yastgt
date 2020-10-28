# TODO

[ ] multiple fill patterns
    [ ] horizontal (red)
    [ ] vertical (green)
    [ ] both (blue)
[x] automatically rasterize svg to derive fill pattern
[x] build and deploy binaries to github
[ ] think about pen change

# Yet another SVG to G-code tool

As the name suggest this tool generates g-code for our xy-plotter directly from an SVG file with the following features:
* use a servo motor to handle the pen
* use simple non feeding motion g-codes
* support lines, circles, radis, filled shapes
* simple optimisation

> **DISCLAIMER**
> Beware that this is a specialized tool written for our xy-plotter. Feel free none the less to fork this repo and adopt it to your needs or try reaching out to us for support.
>
> Meaning this tool is not meant as a general purpose solution to generate g-code out of svg files.

# Usage

```text
NAME:
   stg - Generates GCode from SVG

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --svg value          Path to the SVG file to generate GCode from.
   --out value          Path to the output GCode file.
   --curveSpeed value   Divisor to normalize the speed of curves. (default: 3000)
   --linearSpeed value  Flat feed value for linear move commands. (default: 4000)
   --gap value          Gap between fill lines. (default: 10)
   --threshold value    Minimum line length for fill pattern. (default: 4)
   --dpi value          DPI of the rasterized SVG image. Used to calculate the fill pattern. (default: 96)
   --inkscape value     The path to a inkscape commandline binary version >= 1.x
   --no-fill            Set to disable filling the shapes with patterns. (default: false)
   --padding value      Set a padding in mm for fill pattern. (default: 0)
   --help, -h           show help (default: false)
```