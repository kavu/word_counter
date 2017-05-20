package main

import "testing"

func TestURLCounter(t *testing.T) {
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

	cntr := &URLCounter{searchString}

	for _, ex := range examples {
		if got := cntr.Count(ex.line); got != ex.expected {
			t.Errorf("Expected %d, but got %d", ex.expected, got)
		}
	}
}
