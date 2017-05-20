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

	counter := NewURLCounter(searchString)
	processor := NewProcessor(maxProcessors, source, counter)

	processor.Start()
	processor.Wait()

	log.Println(processor.GetResult())
}
