package input

import (
	"flag"
	"strings"
)

var (
	Input  string
	Output string
)

func init() {
	flag.StringVar(&Input, "i", "", "An input value")
	flag.StringVar(&Output, "o", "", "An output measurement type")
	flag.Parse()
	Input = strings.Trim(Input, " \n")
	Output = strings.Trim(Output, " \n")
}
