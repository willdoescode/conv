package input

import "flag"

var (
	Input  string
	Output string
)

func init() {
	flag.StringVar(&Input, "i", "", "An input value")
	flag.StringVar(&Output, "o", "", "An output measurement type")
	flag.Parse()
}
