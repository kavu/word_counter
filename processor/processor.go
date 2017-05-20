package processor

import (
	"bufio"
	"io"

	"github.com/kavu/word_counter/counters"
)

// Processor reads the input data from the Source and counts it with
// the ConcurrentAccumulatingCounter.
type Processor struct {
	Source  io.Reader
	Counter counters.ConcurrentAccumulatingCounter
}

// NewProcessor is just a factory function for Processor
func NewProcessor(ior io.Reader, cntr counters.ConcurrentAccumulatingCounter) *Processor {
	return &Processor{Source: ior, Counter: cntr}
}

// GetResult returns Count result from ConcurrentAccumulatingCounter
func (processor *Processor) GetResult() int {
	return processor.Counter.Counted()
}

// Start starts the buffered reading and counting
func (processor *Processor) Start() {
	scanner := bufio.NewScanner(processor.Source)

	for scanner.Scan() {
		line := scanner.Text()

		processor.Counter.Count(line)
	}
}

// Wait allows Processor to wait until ConcurrentAccumulatingCounter complete
// all jobs.
func (processor *Processor) Wait() {
	processor.Counter.Wait()
}
