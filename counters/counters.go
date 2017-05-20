package counters

type ConcurrentAccumulatingCounter interface {
	AccumulatingCounter
	Wait()
}

type AccumulatingCounter interface {
	Count(line string)
	Counted() int
	addCounted(int)
}
