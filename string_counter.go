package main

import "strings"

type StringCounter struct {
	searchString string
	counted      int
}

func NewStringCounter(search string) *StringCounter {
	return &StringCounter{searchString: search}
}

func (counter *StringCounter) Count(line string) int {
	if len(line) == 0 {
		return 0
	}

	n := strings.Count(line, counter.searchString)

	counter.addCounted(n)

	return n
}

func (counter *StringCounter) Counted() int {
	return counter.counted
}

func (counter *StringCounter) addCounted(n int) {
	counter.counted += n
}
