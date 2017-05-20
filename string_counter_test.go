package main

import "testing"

func TestStringCounter(t *testing.T) {
	searchString := "123"
	examples := []struct {
		line     string
		expected int
	}{
		{"123", 1},
		{"1236123", 2},
		{"1236123\n123", 3},
		{"", 0},
	}
	expectedTotal := 6

	cntr := NewStringCounter(searchString)

	for _, ex := range examples {
		cntr.Count(ex.line)
	}

	if cntr.Counted() != expectedTotal {
		t.Errorf("Expected %d, but got %d", expectedTotal, cntr.Counted())
	}
}

func TestNewStringCounter(t *testing.T) {
	expected := "123"
	counter := NewStringCounter(expected)

	if got := counter.searchString; got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
	}
}
