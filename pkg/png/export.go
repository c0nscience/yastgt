package png

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

var dpi = 96
var inkscapePath = ""

func SetDpi(i int) {
	dpi = i
}

func SetInkscapePath(p string) {
	inkscapePath = p
}

func Export(svg string) (*os.File, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "svgexport-*.png")
	if err != nil {
		fmt.Printf("Could not create temp file. Error: %s\n", err.Error())
		return nil, err
	}

	inkscape := inkscapeCmd()

	if inkscape == "" {
		return nil, errors.New("Operating system not supported, could not create inkscape command")
	}

	cmd := exec.Command(
		inkscapeCmd(),
		fmt.Sprintf("--export-filename=%s", tmpFile.Name()),
		fmt.Sprintf("--export-dpi=%d", dpi),
		svg,
	)

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Could not execute inkscape. Error: %s\n", err.Error())
		return nil, err
	}

	return tmpFile, nil
}

func inkscapeCmd() string {
	if inkscapePath != "" {
		return inkscapePath
	}

	switch runtime.GOOS {
	case "windows":
		return "C:\\Program Files\\Inkscape\\bin\\inkscape.com"
	case "linux", "darwin":
		return "inkscape"
	default:
		return ""
	}
}
