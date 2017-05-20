package main

type AccumulatingCounter interface {
	Counter

	Counted() int
	addCounted(int)
}

type Counter interface {
	Count(line string) int
}
