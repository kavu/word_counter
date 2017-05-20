package main

type ConcurrentAccumulatingCounter interface {
	Wait()

	AccumulatingCounter
}

type AccumulatingCounter interface {
	Count(line string)
	Counted() int
	addCounted(int)
}
