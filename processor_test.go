package main

import (
	"os"
	"strings"
	"testing"
)

func TestProcessor(t *testing.T) {
	cntr := NewStringCounter("123")
	source := strings.NewReader("123asd\n\n123\nzzsd\123x123\n666\n\n\n444\n12")
	expected := 3

	proc := NewProcessor(1, source, cntr)

	proc.Start()

	if got := proc.GetResult(); got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
	}
}

func TestNewProcessor(t *testing.T) {
	cntr := NewStringCounter("123")
	source := os.Stdin

	processor := NewProcessor(1, source, cntr)

	if got := processor.MaxJobs; got != 1 {
		t.Errorf("Expected %d, but got %d", 1, got)
	}

	if got := processor.Counter.(*StringCounter).searchString; got != "123" {
		t.Errorf("Expected %s, but got %s", "123", got)
	}

	if got := processor.Source; got != os.Stdin {
		t.Errorf("Expected %#v, but got %#v", os.Stdin, got)
	}
}
