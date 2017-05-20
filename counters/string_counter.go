package counters

import "strings"

type StringCounter struct {
	searchString string
	counted      int
}

func NewStringCounter(search string) *StringCounter {
	return &StringCounter{searchString: search}
}

func (counter *StringCounter) Count(line string) {
	if len(line) == 0 {
		return
	}

	n := strings.Count(line, counter.searchString)

	counter.addCounted(n)
}

func (counter *StringCounter) Counted() int {
	return counter.counted
}

func (counter *StringCounter) addCounted(n int) {
	counter.counted += n
}

func (counter *StringCounter) Wait() {
}
