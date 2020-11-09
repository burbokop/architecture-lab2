package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/burbokop/architecture-lab2"
)

var (
	inputExpression = flag.String("e", "", "expression")
	inputFile       = flag.String("f", "", "input file")
	outputFile      = flag.String("o", "", "output file")
)

func main() {
	flag.Parse()
	var in io.Reader
	var out io.Writer

	if *inputExpression != "" {
		in = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, e := os.Open(*inputFile)
		if e != nil {
			fmt.Println(e)
			return
		}
		in = file
	} else {
		fmt.Println("flags help:\n\t-e - expression\n\t-f - input file with expression\n\t-o - output file (default stdout)")
		return
	}

	if *outputFile != "" {
		file, e := os.Create(*outputFile)
		if e != nil {
			fmt.Println(e)
			return
		}
		out = file
	} else {
		out = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  in,
		Output: out,
	}
	err := handler.Compute()

	if err != nil {
		fmt.Println(err)
	}
}
