package main

import (
	"strings"
	"sync"
)

type URLCounter struct {
	searchString string
	counted      int
	maxJobsChan  chan bool // FIXME: unused
	wg           sync.WaitGroup
	sync.RWMutex
}

func NewURLCounter(search string) *URLCounter {
	ch := make(chan bool, 5) // FIXME: unused, move from processor

	return &URLCounter{searchString: search, maxJobsChan: ch}
}

func (counter *URLCounter) Count(line string) {
	if len(line) == 0 {
		return
	}

	counter.wg.Add(1)
	go counter.countInHTTPResponse(line)
}

func (counter *URLCounter) countInHTTPResponse(url string) {
	n := strings.Count(url, counter.searchString)

	counter.addCounted(n)
	counter.wg.Done()
}

func (counter *URLCounter) Counted() int {
	counter.RLock()
	defer counter.RUnlock()
	return counter.counted
}

func (counter *URLCounter) addCounted(n int) {
	counter.Lock()
	defer counter.Unlock()
	counter.counted += n
}

func (counter *URLCounter) Wait() {
	counter.wg.Wait()
}
