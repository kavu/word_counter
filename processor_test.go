package main

import (
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
