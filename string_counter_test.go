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
		if got := cntr.Count(ex.line); got != ex.expected {
			t.Errorf("Expected %d, but got %d", ex.expected, got)
		}
	}

	if cntr.Counted() != expectedTotal {
		t.Errorf("Expected %d, but got %d", expectedTotal, cntr.Counted())
	}
}
