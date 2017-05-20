package counters

// ConcurrentAccumulatingCounter is an extension of AccumulatingCounter. It
// should work concurrently and allow its caller to Wait.
type ConcurrentAccumulatingCounter interface {
	AccumulatingCounter
	Wait()
}

// AccumulatingCounter can Count for each line and add the result to its
// internal counter.
type AccumulatingCounter interface {
	Count(line string)
	Counted() int
	addCounted(int)
}
