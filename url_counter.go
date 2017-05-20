package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type URLCounter struct {
	searchString string
	counted      int
	maxJobs      chan bool
	wg           sync.WaitGroup
	sync.RWMutex
}

func NewURLCounter(search string, maxJobs int) *URLCounter {
	ch := make(chan bool, maxJobs)

	return &URLCounter{searchString: search, maxJobs: ch}
}

func (counter *URLCounter) Count(line string) {
	if len(line) == 0 {
		return
	}

	counter.maxJobs <- true
	counter.wg.Add(1)
	go counter.countInHTTPResponse(line)
}

func (counter *URLCounter) countInHTTPResponse(line string) {
	defer counter.jobDone()

	_, err := url.Parse(line)
	if err != nil {
		// Not a valid URL, skipping
		return
	}

	response, err := http.Get(line)
	if err != nil {
		// Something went wrong, skipping
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// Something went wrong, skipping
		return
	}

	n := bytes.Count(body, []byte(counter.searchString))

	log.Printf("Count for %s: %d\n", line, n)

	counter.addCounted(n)
}

func (counter *URLCounter) jobDone() {
	<-counter.maxJobs
	counter.wg.Done()
}

func (counter *URLCounter) addCounted(n int) {
	counter.Lock()
	defer counter.Unlock()
	counter.counted += n
}

func (counter *URLCounter) Counted() int {
	counter.RLock()
	defer counter.RUnlock()
	return counter.counted
}

func (counter *URLCounter) Wait() {
	counter.wg.Wait()
}
