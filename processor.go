package main

import (
	"bufio"
	"io"
)

type Processor struct {
	MaxJobs int
	Source  io.Reader
	Counter AccumulatingCounter
}

func NewProcessor(max int, ior io.Reader, cntr AccumulatingCounter) *Processor {
	return &Processor{MaxJobs: max, Source: ior, Counter: cntr}
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
