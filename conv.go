package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/willdoescode/conv/input"
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
	fmt.Println(input.Input, input.Output)
}
