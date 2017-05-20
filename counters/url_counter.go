package counters

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

// URLCounter counts a string inside the HTTP bodies, which it get from the URLs
// passed to its Count method. It implements ConcurrentAccumulatingCounter.
// It have a cap of max concurrent jobs running.
type URLCounter struct {
	searchString string
	counted      int
	maxJobs      chan bool
	wg           sync.WaitGroup
	sync.RWMutex
}

// NewURLCounter is just a factory function for URLCounter. You need to
// pass a string you want to count and max concurrent jobs to run.
func NewURLCounter(search string, maxJobs int) *URLCounter {
	ch := make(chan bool, maxJobs)

	return &URLCounter{searchString: search, maxJobs: ch}
}

// Count invokes a http.Get on each passed URL and counts the string in the
// body. It blocks until there will be a free job.
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

// Counted returns accumulator value.
func (counter *URLCounter) Counted() int {
	counter.RLock()
	defer counter.RUnlock()
	return counter.counted
}

// Wait waits for all job workers to complete.
func (counter *URLCounter) Wait() {
	counter.wg.Wait()
}
