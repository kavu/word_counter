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
	expectedTotal := 6

	counter := NewURLCounter(searchString)

	for _, ex := range examples {
		counter.Count(ex.line)
	}
	counter.Wait()

	if counter.Counted() != expectedTotal {
		t.Errorf("Expected %d, but got %d", expectedTotal, counter.Counted())
	}
}
