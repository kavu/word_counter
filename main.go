package main

import (
	"flag"
	"io"
	"log"
	"os"
)

var (
	searchString  string
	maxProcessors int
	source        io.Reader = os.Stdin
)

func init() {
	flag.StringVar(&searchString, "s", "Go", "A phrase to search")
	flag.IntVar(&maxProcessors, "k", 5, "Maximum concurrent running URL processors")
	// TODO: Implement flag for file openning
}

func main() {
	flag.Parse()

	cntr := NewStringCounter(searchString)
	proc := NewProcessor(maxProcessors, source, cntr)

	proc.Start()

	log.Println(proc.GetResult())
}
