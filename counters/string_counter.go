package counters

import "strings"

// StringCounter counts a string inside the lines passed to its Count method.
// It implements ConcurrentAccumulatingCounter, but it is not concurrent.
// Its Wait function does nothing.
type StringCounter struct {
	searchString string // string to count
	counted      int    // accumulator
}

// NewStringCounter is just a factory function for StringCounter. You need to
// pass a string you want to count.
func NewStringCounter(search string) *StringCounter {
	return &StringCounter{searchString: search}
}

// Count counts searchString in the line and stores the value inside
// accumulator.
func (counter *StringCounter) Count(line string) {
	if len(line) == 0 {
		return
	}

	n := strings.Count(line, counter.searchString)

	counter.addCounted(n)
}

// Counted returns accumulator value.
func (counter *StringCounter) Counted() int {
	return counter.counted
}

func (counter *StringCounter) addCounted(n int) {
	counter.counted += n
}

// Wait is needed to implement ConcurrentAccumulatingCounter. It does nothing
// here.
func (counter *StringCounter) Wait() {
}
