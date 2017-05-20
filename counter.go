package main

import "strings"

type Counter interface {
	Count(line string) int
}

type URLCounter struct {
	s string
}

func (cntr *URLCounter) Count(line string) int {
	if len(line) == 0 {
		return 0
	}

	return strings.Count(line, cntr.s)
}
