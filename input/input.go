package input

import "github.com/alexflint/go-arg"

var Args struct {
	Types string `arg:"positional"`
}

func init() {
	arg.MustParse(&Args)
}
