package main

import (
	"flag"
	"log"
)

var (
	searchString  string
	maxProcessors int64
)

func init() {
	flag.StringVar(&searchString, "s", "Go", "A phrase to search")
	flag.Int64Var(&maxProcessors, "k", 5, "Maximum concurrent running URL processors")
}

func main() {
	flag.Parse()

	log.Println(searchString)
	log.Println(maxProcessors)
}
