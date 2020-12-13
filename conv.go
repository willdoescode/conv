package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/willdoescode/conv/input"
	"github.com/willdoescode/conv/utils"
)

func init() {
	if input.Input == "" {
		color.Error.Println("Please give an input measurement")
		os.Exit(0)
		// Exit with 1 displays error code in terminal
	}
	if input.Output == "" {
		color.Error.Println("Please give an output measurement type")
		os.Exit(0)
	}
}

func main() {
	for _, v := range utils.ReverseSlice(strings.Split(input.Input, "")).([]string) { /* Cannot use range for interface{} */
		fmt.Printf("%s: %v\n", v, utils.IsFloat(v))
	}
}
