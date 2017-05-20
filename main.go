package main

import (
	"flag"
	"io"
	"os"
)

var (
	searchString string
	maxJobs      int
	source       io.Reader = os.Stdin
)

func init() {
	flag.StringVar(&searchString, "s", "Go", "A phrase to search")
	flag.IntVar(&maxJobs, "k", 5, "Maximum concurrent running URL processors")
	// TODO: Implement flag for file openning
}

func main() {
	flag.Parse()

	counter := NewURLCounter(searchString, maxJobs)
	processor := NewProcessor(source, counter)

	processor.Start()
	processor.Wait()
}
