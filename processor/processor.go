package processor

import (
	"bufio"
	"io"

	"github.com/kavu/word_counter/counters"
)

type Processor struct {
	Source  io.Reader
	Counter counters.ConcurrentAccumulatingCounter
}

func NewProcessor(ior io.Reader, cntr counters.ConcurrentAccumulatingCounter) *Processor {
	return &Processor{Source: ior, Counter: cntr}
}

func (processor *Processor) GetResult() int {
	return processor.Counter.Counted()
}

func (processor *Processor) Start() {
	scanner := bufio.NewScanner(processor.Source)

	for scanner.Scan() {
		line := scanner.Text()

		processor.Counter.Count(line)
	}
}

func (processor *Processor) Wait() {
	processor.Counter.Wait()
}
