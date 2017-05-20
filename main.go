package main

import (
	"flag"
	"io"
	"os"
)

var (
	searchString  string
	maxProcessors int64
	source        io.Reader = os.Stdin
)

func init() {
	flag.StringVar(&searchString, "s", "Go", "A phrase to search")
	flag.Int64Var(&maxProcessors, "k", 5, "Maximum concurrent running URL processors")
	// TODO: Implement flag for file openning
}

func main() {
	flag.Parse()

	cntr := &URLCounter{searchString}

	Processor(maxProcessors, source, cntr)
}
