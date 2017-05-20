package main

import (
	"bufio"
	"io"
)

func Processor(max int64, ior io.Reader, cntr Counter) {
	scanner := bufio.NewScanner(ior)
	for scanner.Scan() {
		line := scanner.Text()

		cntr.Count(line)
	}
}
